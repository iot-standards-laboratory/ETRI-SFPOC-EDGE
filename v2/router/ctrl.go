package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/v2/consulapi"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	payload := map[string]interface{}{}
	err := c.BindJSON(&payload)
	if err != nil {
		panic(err)
	}
	agentId, ok := payload["agent_id"].(string)
	if !ok {
		panic(errors.New("invalid agent id error"))
	}

	svcName, ok := payload["service_name"].(string)
	if !ok {
		panic(errors.New("invalid service name error"))
	}

	fmt.Println("payload:", payload)
	ctrlId, ok := payload["id"].(string)
	if !ok {
		panic(errors.New("invalid controller id error"))
	}

	json_payload, _ := json.Marshal(payload)
	err = consulapi.Put(
		fmt.Sprintf("ctrls/%s/%s", agentId, ctrlId),
		json_payload,
	)
	if err != nil {
		panic(err)
	}

	err = consulapi.Put(
		fmt.Sprintf("ctrls/%s/%s/%s", svcName, agentId, ctrlId),
		json_payload,
	)
	if err != nil {
		panic(err)
	}

	svcJson, err := json.Marshal(map[string]interface{}{
		"name": svcName,
		"id":   "",
	})
	if err != nil {
		panic(err)
	}

	b, err := consulapi.Get(fmt.Sprintf("svcs/%s", svcName))
	if b == nil || err != nil {
		fmt.Println(err)
		err = consulapi.Put(
			fmt.Sprintf("svcs/%s", svcName),
			svcJson,
		)
		if err != nil {
			panic(err)
		}
	}

	c.String(http.StatusOK, "OK")
}

func DeleteCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	payload := map[string]interface{}{}
	err := c.BindJSON(&payload)
	if err != nil {
		panic(err)
	}
	agentId, ok := payload["agent_id"].(string)
	if !ok {
		panic(errors.New("invalid agent id error"))
	}

	svcName, ok := payload["service_name"].(string)
	if !ok {
		panic(errors.New("invalid service name error"))
	}

	fmt.Println("payload:", payload)
	ctrlId, ok := payload["id"].(string)
	if !ok {
		panic(errors.New("invalid controller id error"))
	}

	err = consulapi.Delete(fmt.Sprintf("ctrls/%s/%s", agentId, ctrlId))
	if err != nil {
		panic(err)
	}

	err = consulapi.Delete(fmt.Sprintf("ctrls/%s/%s/%s", svcName, agentId, ctrlId))
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, "OK")
}
