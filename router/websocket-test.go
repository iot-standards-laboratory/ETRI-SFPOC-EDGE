package router

import (
	"etri-sfpoc-edge/notifier"
	"time"
)

func fire() {
	for i := 0; i < 10; i++ {
		box.Publish(notifier.NewStatusChangedEvent("Hello world", "Hello world", notifier.SubtokenStatusChanged))
		time.Sleep(time.Second * 2)
	}
}
