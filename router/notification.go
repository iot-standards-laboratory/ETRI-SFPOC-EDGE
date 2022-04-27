package router

import (
	"etri-sfpoc-edge/logger"
	"etri-sfpoc-edge/notifier"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func GetPublish(c *gin.Context) {
	_complete := make(chan int)
	_uuid, _ := uuid.NewV4()

	path := c.Param("any")

	var subtoken string
	if len(path) <= 9 {
		subtoken = notifier.SubtokenStatusChanged
	} else if path[8] != '/' {
		return
	} else {
		subtoken = path[9:]
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	subscriber := notifier.NewWebsocketSubscriber(
		_uuid.String(),
		subtoken,
		notifier.SubtypeCont,
		_complete,
		conn,
	)

	box.AddSubscriber(subscriber)
	defer box.RemoveSubscriber(subscriber)

	closeCh := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				closeCh <- true
			}
		}()
		for {
			// Read Messages
			_, _, err := conn.ReadMessage()

			if c, k := err.(*websocket.CloseError); k {
				if c.Code == 1000 {
					// Never entering since c.Code == 1005
					logger.Println(err)
					panic(err)
				}
			}
		}
	}()

	select {
	case <-closeCh:
		box.RemoveSubscriber(subscriber)
		return
	case <-_complete:
		return
	}
}

func PostPublish(c *gin.Context) {
	defer handleError(c)

	path := c.Param("any")

	var subtoken string
	if len(path) <= 9 {
		subtoken = notifier.SubtokenStatusChanged
	} else if path[8] != '/' {
		return
	} else {
		subtoken = path[9:]
	}

	fmt.Printf("subtoken:%s\n", subtoken)

	var body map[string]interface{}
	err := c.BindJSON(&body)
	if err != nil {
		panic(err)
	}

	box.Publish(notifier.NewPushEvent("control", body, subtoken))
	c.String(http.StatusOK, "Sended PUSH")
}

func GetSubscriberList(c *gin.Context) {
	defer handleError(c)
	c.JSON(http.StatusOK, box.GetSubscriberList())
}

// func fire() {
// 	for i := 0; i < 10; i++ {
// 		box.Publish(notifier.NewStatusChangedEvent("Hello world", "Hello world", notifier.SubtokenStatusChanged))
// 		time.Sleep(time.Second * 2)
// 	}
// }
