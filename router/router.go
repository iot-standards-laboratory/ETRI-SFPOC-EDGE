package router

import (
	"etri-sfpoc-edge/model"
	"etri-sfpoc-edge/notifier"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type RequestBox struct {
	notifier.INotiManager
}

var box *RequestBox

var db model.DBHandlerI

func init() {
	box = &RequestBox{notifier.NewNotiManager()}

	var err error
	db, err = model.NewSqliteHandler("dump.db")
	if err != nil {
		panic(err)
	}
	// go fire()
}

func NewRouter() *gin.Engine {
	apiEngine := gin.New()
	apiv1 := apiEngine.Group("api/v1")
	{
		apiv1.GET("/subs", Subscribe)
		apiv1.POST("/ctrls", PostCtrl)
		apiv1.GET("/ctrls/list", GetCtrlList)
	}

	r := gin.New()

	assetEngine := gin.New()
	assetEngine.Static("/", "./static")
	r.Any("/*any", func(c *gin.Context) {
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

func Subscribe(c *gin.Context) {
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

	subscriber := notifier.NewWebsocketSubscriber(_uuid.String(), notifier.SubtokenStatusChanged, notifier.SubtypeCont, _complete, conn)
	box.AddSubscriber(subscriber)
	defer box.RemoveSubscriber(subscriber)

	<-_complete
}
