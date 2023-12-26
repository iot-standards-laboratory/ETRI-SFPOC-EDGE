package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/mqtthandler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCtrl(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	id := c.Request.Header.Get("id")
	key := "svcCtrls"
	if len(id) > 0 {
		key = fmt.Sprintf("svcCtrls/%s", id)
	}

	ctrlKeys, err := consulapi.GetKeys(key)
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

	ctrlName, ok := payload["name"].(string)
	if !ok {
		panic(errors.New("invalid agent id error"))
	}

	ctrl, err := DB.InsertCtrl(ctrlName)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, ctrl)
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

	svcId, ok := payload["service_id"].(string)
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

	err = consulapi.Delete(fmt.Sprintf("svcCtrls/%s/%s/%s", svcId, agentId, ctrlId))
	if err != nil {
		panic(err)
	}
	mqtthandler.Publish("public/statuschanged", []byte("changed"))
	c.String(http.StatusOK, "OK")
}
