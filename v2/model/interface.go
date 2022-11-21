package model

import "io"

type I_DBHandler interface {
	// GetDevices(string) ([]*Device, int, error)
	AddAgentWithJsonReader(r io.Reader) (*Agent, error)
	GetAgent(id string) (*Agent, error)
	GetAgents() ([]*Agent, error)
	// IsExistController(cid string) bool
}
