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

// func GetSvcs(c *gin.Context) {
// 	defer handleError(c)

// 	w := c.Writer
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

// 	svcKeys, err := consulapi.GetKeys("svcs")
// 	if err != nil {
// 		panic(err)
// 	}
// 	svcs := make([]map[string]interface{}, 0, len(svcKeys))

// 	for _, e := range svcKeys {
// 		b_svc, err := consulapi.Get(e)
// 		if err != nil {
// 			panic(err)
// 		}
// 		j_svc := map[string]interface{}{}
// 		err = json.Unmarshal(b_svc, &j_svc)
// 		if err != nil {
// 			panic(err)
// 		}

// 		j_svc["status"] = "enabled"

// 		svcId, ok := j_svc["id"]
// 		if !ok {
// 			panic(errors.New("invalid service name error"))
// 		}

// 		ctrlKeys, err := consulapi.GetKeys(fmt.Sprintf("svcCtrls/%s", svcId))
// 		if err != nil {
// 			panic(err)
// 		}
// 		j_svc["num_clnts"] = len(ctrlKeys)

// 		cid, ok := j_svc["cid"].(string)
// 		if !ok || len(cid) == 0 {
// 			j_svc["status"] = "disabled"
// 		} else {
// 			status, err := consulapi.GetStatus(fmt.Sprintf("svcs/%s", cid))
// 			if err != nil || strings.Compare(status, "passing") != 0 {
// 				j_svc["status"] = "disabled"
// 			}
// 		}

// 		svcs = append(svcs, j_svc)
// 	}

// 	c.JSON(http.StatusOK, svcs)
// }

func PutSvcs(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	id := c.Request.Header.Get("id")
	if len(id) <= 0 {
		panic(errors.New("invalid service id error"))
	}

	_uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	fmt.Println(_uuid)
	svc, err := DB.SelectSvc(_uuid)
	if err != nil {
		panic(err)
	}

	originAddr, _ := net.ResolveTCPAddr("tcp", c.Request.RemoteAddr)
	svc.Status = "running"
	svc.Addr = originAddr.IP.String()
	_, err = DB.UpdateSvc(svc)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"conn_param": connectionParams(c),
		"origin":     originAddr.IP,
	}

	fmt.Println(payload)
	c.JSON(http.StatusOK, payload)
}

func PostSvcs(c *gin.Context) {
	defer handleError(c)

	var obj map[string]interface{}

	dec := json.NewDecoder(c.Request.Body)
	err := dec.Decode(&obj)
	if err != nil {
		panic(err)
	}

	// fmt.Println(obj)
	t, ok := obj["type"].(string)
	if !ok {
		fmt.Println("Invalid!!!")
		panic(errors.New("invalid message error"))
	}

	if t == "DELETE" {
		// delete svcs
	} else if t == "INSERT" {
		// install svcs

		imgId, err := uuid.Parse(obj["record"].(map[string]interface{})["img_id"].(string))
		if err != nil {
			panic(err)
		}

		fmt.Println(imgId)
		img, err := DB.SelectImg(imgId)
		if err != nil {
			panic(err)
		}

		fmt.Println(img.ImageURL)
	}
	// // 초기 등록
	// id := c.Request.Header.Get("service_id")
	// if len(id) <= 0 {
	// 	panic(errors.New("invalid service name error"))
	// }

	// b_svcInfo, err := consulapi.Get(fmt.Sprintf("svcs/%s", id))
	// if err != nil {
	// 	panic(err)
	// }

	// m_svcInfo := map[string]interface{}{}
	// err = json.Unmarshal(b_svcInfo, &m_svcInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// // check the identifier of container for the service
	// cid := m_svcInfo["cid"].(string)
	// if len(cid) > 0 {
	// 	panic(errors.New("already installed service"))
	// }

	// cid = uuid.New().String()

	// m_svcInfo["cid"] = cid
	// b_svcInfo, err = json.Marshal(m_svcInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// err = consulapi.Put(fmt.Sprintf("svcs/%s", id), b_svcInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// err = createContainer(cid, id)
	// if err != nil {
	// 	// remove the container information from consul
	// 	panic(err)
	// }

	c.String(http.StatusCreated, "installed")
}

func DeleteSvcs(c *gin.Context) {
	defer handleError(c)

	id := c.Request.Header.Get("service_id")
	if len(id) <= 0 {
		panic(errors.New("invalid service id error"))
	}

	b_svcInfo, err := consulapi.Get(fmt.Sprintf("svcs/%s", id))
	if err != nil {
		panic(err)
	}

	m_svcInfo := map[string]interface{}{}
	err = json.Unmarshal(b_svcInfo, &m_svcInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println(m_svcInfo)

	cid, ok := m_svcInfo["cid"].(string)
	if !ok {
		panic(errors.New("invalid service id error"))
	}

	consulapi.DeregisterCtrl("svcs/" + cid)

	m_svcInfo["cid"] = ""

	b_obj, err := json.Marshal(m_svcInfo)
	if err != nil {
		panic(err)
	}

	err = consulapi.Put(fmt.Sprintf("svcs/%s", id), b_obj)
	if err != nil {
		panic(err)
	}

	fmt.Println("remove container:", cid)
	err = deleteContainer(cid)
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

func createContainer(id, img_url, name string) error {

	if isExist(id) {
		return nil
	}
	args := strings.Split(fmt.Sprintf("container\\run\\--restart\\always\\--env\\id=%s\\--add-host=host.docker.internal:host-gateway\\--name\\%s\\-d\\%s", id, name, img_url), "\\")
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
