package model

type IDBHandler interface {
	GetDevices(string) ([]*Device, int, error)
}
