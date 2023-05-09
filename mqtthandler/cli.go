package mqtthandler

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const user = "etrimqtt"
const passwd = "fainal2311"

var client mqtt.Client

func mqttHandleFunc(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(msg)
}

func ConnectMQTT(mqttAddr string) error {
	opts := mqtt.NewClientOptions().AddBroker(mqttAddr).SetClientID("edge-master-debug")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(mqttHandleFunc)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetUsername(user)
	opts.SetPassword(passwd)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func Publish(topic string, payload []byte) error {
	tkn := client.Publish(topic, 0, false, payload)
	return tkn.Error()
}

func Subscribe(topic string) error {
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Unsubscribe(topic string) {
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
