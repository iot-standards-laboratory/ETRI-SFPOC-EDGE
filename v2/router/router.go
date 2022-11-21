package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	apiEngine := gin.New()
	apiv2 := apiEngine.Group("api/v2")
	{
		// _ = apiv2
		apiv2.POST("/agents/*any", PostAgent)
		apiv2.GET("/agents/*any", GetAgent)

		apiv2.GET("/ctrls/*any", GetCtrl)
		apiv2.POST("/ctrls/*any", PostCtrl)
		apiv2.DELETE("/ctrls/*any", DeleteCtrl)

		apiv2.GET("/svcs/*any", GetSvcs)

	}

	assetEngine := gin.New()
	assetEngine.Static("/", "./front/build/web")

	r := gin.New()
	r.Any("/*any", func(c *gin.Context) {
		remote, err := getRemoteURL(c.Request.Host)
		if err != nil {
			c.String(http.StatusNoContent, "wrong host is indicated")
			return
		}

		if remote == nil {
			path := c.Param("any")
			if strings.HasPrefix(path, "/api/v2") {
				apiEngine.HandleContext(c)
			} else {
				assetEngine.HandleContext(c)
			}
			return
		}

		reverseProxy(c, remote)

	})

	return r
}

func getRemoteURL(host string) (*url.URL, error) {
	// return nil, nil

	if !strings.HasPrefix(host, "svc.") {
		return nil, nil
	}
	remote, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	return remote, nil
}

func reverseProxy(c *gin.Context, remote *url.URL) {
	rp := httputil.NewSingleHostReverseProxy(remote)
	rp.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("any")
	}

	rp.ServeHTTP(c.Writer, c.Request)
}
