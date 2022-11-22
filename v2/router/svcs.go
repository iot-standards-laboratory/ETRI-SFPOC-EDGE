package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/v2/consulapi"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSvcs(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	svcKeys, err := consulapi.GetKeys("svcs")
	if err != nil {
		panic(err)
	}
	svcs := make([]map[string]interface{}, 0, len(svcKeys))

	for _, e := range svcKeys {
		b_svc, err := consulapi.Get(e)
		if err != nil {
			panic(err)
		}
		j_svc := map[string]interface{}{}
		err = json.Unmarshal(b_svc, &j_svc)
		if err != nil {
			panic(err)
		}

		svcName, ok := j_svc["name"]
		if !ok {
			panic(errors.New("invalid service name error"))
		}
		ctrlKeys, err := consulapi.GetKeys(fmt.Sprintf("svcCtrls/%s", svcName))
		if err != nil {
			panic(err)
		}

		j_svc["num_clnts"] = len(ctrlKeys)
		svcs = append(svcs, j_svc)
	}

	c.JSON(http.StatusOK, svcs)
}

func PutSvcs(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	name := c.Request.Header.Get("name")
	b_svcInfo, err := consulapi.Get(fmt.Sprintf("/svcs/%s", name))
	if err != nil {
		panic(err)
	}

	m_svcInfo := map[string]interface{}{}
	err = json.Unmarshal(b_svcInfo, &m_svcInfo)
	if err != nil {
		panic(err)
	}

	id, ok := m_svcInfo["id"].(string)
	if !ok || len(id) <= 0 {
		id := uuid.New().String()
		m_svcInfo["id"] = id
		b_svcInfo, err = json.Marshal(m_svcInfo)
		if err != nil {
			panic(err)
		}
		err = consulapi.Put(fmt.Sprintf("svcs/%s", name), b_svcInfo)
		if err != nil {
			panic(err)
		}
	}

	originAddr, _ := net.ResolveTCPAddr("tcp", c.Request.RemoteAddr)

	payload := map[string]interface{}{
		"conn_param": connectionParams(),
		"info":       m_svcInfo,
		"origin":     originAddr.IP,
	}

	c.JSON(http.StatusOK, payload)
}

func PostSvcs(c *gin.Context) {
	defer handleError(c)

	// 초기 등록
	if len(c.Param("any")) <= 1 {

		// controller에 의해 등록되어 있는지 확인

	}

	// 내용 수정
}
