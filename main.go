package main

import (
	"context"
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/mqtthandler"
	v1router "etri-sfpoc-edge/v1/router"
	"etri-sfpoc-edge/v2/consulapi"
	v2router "etri-sfpoc-edge/v2/router"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	cfg := flag.Bool("init", false, "create initial config file")
	version := flag.String("version", "v2", "specify version")
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
		mqtthandler.ConnectMQTT("localhost:2883")
		err := consulapi.Connect("http://localhost:9999")
		if err != nil {
			panic(err)
		}

		go consulapi.Monitor(func(what string) {
			if strings.Contains(what, "Synced check") {
				fmt.Println("What:", what)
				mqtthandler.Publish("public/statuschanged", []byte("changed"))
			}
		}, context.Background())

		if strings.Compare(*version, "v1") == 0 {
			v1router.NewRouter().Run(config.Params["bind"].(string))
		} else {
			v2router.NewRouter().Run(config.Params["bind"].(string))
		}
	}

	// etrisfpocctnmgmt.CreateContainer("hello-world")
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
