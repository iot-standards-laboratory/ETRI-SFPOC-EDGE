package router

import (
	"etri-sfpoc-edge/consulapi"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func reverseProxyHandle(c *gin.Context) {
	defer handleError(c)
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
