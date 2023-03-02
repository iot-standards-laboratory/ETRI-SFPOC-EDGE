package router

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/mqtthandler"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostAgent(c *gin.Context) {
	if len(c.Param("any")) <= 1 {
		agent, err := DB.AddAgentWithJsonReader(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusCreated, agent)
	} else {
		id := c.Param("any")[1:]
		_, err := DB.GetAgent(id)
		if err != nil {
			panic(err)
		}
		params := connectionParams()

		addr, err := net.ResolveTCPAddr("tcp", c.Request.RemoteAddr)
		if err != nil {
			params["origin"] = ""
		} else {
			params["origin"] = addr.IP
		}

		c.JSON(http.StatusOK, params)
	}
}

func DeleteAgent(c *gin.Context) {
	agent_id := c.Request.Header.Get("agent_id")

	if len(agent_id) <= 1 {
		panic(errors.New("invalid agent id error"))
	} else {
		err := removeCtrlsWithAgentId(agent_id)
		if err != nil {
			panic(err)
		}
		err = consulapi.DeregisterCtrl(fmt.Sprintf("agent/%s", agent_id))
		if err != nil {
			panic(err)
		}
		err = DB.DeleteAgent(agent_id)
		if err != nil {
			panic(err)
		}

		mqtthandler.Publish("public/statuschanged", []byte("changed"))

		c.JSON(http.StatusOK, "deleted")
	}
}

func removeCtrlsWithAgentId(agentId string) error {
	// remove ctrls/{agentid}/ controller
	fmt.Printf("remove agentCtrls/%s\n", agentId)
	ctrlKeys, err := consulapi.GetKeys(fmt.Sprintf("agentCtrls/%s", agentId))
	if err != nil {
		return err
	}

	for _, key := range ctrlKeys {
		err = consulapi.Delete(key)
		if err != nil {
			return err
		}
	}

	ctrlKeys, err = consulapi.GetKeys("svcCtrls")
	if err != nil {
		return err
	}

	for _, key := range ctrlKeys {
		if strings.Contains(key, agentId) {
			err = consulapi.Delete(key)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetAgent(c *gin.Context) {
	agents, err := DB.GetAgents()
	if err != nil {
		panic(err)
	}

	agentsWithStatus := make([]map[string]interface{}, 0, len(agents))
	for _, agent := range agents {
		status, err := consulapi.GetStatus(fmt.Sprintf("agent/%s", agent.ID))
		if err != nil {
			continue
		}

		ctrls, err := getCtrlsWithAgentId(agent.ID)
		if err != nil {
			panic(err)
		}

		agentsWithStatus = append(agentsWithStatus, map[string]interface{}{
			"name":   agent.Name,
			"id":     agent.ID,
			"status": status,
			"ctrls":  ctrls,
		})

	}

	c.JSON(http.StatusOK, agentsWithStatus)
}

func getCtrlsWithAgentId(agentId string) ([]map[string]interface{}, error) {
	keys, err := consulapi.GetKeys(fmt.Sprintf("agentCtrls/%s", agentId))
	if err != nil {
		return nil, err
	}

	ctrls := make([]map[string]interface{}, 0, len(keys))
	for _, e := range keys {
		b_ctrl, err := consulapi.Get(e)
		if err != nil {
			return nil, err
		}
		m_ctrl := map[string]interface{}{}
		err = json.Unmarshal(b_ctrl, &m_ctrl)
		if err != nil {
			return nil, err
		}
		ctrls = append(ctrls, m_ctrl)
	}

	return ctrls, nil
}
