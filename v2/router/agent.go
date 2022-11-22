package router

import (
	"encoding/json"
	"etri-sfpoc-edge/v2/consulapi"
	"etri-sfpoc-edge/v2/model/dbstorage"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)



func PostAgent(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if len(c.Param("any")) <= 1 {
		agent, err := dbstorage.DefaultDB.AddAgentWithJsonReader(c.Request.Body)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(http.StatusCreated, agent)
	} else {
		id := c.Param("any")[1:]
		_, err := dbstorage.DefaultDB.GetAgent(id)
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

func GetAgent(c *gin.Context) {
	defer handleError(c)

	w := c.Writer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	agents, err := dbstorage.DefaultDB.GetAgents()
	if err != nil {
		panic(err)
	}

	agentsWithStatus := make([]map[string]interface{}, 0, len(agents))
	for _, agent := range agents {
		status, err := consulapi.GetStatus(fmt.Sprintf("agent/%s", agent.ID))
		if err != nil {
			panic(err)
		}

		// ctrls, err := getCtrlsWithAgentId(agent.ID)
		// if err != nil {
		// 	panic(err)
		// }

		agentsWithStatus = append(agentsWithStatus, map[string]interface{}{
			"name":   agent.Name,
			"id":     agent.ID,
			"status": status,
			// "ctrls":  ctrls,
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
