package main

import (
	"context"
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/controller/connmgmt"
	"etri-sfpoc-edge/controller/state"
	"etri-sfpoc-edge/v2/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	stateStream := state.Subscribe()
	if _, err := os.Stat("./config.properties"); errors.Is(err, os.ErrNotExist) {
		state.Put(state.STATE_INITIALIZING)
	} else {
		state.Put(state.STATE_RUNNING)
	}

	// go v2router.NewRouter().Run(config.Params["bind"].(string))
	var srv *http.Server = nil

	for currentState := range stateStream {
		fmt.Println(currentState)
		switch currentState {
		case state.STATE_INITIALIZING:
			mux, err := router.NewRouter(state.STATE_INITIALIZING)
			if err != nil {
				panic(err)
			}
			srv = &http.Server{
				Addr:    ":9995",
				Handler: mux,
			}

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}

				log.Println("close initializing")
			}()

		case state.STATE_INITIALIZED:
			err := srv.Shutdown(context.Background())
			if err != nil {
				panic(err)
			}

			config.CreateInitFile()
			state.Put(state.STATE_RUNNING)

		case state.STATE_RUNNING:
			config.LoadConfig()
			err := connmgmt.Connect(config.Params["mqttAddr"].(string), config.Params["consulAddr"].(string))
			if err != nil {
				panic(err)
			}

			// go consulapi.Monitor(func(what string) {
			// 	if strings.Contains(what, "Synced check") {
			// 		// agents, err := consulstorage.DefaultDB.GetAgents()
			// 		if err != nil {
			// 			return
			// 		}
			// 		for _, agent := range agents {
			// 			status, err := consulapi.GetStatus(fmt.Sprintf("agent/%s", agent.ID))
			// 			if err != nil {
			// 				return
			// 			}

			// 			if strings.Compare(status, "passing") != 0 {
			// 				err := removeCtrlsWithAgentId(agent.ID)
			// 				if err != nil {
			// 					return
			// 				}
			// 			}
			// 		}
			// 		mqtthandler.Publish("public/statuschanged", []byte("changed"))
			// 	}
			// }, context.Background())

			mux, err := router.NewRouter(state.STATE_RUNNING)
			if err != nil {
				panic(err)
			}

			srv = &http.Server{
				Addr:    ":9995",
				Handler: mux,
			}
			srv.ListenAndServe()
		}
	}

	// if *cfg {
	// 	config.CreateInitFile()
	// } else {
	// 	// 상태 변경 시 알림 전달

	// }
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
