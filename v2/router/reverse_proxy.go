package router

import (
	"context"
	"etri-sfpoc-edge/consulapi"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-zoox/proxy"
)

var mqttReverseProxy http.Handler
var consulReverseProxy http.Handler
var serviceReverseProxy http.Handler

func init() {
	mqttReverseProxy = proxy.New(&proxy.Config{
		OnRequest: func(req *http.Request, inReq *http.Request) error {
			fmt.Println(inReq.Header.Get("Sec-Websocket-Protocol") == "mqtt")
			req.URL.Host = "127.0.0.1:9998"
			return nil
		},
		OnResponse: func(res *http.Response, inReq *http.Request) error {
			fmt.Println(res)
			return nil
		},
		OnError: func(err error, rw http.ResponseWriter, req *http.Request) {
		},
		OnContext: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
		IsAnonymouse: false,
	})

	consulReverseProxy = proxy.New(&proxy.Config{
		OnRequest: func(req *http.Request, inReq *http.Request) error {
			req.URL.Host = "127.0.0.1:9999"
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/consul")
			return nil
		},
		OnResponse: func(res *http.Response, inReq *http.Request) error {

			if res.StatusCode == http.StatusMovedPermanently && len(inReq.Referer()) == 0 {
				res.Header.Set("Location", "/consul"+res.Header.Get("Location"))
				fmt.Println(res.Header)
				return nil
			}

			return nil
		},
		OnError: func(err error, rw http.ResponseWriter, req *http.Request) {
		},
		OnContext: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
		IsAnonymouse: false,
	})

	serviceReverseProxy = proxy.New(&proxy.Config{
		OnRequest: func(req *http.Request, inReq *http.Request) error {
			fmt.Println("service reverse proxy")
			fmt.Println(req.URL.Path)

			return nil
		},
	})
}

func reverseProxyHandle(c *gin.Context) {
	path := c.Param("any")
	id, _, ok := strings.Cut(path[5:], "/")
	if !ok {
		id = path[5:]
	}

	svcAddr, err := consulapi.GetSvcAddr(fmt.Sprintf("svcs/%s", id))
	if err != nil {
		panic(err)
	}

	remote, err := url.Parse("http://" + svcAddr)
	if err != nil {
		panic(err)
	}

	// serviceReverseProxies[svcAddr]

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("any")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
