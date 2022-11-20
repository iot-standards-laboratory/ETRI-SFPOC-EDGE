package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func PostCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if len(c.Param("any")) <= 1 {
		ctrl, err := db.AddControllerWithJsonReader(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusCreated, ctrl)
	} else {
		cid := c.Param("any")[1:]
		_, err := db.GetController(cid)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, connectionParams())
	}
}
