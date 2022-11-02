package config

import (
	"io"
	"os"
	"strings"
)

var compose = `version: '3'

services:
  consul:
    command: agent -ui -bootstrap -server -client 0.0.0.0
    container_name: consul
    hostname: consul
    image: consul:latest
    # networks:
    #   edgex-network: {}
    ports:
      - 9999:8500/tcp
    read_only: true
    restart: always
    security_opt:
      - no-new-privileges:true
    user: root:root
    volumes:
      - ./config:/consul/config:z
    #   - consul-data:/consul/data:z
`

func CreateCompose() error {
	f, err := os.Create("docker-compose.yaml")
	if err != nil {
		return err
	}

	io.Copy(f, strings.NewReader(compose))
	return nil
}
