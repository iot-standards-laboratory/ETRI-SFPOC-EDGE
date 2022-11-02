package notimanager_test

import (
	"etri-sfpoc-edge/v2/notimanager"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNotiManager(t *testing.T) {
	assert := assert.New(t)

	manager := notimanager.NewNotiManager()

	id := uuid.New().String()
	token, err := notimanager.GenerateSecureToken(16)
	channel := make(chan notimanager.IEvent)
	assert.NoError(err)

	manager.AddSubscriber(notimanager.NewChanSubscriber(id, token, notimanager.SubtypeCont, channel))
	manager.AddSubscriber(notimanager.NewChanSubscriber(id, token, notimanager.SubtypeOnce, channel))

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second)
			manager.Publish(notimanager.NewSimpleEvent(token, "Hello world"))
		}
	}()
	e := <-channel

	assert.Equal(e.Message(), "Hello world")

	af := time.After(time.Second * 3)

	// subscriber type check
	select {
	case <-af:
		t.Fatal("timeout")
	case <-channel:
		assert.Equal(e.Message(), "Hello world")
	}
}
