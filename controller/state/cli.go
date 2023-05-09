package state

import (
	"sync"
)

var currentState IState = STATE_RUNNING
var mutex sync.Mutex
var currentStateStream = make(chan IState, 1)

type IState interface {
	String() string
}

func Subscribe() <-chan IState {
	return currentStateStream
}

func Put(sv IState) {
	mutex.Lock()
	defer mutex.Unlock()
	currentState = sv
	currentStateStream <- currentState
}

func Get() IState {
	mutex.Lock()
	defer mutex.Unlock()
	return currentState
}
