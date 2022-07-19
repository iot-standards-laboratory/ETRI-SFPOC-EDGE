package model

import (
	"io"

	"gorm.io/gorm"
)

// func newDBHandler(dbtype, path string) (*gorm.DB, error) {
// 	if dbtype == "sqlite" {
// 		return gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
// 	} else {
// 		dsn := "host=localhost user=user password=user_password dbname=godopudb port=5432 sslmode=disable TimeZone=Asia/Seoul"
// 		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	}
// }

type DBHandlerI interface {
	GetDevices(string) ([]*Device, int, error)
	AddDevice(d *Device) error
	AddDiscoveredDevice(device *Device)
	GetDiscoveredDevices() []*Device
	RemoveDiscoveredDevice(device *Device)
	QueryDevice(dname string) (*Device, error)
	DeleteDevice(device *Device) error
	IsExistDevice(dname string) bool
	AddController(r io.Reader) (*Controller, error)
	GetController(cid string) (*Controller, error)
	GetControllers() ([]*Controller, error)
	IsExistController(cid string) bool
	GetServices() ([]*Service, error)
	AddService(name string) error
	RegisterService(sname, addr string) (*Service, error)
	UpdateService(sid, addr string) (*Service, error)
	GetAddr(sid string) (string, error)
	GetSID(name string) (string, error)
	IsExistService(name string) bool
}

type _DBHandler struct {
	db *gorm.DB
}
