package router

import (
	"etri-sfpoc-edge/model"
	"etri-sfpoc-edge/notifier"
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
		apiv1.GET("/pub", Subscribe)
		apiv1.GET("/subs/list", GetSubscriberList)
		apiv1.POST("/ctrls", PostCtrl)
		apiv1.GET("/ctrls/list", GetCtrlList)
		apiv1.GET("/devs/list", GetDeviceList)
		apiv1.DELETE("/devs", DeleteDevice)
		apiv1.GET("/devs/discover/list", GetDiscoveredDevices)
		apiv1.POST("/devs/discover", PostDiscoveredDevice)
		apiv1.PUT("/devs/discover", PutDiscoveredDevice)

	}

	r := gin.New()

	assetEngine := gin.New()
	assetEngine.Static("/", "./static")
	r.Any("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiEngine.HandleContext(c)
		} else {
			assetEngine.HandleContext(c)
		}
	})

	return r
}

// Alarm
