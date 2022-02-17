package router

import (
	"etri-sfpoc-edge/notifier"

	"github.com/gorilla/websocket"
)

type WebsocketSubscriber struct {
	notifier.INotiManager
	_id       string
	_token    string
	_type     int
	_complete chan<- int
	_conn     *websocket.Conn
	// _h     func(msg string)
}

func (rs *WebsocketSubscriber) ID() string {
	return rs._id
}

func (rs *WebsocketSubscriber) Token() string {
	return rs._token
}

func (rs *WebsocketSubscriber) Handle(e notifier.IEvent) {
	payload := map[string]interface{}{"key": e.Title()}

	if rs._conn.WriteJSON(payload) != nil {
		rs._complete <- 200
	}
}

func (rs *WebsocketSubscriber) Type() int {
	return rs._type
}
