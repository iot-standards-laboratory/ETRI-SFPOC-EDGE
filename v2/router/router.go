package router

import (
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/controller/state"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewRouter(st state.IState) (*gin.Engine, error) {
	switch st {
	case state.STATE_INITIALIZING:
		return NewInitRouter(), nil
	case state.STATE_RUNNING:
		return NewRunningRouter(), nil
	default:
		return nil, errors.New("invalid state error")
	}
}

func NewInitRouter() *gin.Engine {
	assetEngine := gin.New()
	assetEngine.Static("/", "../ETRI-SFPOC-EDGE_front/web")

	r := gin.New()
	r.Any("/*any", func(c *gin.Context) {
		defer handleError(c)
		w := c.Writer
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		path := c.Param("any")
		// assetEngine.HandleContext(c)
		if path == "/loading" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"page": "/init",
			})

		} else if path == "/init" && c.Request.Method == "POST" {
			payload := map[string]interface{}{}
			err := c.BindJSON(&payload)
			if err != nil {
				panic(err)
			}

			err = parameterCheck(payload, []string{"consulAddr", "mqttAddr"})
			if err != nil {
				panic(err)
			}

			c.Status(http.StatusOK)

			config.Params["consulAddr"] = payload["consulAddr"]
			config.Params["mqttAddr"] = payload["mqttAddr"]
			state.Put(state.STATE_INITIALIZED)

		} else if len(path) <= 1 || strings.Contains(c.Request.Header.Get("Referer"), "localhost:3000") {
			assetEngine.HandleContext(c)
		} else {
			c.Status(http.StatusBadRequest)
		}
	})

	return r
}

func NewRunningRouter() *gin.Engine {
	apiEngine := gin.New()
	apiv2 := apiEngine.Group("api/v2")
	{
		// _ = apiv2
		apiv2.POST("/agents/*any", PostAgent)
		apiv2.GET("/agents/*any", GetAgent)
		apiv2.DELETE("/agents/*any", DeleteAgent)

		apiv2.GET("/ctrls/*any", GetCtrl)
		apiv2.POST("/ctrls/*any", PostCtrl)
		apiv2.DELETE("/ctrls/*any", DeleteCtrl)

		apiv2.GET("/svcs/*any", GetSvcs)
		apiv2.PUT("/svcs/*any", PutSvcs)
		apiv2.POST("/svcs/*any", PostSvcs)
		apiv2.DELETE("/svcs/*any", DeleteSvcs)

		apiv2.GET("/home", GetHome)
	}

	reverseProxyEngine := gin.New()
	reverseProxyEngine.Any("/*any", reverseProxyHandle)

	assetEngine := gin.New()
	assetEngine.Static("/", "../ETRI-SFPOC-EDGE_front/build/web")

	r := gin.New()
	r.Any("/*any", func(c *gin.Context) {
		defer handleError(c)
		w := c.Writer
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		remote, err := getRemoteURL(c.Request.Host)
		if err != nil {
			c.String(http.StatusNoContent, "wrong host is indicated")
			return

		}

		if remote == nil {
			path := c.Param("any")
			if strings.HasPrefix(path, "/api/v2") {
				apiEngine.HandleContext(c)
			} else if strings.HasPrefix(path, "/svc/") {
				reverseProxyEngine.HandleContext(c)
			} else if path == "/loading" {
				c.JSON(http.StatusOK, map[string]interface{}{
					"page": "/home",
				})
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
