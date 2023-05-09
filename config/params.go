package config

import (
	"os"

	"github.com/magiconair/properties"
)

const Mode = "debug"

var Params = map[string]interface{}{}

func LoadConfig() {

	p := properties.MustLoadFile("./config.properties", properties.UTF8)
	Params["bind"] = p.GetString("bind", ":3000")
	Params["consulAddr"] = p.GetString("consulAddr", "http://localhost:9999")
	Params["mqttAddr"] = p.GetString("mqttAddr", "ws://localhost:9998")
}

func CreateInitFile() {
	f, err := os.Create("./config.properties")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := properties.NewProperties()

	for k, v := range Params {
		p.SetValue(k, v)
	}
	p.SetValue("bind", ":3000")

	p.Write(f, properties.UTF8)
}
