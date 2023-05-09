package connmgmt

import (
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/mqtthandler"
)

var connected = false

func Connect(mqttAddr, consulAddr string) error {
	if connected {
		return nil
	}
	err := mqtthandler.ConnectMQTT(mqttAddr)
	if err != nil {
		return err
	}

	err = consulapi.Connect(consulAddr)
	if err != nil {
		return err
	}

	connected = true
	return nil
}
