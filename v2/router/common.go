package router

import (
	"etri-sfpoc-edge/config"
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
		"consulAddr": config.Params["consulAddr"].(string),
		"mqttAddr":   config.Params["mqttAddr"].(string),
	}
}
