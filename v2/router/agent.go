package router

import (
	"etri-sfpoc-edge/v2/consulapi"
	"fmt"
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

func PostAgent(c *gin.Context) {
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

func GetAgent(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	ctrls, err := db.GetControllers()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(ctrls))
	ctrlsWithStatus := make([]map[string]interface{}, 0, len(ctrls))
	for _, ctrl := range ctrls {
		status, err := consulapi.GetStatus(fmt.Sprintf("ctrl/%s", ctrl.CID))
		if err != nil {
			panic(err)
		}
		ctrlsWithStatus = append(ctrlsWithStatus, map[string]interface{}{
			"ctrl":   ctrl,
			"status": status,
		})

		fmt.Println(status)
	}
	consulapi.GetStatus("")
	c.JSON(http.StatusOK, ctrlsWithStatus)
}
