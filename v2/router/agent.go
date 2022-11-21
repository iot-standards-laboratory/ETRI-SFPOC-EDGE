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
		agent, err := db.AddAgentWithJsonReader(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusCreated, agent)
	} else {
		id := c.Param("any")[1:]
		_, err := db.GetAgent(id)
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

	agents, err := db.GetAgents()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(agents))
	ctrlsWithStatus := make([]map[string]interface{}, 0, len(agents))
	for _, agent := range agents {
		status, err := consulapi.GetStatus(fmt.Sprintf("agent/%s", agent.ID))
		if err != nil {
			panic(err)
		}
		ctrlsWithStatus = append(ctrlsWithStatus, map[string]interface{}{
			"name":   agent.Name,
			"id":     agent.ID,
			"status": status,
		})

		fmt.Println(status)
	}
	consulapi.GetStatus("")
	c.JSON(http.StatusOK, ctrlsWithStatus)
}
