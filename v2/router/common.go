package router

import (
	"etri-sfpoc-edge/logger"
	"etri-sfpoc-edge/model/consulstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB = consulstorage.DefaultDB

func handleError(c *gin.Context) {
	if r := recover(); r != nil {
		logger.Println(r)
		c.String(http.StatusBadRequest, r.(error).Error())
	}
}

func connectionParams() map[string]interface{} {
	// wsAddr, _ := config.Params("wsAddr")
	// wsAddr, _ := config.Params("wsAddr")
	// wsAddr, _ := config.Params("wsAddr")
	return map[string]interface{}{
		// "wsAddr":     "wss://mqtt.godopu.com:8000/connection/websocket",
		"consulAddr": "http://etri.godopu.com:9999",
		"mqttAddr":   "wss://mqtt.godopu.com",
	}
}
