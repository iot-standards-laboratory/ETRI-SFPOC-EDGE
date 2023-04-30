package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/mqtthandler"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"

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

		j_svc["status"] = "enabled"
		svcId, ok := j_svc["id"]
		if !ok {
			j_svc["status"] = "disabled"
		}
		// _, err = consulapi.GetSvcAddr(fmt.Sprintf("svcs/%s", svcId))
		// if err != nil {
		// 	j_svc["status"] = "disabled"
		// }

		status, err := consulapi.GetStatus(fmt.Sprintf("svcs/%s", svcId))
		if err != nil || strings.Compare(status, "passing") != 0 {
			j_svc["status"] = "disabled"
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
	if len(name) <= 0 {
		panic(errors.New("invalid service name error"))
	}
	b_svcInfo, err := consulapi.Get(fmt.Sprintf("svcs/%s", name))
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
		panic(errors.New("invalid id"))
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
	name := c.Request.Header.Get("service_name")
	if len(name) <= 0 {
		panic(errors.New("invalid service name error"))
	}
	b_svcInfo, err := consulapi.Get(fmt.Sprintf("svcs/%s", name))
	if err != nil {
		panic(err)
	}

	m_svcInfo := map[string]interface{}{}
	err = json.Unmarshal(b_svcInfo, &m_svcInfo)
	if err != nil {
		panic(err)
	}

	svcId := m_svcInfo["id"].(string)
	if len(svcId) > 0 {
		panic(errors.New("already installed service"))
	}

	svcId = uuid.New().String()
	m_svcInfo["id"] = svcId
	b_svcInfo, err = json.Marshal(m_svcInfo)
	if err != nil {
		panic(err)
	}
	err = consulapi.Put(fmt.Sprintf("svcs/%s", name), b_svcInfo)
	if err != nil {
		panic(err)
	}

	err = createContainer(svcId, name)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusCreated, "installed")
	// 내용 수정
}

func DeleteSvcs(c *gin.Context) {
	defer handleError(c)

	// 초기 등록
	name := c.Request.Header.Get("service_name")
	if len(name) <= 0 {
		panic(errors.New("invalid service name error"))
	}

	b_svcInfo, err := consulapi.Get(fmt.Sprintf("svcs/%s", name))
	if err != nil {
		panic(err)
	}

	m_svcInfo := map[string]interface{}{}
	err = json.Unmarshal(b_svcInfo, &m_svcInfo)
	if err != nil {
		panic(err)
	}
	svcId, ok := m_svcInfo["id"].(string)
	if !ok {
		panic(errors.New("invalid service id error"))
	}

	consulapi.DeregisterCtrl("svcs/" + svcId)

	m_svcInfo["id"] = ""

	b_obj, err := json.Marshal(m_svcInfo)
	if err != nil {
		panic(err)
	}

	err = consulapi.Put(fmt.Sprintf("svcs/%s", name), b_obj)
	if err != nil {
		panic(err)
	}

	err = deleteContainer(svcId)
	if err != nil {
		panic(err)
	}

	mqtthandler.Publish("public/statuschanged", []byte("changed"))
	c.String(http.StatusOK, "removed")
}

func isExist(id string) bool {
	cmd := strings.Split("container\\ls\\--format\\{{.Names}}\\-a", "\\")
	fmt.Println(cmd)
	bout, err := exec.Command("docker", cmd...).Output()
	if err != nil {
		log.Fatalln(err)
	}

	sout := strings.Split(string(bout), "\n")

	for _, e := range sout {
		if strings.Compare(e, id) == 0 {
			return true
		}
	}

	return false
}

func createContainer(id, name string) error {

	if isExist(id) {
		return nil
	}
	args := strings.Split(fmt.Sprintf("container\\run\\--restart\\always\\--name\\%s\\-d\\%s", id, name), "\\")
	fmt.Println(args)
	_, err := exec.Command("docker", args...).Output()
	if err != nil {
		return err
	}

	return nil
}

func deleteContainer(id string) error {
	if !isExist(id) {
		fmt.Println("not exist")
		return nil
	}

	args := strings.Split(fmt.Sprintf("container\\rm\\-v\\-f\\%s", id), "\\")
	fmt.Println(args)
	_, err := exec.Command("docker", args...).Output()
	if err != nil {
		return err
	}

	return nil
}
