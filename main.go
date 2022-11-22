package main

import (
	"context"
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/mqtthandler"
	"etri-sfpoc-edge/v2/consulapi"
	"etri-sfpoc-edge/v2/model/dbstorage"
	v2router "etri-sfpoc-edge/v2/router"
	"flag"
	"fmt"
	"os"
	"strings"
)

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
		err := mqtthandler.ConnectMQTT("tcp://localhost:2883")
		if err != nil {
			panic(err)
		}
		err = consulapi.Connect("http://localhost:9999")
		if err != nil {
			panic(err)
		}

		go consulapi.Monitor(func(what string) {
			if strings.Contains(what, "Synced check") {
				mqtthandler.Publish("public/statuschanged", []byte("changed"))

				agents, err := dbstorage.DefaultDB.GetAgents()
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

			}
		}, context.Background())

		v2router.NewRouter().Run(config.Params["bind"].(string))
	}

	// etrisfpocctnmgmt.CreateContainer("hello-world")
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

// var box = &RequestBox{notifier.NewNotiManager()}

// func WaitHandler(c *gin.Context) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	_uuid, _ := uuid.NewUUID()
// 	box.AddSubscriber(notifier.NewCallbackSubscriber(
// 		_uuid.String(),
// 		notifier.GenerateSecureToken(14),
// 		notifier.SubtypeOnce,
// 		func(msg string) {
// 			c.JSON(200, gin.H{
// 				"message": msg,
// 			})
// 			wg.Done()
// 		}))
// 	go routine()
// 	wg.Wait()
// }

// func routine() {
// 	time.Sleep(time.Second * 10)
// 	box.Publish(notifier.NewStatusChangedEvent("Hello world", "Hello world", notifier.SubtokenStatusChanged))
// }

// func WaitWithChannelHandler(c *gin.Context) {
// 	_uuid, _ := uuid.NewUUID()
// 	_channel := make(chan notifier.IEvent)
// 	box.AddSubscriber(notifier.NewChanSubscriber(
// 		_uuid.String(),
// 		// notifier.GenerateSecureToken(14),
// 		notifier.SubtokenStatusChanged,
// 		notifier.SubtypeOnce,
// 		_channel,
// 	))

// 	e := <-_channel
// 	c.Writer.Write([]byte(e.Title()))
// }
