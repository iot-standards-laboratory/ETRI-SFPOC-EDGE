package main

import (
	"etri-sfpoc-edge/canary-example/centrifuge/client/centrifuge_api"
	"log"

	"github.com/centrifugal/centrifuge-go"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New().String()
	token, _ := centrifuge_api.IssueJWT(id, "godopu", "ctrl", "/", nil)
	client := centrifuge_api.NewClient("ws://dsad:8000/connection/websocket", token)
	defer client.Close()
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	sub, err := centrifuge_api.NewSubscription(client, "public:blank")
	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		log.Printf("Someone says via channel %s: %s (offset %d)", sub.Channel, e.Data, e.Offset)
	})
	if err != nil {
		panic(err)
	}

	err = sub.Subscribe()
	if err != nil {
		panic(err)
	}

	select {}
}
