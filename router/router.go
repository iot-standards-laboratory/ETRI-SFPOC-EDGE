package router

import (
	"etri-sfpoc-edge/notifier"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type RequestBox struct {
	notifier.INotiManager
}

var box *RequestBox

func init() {
	box = &RequestBox{notifier.NewNotiManager()}
}

func fire() {
	for i := 0; i < 5; i++ {
		box.Publish(notifier.NewStatusChangedEvent("Hello world", "Hello world", notifier.SubtokenStatusChanged))
		time.Sleep(time.Second * 2)
	}
}

func NewRouter() *gin.Engine {
	apiEngine := gin.New()
	apiv1 := apiEngine.Group("api/v1")
	{
		apiv1.GET("/notification", WsTest)
		apiv1.GET("/test", func(c *gin.Context) {
			c.String(200, "Hello world")
		})
	}

	r := gin.New()

	assetEngine := gin.New()
	assetEngine.Static("/", "./static")
	r.GET("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiEngine.HandleContext(c)
		} else {
			assetEngine.HandleContext(c)
		}
	})

	return r
}

// Alarm
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WsTest(c *gin.Context) {
	_complete := make(chan int)
	_uuid, _ := uuid.NewV4()

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	// conn.SetPingHandler((func(ping string) error {
	// 	_h := conn.PingHandler()
	// 	fmt.Println("ping : ", ping)
	// 	return _h(ping)
	// }))

	// conn.SetPongHandler((func(pong string) error {
	// 	_h := conn.PongHandler()
	// 	fmt.Println("pong : ", pong)
	// 	return _h(pong)
	// }))

	// conn.SetCloseHandler((func(code int, text string) error {
	// 	_h := conn.CloseHandler()
	// 	fmt.Println("Close!!")
	// 	return _h(code, text)
	// }))

	subscriber := &WebsocketSubscriber{
		_id:       _uuid.String(),
		_token:    notifier.SubtokenStatusChanged,
		_type:     notifier.SubtypeCont,
		_complete: _complete,
		_conn:     conn,
	}
	box.AddSubscriber(subscriber)
	defer box.RemoveSubscriber(subscriber)

	<-_complete

	// c.Writer.Write([]byte(e.Title()))
}
