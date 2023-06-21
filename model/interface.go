package model

import "io"

type I_DBHandler interface {
	// GetDevices(string) ([]*Device, int, error)
	AddAgentWithJsonReader(r io.Reader) (*Agent, error)
	GetAgent(id string) (*Agent, error)
	GetAgents() ([]*Agent, error)
	DeleteAgent(id string) error
	UpdateAgentWithJsonReader(id string, body io.ReadCloser) error
	// IsExistController(cid string) bool
}
