package model

import "github.com/google/uuid"

type I_DBHandler interface {
	InsertCtrl(name string) (*Controller, error)

	SelectSvc(id uuid.UUID) (*Service, error)
	UpdateSvc(svc *Service) (*Service, error)

	SelectImg(id uuid.UUID) (*SG_Image, error)
	// GetDevices(string) ([]*Device, int, error)
	// AddAgentWithJsonReader(r io.Reader) (*Agent, error)
	// GetAgent(id string) (*Agent, error)
	// GetAgents() ([]*Agent, error)
	// DeleteAgent(id string) error
	// UpdateAgentWithJsonReader(id string, body io.ReadCloser) error
	// IsExistController(cid string) bool
}
