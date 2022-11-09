package model

import "io"

type I_DBHandler interface {
	// GetDevices(string) ([]*Device, int, error)
	AddControllerWithJsonReader(r io.Reader) (*Controller, error)
	GetController(cid string) (*Controller, error)
	GetControllers() ([]*Controller, error)
	// IsExistController(cid string) bool
}
