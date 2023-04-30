package main

import (
	"context"
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/model/consulstorage"
	"etri-sfpoc-edge/mqtthandler"
	v2router "etri-sfpoc-edge/v2/router"
	"flag"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	err := consulapi.Connect("http://etri.godopu.com:9999")
// 	if err != nil {
// 		panic(err)
// 	}

// 	agent, err := consulstorage.DefaultDB.GetAgent("ed62e012-71f9-4e4d-8a84-d28bb0862cc5")
// 	if err != nil {
// 		panic(err)
// 	}
// 	consulapi.RegisterAgent(*agent, "http://129.254.164.223:4000")
// }

func main() {

	cfg := flag.Bool("init", false, "create initial config file")
	flag.Parse()
	if *cfg {
		config.CreateInitFile()
	} else {
		if _, err := os.Stat("./config.properties"); errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does not exist
			fmt.Println("config file doesn't exist")
			fmt.Println("please add -init option to create config file")
			return
		}
		config.LoadConfig()
		err := mqtthandler.ConnectMQTT("wss://mqtt.godopu.com")
		if err != nil {
			panic(err)
		}
		err = consulapi.Connect("http://localhost:9999")
		if err != nil {
			panic(err)
		}

		// 상태 변경 시 알림 전달
		go consulapi.Monitor(func(what string) {
			if strings.Contains(what, "Synced check") {
				agents, err := consulstorage.DefaultDB.GetAgents()
				if err != nil {
					return
				}
				for _, agent := range agents {
					status, err := consulapi.GetStatus(fmt.Sprintf("agent/%s", agent.ID))
					if err != nil {
						return
					}

					if strings.Compare(status, "passing") != 0 {
						err := removeCtrlsWithAgentId(agent.ID)
						if err != nil {
							return
						}
					}
				}
				mqtthandler.Publish("public/statuschanged", []byte("changed"))
			}
		}, context.Background())

		v2router.NewRouter().Run(config.Params["bind"].(string))
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
