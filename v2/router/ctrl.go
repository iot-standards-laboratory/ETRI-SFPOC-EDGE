package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/mqtthandler"
	"etri-sfpoc-edge/v2/consulapi"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	ctrlKeys, err := consulapi.GetKeys("svcCtrls")
	if err != nil {
		panic(err)
	}

	l_ctrls := make([]map[string]interface{}, 0, len(ctrlKeys))

	for _, key := range ctrlKeys {
		b, err := consulapi.Get(key)
		if err != nil {
			panic(err)
		}

		m_ctrl := map[string]interface{}{}
		err = json.Unmarshal(b, &m_ctrl)
		if err != nil {
			panic(err)
		}

		l_ctrls = append(l_ctrls, m_ctrl)
	}

	c.JSON(http.StatusOK, l_ctrls)
}
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

	ctrlId, ok := payload["id"].(string)
	if !ok {
		panic(errors.New("invalid controller id error"))
	}

	json_payload, _ := json.Marshal(payload)
	err = consulapi.Put(
		fmt.Sprintf("agentCtrls/%s/%s", agentId, ctrlId),
		json_payload,
	)
	if err != nil {
		panic(err)
	}

	err = consulapi.Put(
		fmt.Sprintf("svcCtrls/%s/%s/%s", svcName, agentId, ctrlId),
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
		err = consulapi.Put(
			fmt.Sprintf("svcs/%s", svcName),
			svcJson,
		)
		if err != nil {
			panic(err)
		}
	}

	mqtthandler.Publish("public/statuschanged", []byte("changed"))
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

	ctrlId, ok := payload["id"].(string)
	if !ok {
		panic(errors.New("invalid controller id error"))
	}

	err = consulapi.Delete(fmt.Sprintf("agentCtrls/%s/%s", agentId, ctrlId))
	if err != nil {
		panic(err)
	}

	err = consulapi.Delete(fmt.Sprintf("svcCtrls/%s/%s/%s", svcName, agentId, ctrlId))
	if err != nil {
		panic(err)
	}
	mqtthandler.Publish("public/statuschanged", []byte("changed"))
	c.String(http.StatusOK, "OK")
}