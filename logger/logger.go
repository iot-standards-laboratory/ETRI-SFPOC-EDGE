package logger

import (
	"etri-sfpoc-edge/config"
	"log"
)

func Println(v ...interface{}) {
	if config.Mode == "debug" {
		log.Println(v...)
	}
}
