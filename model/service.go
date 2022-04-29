package model

import (
	"errors"
	"etrisfpocdatamodel"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service etrisfpocdatamodel.Service

func (s *_DBHandler) GetServices() ([]*Service, error) {
	var services []*Service

	result := s.db.Find(&services)

	if result.Error != nil {
		return nil, result.Error
	}

	// var devices []*Device
	for _, service := range services {
		tx := s.db.Where("sname=?", service.SName).Find(&[]*Device{})
		if tx.Error == gorm.ErrRecordNotFound {
			service.NumOfDevs = 0
		} else if tx.Error != nil {
			return nil, tx.Error
		} else {
			service.NumOfDevs = int(tx.RowsAffected)
		}
	}
	return services, nil
}

func (s *_DBHandler) AddService(name string) error {
	result := s.db.First(&Service{}, "sname=?", name)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			tx := s.db.Create(&Service{SName: name})
			if tx.Error != nil {
				return tx.Error
			}
		}

		return result.Error
	}

	return nil
}

func (s *_DBHandler) IsExistService(name string) bool {
	var service Service
	tx := s.db.Select("sid").First(&service, "sname=?", name)

	return tx.Error == nil
}

func (s *_DBHandler) RegisterService(sname, addr string) (*Service, error) {
	tx := s.db.Model(&Service{}).Where("sname = ?", sname).Updates(Service{SName: sname, SID: uuid.NewString(), Addr: addr})
	if tx.Error != nil {
		return nil, tx.Error
	}

	var service Service
	tx.First(&service, "sname=?", sname)
	return &service, nil
}

func (s *_DBHandler) UpdateService(sid, addr string) (*Service, error) {
	tx := s.db.Model(&Service{}).Where("sid = ?", sid).Updates(Service{Addr: addr})
	if tx.Error != nil {
		return nil, tx.Error
	}

	var service Service
	tx.First(&service, "sid=?", sid)
	return &service, nil
}

func (s *_DBHandler) GetSID(name string) (string, error) {
	var service Service
	tx := s.db.Select("sid").First(&service, "sname=?", name)
	if tx.Error != nil {
		return "", tx.Error
	}

	if len(service.SID) == 0 {
		return "", errors.New("not installed service")
	}
	return service.SID, nil
}

func (s *_DBHandler) GetAddr(sid string) (string, error) {
	var service Service

	result := s.db.First(&service, "sid=?", sid)

	if result.Error != nil {
		return "", result.Error
	}

	return service.Addr, nil
}
