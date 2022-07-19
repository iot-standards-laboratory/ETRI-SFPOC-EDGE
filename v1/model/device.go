package model

import (
	"etrisfpocdatamodel"
	"sync"

	"gorm.io/gorm"
)

type Device etrisfpocdatamodel.Device

var discoveredDevices = []*Device{}
var discoveredDeviceMutex sync.Mutex

func (s *_DBHandler) AddDiscoveredDevice(device *Device) {
	discoveredDeviceMutex.Lock()
	defer discoveredDeviceMutex.Unlock()
	discoveredDevices = append(discoveredDevices, device)
}

func (s *_DBHandler) GetDiscoveredDevices() []*Device {
	return discoveredDevices
}

func (s *_DBHandler) RemoveDiscoveredDevice(device *Device) {
	discoveredDeviceMutex.Lock()
	defer discoveredDeviceMutex.Unlock()
	for i, e := range discoveredDevices {
		if e == device {
			discoveredDevices[i] = discoveredDevices[len(discoveredDevices)-1]
			discoveredDevices = discoveredDevices[:len(discoveredDevices)-1]
		}
	}
}

func (s *_DBHandler) GetDevices(sname string) ([]*Device, int, error) {
	var devices []*Device

	var result *gorm.DB
	if sname == "all" {
		result = s.db.Find(&devices)
	} else {
		result = s.db.Where("sname=?", sname).Find(&devices)
	}

	if result.Error != nil {
		return nil, -1, result.Error
	}
	return devices, int(result.RowsAffected), nil
}

func (s *_DBHandler) AddDevice(device *Device) error {

	tx := s.db.Create(device)
	if tx.Error != nil {
		return tx.Error
	}

	tx.First(device, "did=?", device.DID)
	return nil

}

func (s *_DBHandler) QueryDevice(did string) (*Device, error) {
	var device Device
	tx := s.db.First(&device, "did=?", did)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &device, nil
}

func (s *_DBHandler) DeleteDevice(device *Device) error {
	tx := s.db.Delete(device)
	if tx.Error != nil {
		return tx.Error
	}

	// tx.First(device, "did=?", device.DID)
	return nil
}

func (s *_DBHandler) IsExistDevice(dname string) bool {
	var device = Device{}

	result := s.db.First(&device, "dname=?", dname)

	return result.Error != nil
}
