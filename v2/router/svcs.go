package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/v2/consulapi"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func PostSvcs(c *gin.Context) {
	// controller에 의해 등록되어 있는지 확인
	// 내용 수정
}
