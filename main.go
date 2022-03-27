package main

import (
	"etri-sfpoc-edge/config"
	"etri-sfpoc-edge/notifier"
	"etri-sfpoc-edge/router"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type RequestBox struct {
	notifier.INotiManager
}

var box = &RequestBox{notifier.NewNotiManager()}

func main() {

	router.NewRouter().Run(config.Params["bind"].(string))
}

func WaitHandler(c *gin.Context) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	_uuid, _ := uuid.NewV4()
	box.AddSubscriber(notifier.NewCallbackSubscriber(
		_uuid.String(),
		notifier.GenerateSecureToken(14),
		notifier.SubtypeOnce,
		func(msg string) {
			c.JSON(200, gin.H{
				"message": msg,
			})
			wg.Done()
		}))
	go routine()
	wg.Wait()
}

func WaitWithChannelHandler(c *gin.Context) {
	_uuid, _ := uuid.NewV4()
	_channel := make(chan notifier.IEvent)
	box.AddSubscriber(notifier.NewChanSubscriber(
		_uuid.String(),
		// notifier.GenerateSecureToken(14),
		notifier.SubtokenStatusChanged,
		notifier.SubtypeOnce,
		_channel,
	))

	e := <-_channel
	c.Writer.Write([]byte(e.Title()))
}

func routine() {
	time.Sleep(time.Second * 10)
	box.Publish(notifier.NewStatusChangedEvent("Hello world", "Hello world", notifier.SubtokenStatusChanged))
}
