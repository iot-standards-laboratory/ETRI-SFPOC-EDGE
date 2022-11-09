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
