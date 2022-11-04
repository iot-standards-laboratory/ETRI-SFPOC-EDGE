package consulapi_test

import (
	"etri-sfpoc-edge/v2/consulapi"
	"etri-sfpoc-edge/v2/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestControllerRegistration(t *testing.T) {
	assert := assert.New(t)
	// cid := uuid.New().String()
	cid := "ctrl/e7e0492c-183c-4504-8561-b8ca9eb4b6d3"
	ctrl := model.Controller{
		CID: cid,
	}

	err := consulapi.Connect("http://localhost:9999")
	assert.NoError(err)
	err = consulapi.RegisterCtrl(ctrl, "http://localhost:8080")
	assert.NoError(err)

	go consulapi.UpdateTTL(func() (bool, error) { return true, nil }, cid)

	go func() {
		time.Sleep(time.Second * 2)
		for {
			state, err := consulapi.GetStatus(cid)
			assert.NoError(err)
			assert.Equal("passing", state)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 20)
	consulapi.DeregisterCtrl(cid)
}

func TestStore(t *testing.T) {
	assert := assert.New(t)

	err := consulapi.Connect("http://localhost:9999")
	assert.NoError(err)

	key := "ctrl/devs"
	value := "Hello world"

	err = consulapi.Put(key, []byte(value))
	assert.NoError(err)

	v, err := consulapi.Get(key)
	assert.NoError(err)
	assert.Equal(string(v), value)
}
