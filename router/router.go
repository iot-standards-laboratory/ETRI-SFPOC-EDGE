package router

import (
	"etri-sfpoc-edge/logger"
	"etri-sfpoc-edge/model"
	"etri-sfpoc-edge/notifier"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestBox struct {
	notifier.INotiManager
}

var box *RequestBox

var db model.DBHandlerI

func init() {
	box = &RequestBox{notifier.NewNotiManager()}

	var err error
	db, err = model.NewSqliteHandler("dump.db")
	if err != nil {
		panic(err)
	}
	// go fire()
}

func NewRouter() *gin.Engine {
	apiEngine := gin.New()

	apiv1 := apiEngine.Group("api/v1")
	{
		apiv1.GET("/subs/list", GetSubscriberList)
		apiv1.POST("/ctrls", PostCtrl)
		apiv1.GET("/ctrls/list", GetCtrlList)
		apiv1.GET("/devs/list", GetDeviceList)
		apiv1.DELETE("/devs", DeleteDevice)
		apiv1.GET("/devs/discover/list", GetDiscoveredDevices)
		apiv1.POST("/devs/discover", PostDiscoveredDevice)
		apiv1.PUT("/devs/discover", PutDiscoveredDevice)
		apiv1.GET("/svcs/list", GetServiceList)
		apiv1.GET("/svcs", GetServiceInfo)
		apiv1.PUT("/svcs", PutService)
		apiv1.POST("/svcs", PostService)
	}

	pushEngine := gin.New()
	pushEngine.GET("/*any", func(c *gin.Context) {
		GetPublish(c)
	})
	pushEngine.POST("/*any", func(c *gin.Context) {
		PostPublish(c)
	})

	svcEngine := gin.New()
	svcEngine.Any("/*any", func(c *gin.Context) {
		SvcBroker(c)
	})

	assetEngine := gin.New()
	assetEngine.Static("/", "./sfpoc_edge_front/build/web")
	// assetEngine.Static("/", "./static")
	r := gin.New()
	r.Any("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api/v1") {
			apiEngine.HandleContext(c)
		} else if strings.HasPrefix(path, "/push/v1") {
			pushEngine.HandleContext(c)
		} else if strings.HasPrefix(path, "/svc/") {
			svcEngine.HandleContext(c)
		} else {
			assetEngine.HandleContext(c)
		}
	})

	return r
}

func Test(c *gin.Context) {
	// logger.Println(c.Query("hello"))
	path := c.Param("any")
	logger.Println(path)
	logger.Println(c.Request.Method)
	c.String(http.StatusOK, "Hello world")
}

// Alarm
