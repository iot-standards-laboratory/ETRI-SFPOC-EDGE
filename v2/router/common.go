package router

import (
	"etri-sfpoc-edge/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		"wsAddr":     "ws://localhost:8000/connection/websocket",
		"consulAddr": "http://localhost:9999",
		"mqttAddr":   "tcp://localhost:2883",
	}
}
