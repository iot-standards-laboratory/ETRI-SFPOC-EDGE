package main

import (
	"errors"
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/notifier"
	"etri-sfpoc-edge/router"
	"flag"
	"fmt"
	"os"
)

type RequestBox struct {
	notifier.INotiManager
}

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
		router.NewRouter().Run(config.Params["bind"].(string))
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
