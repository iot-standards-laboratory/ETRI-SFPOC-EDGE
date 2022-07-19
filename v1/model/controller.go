package model

import (
	"encoding/json"
	"etrisfpocdatamodel"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type Controller etrisfpocdatamodel.Controller

func (s *_DBHandler) AddController(r io.Reader) (*Controller, error) {
	decoder := json.NewDecoder(r)
	var controller = &Controller{}

	err := decoder.Decode(controller)
	if err != nil {
		return nil, err
	}

	controller.CID = uuid.NewString()
	controller.Key = controller.CID

	tx := s.db.Create(controller)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx.First(&controller, "cid=?", controller.CID)
	fmt.Println(controller)
	return controller, nil
}

func (s *_DBHandler) GetController(cid string) (*Controller, error) {
	var ctrl Controller
	result := s.db.First(&ctrl, "cid=?", cid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &ctrl, nil
}

func (s *_DBHandler) GetControllers() ([]*Controller, error) {
	var list []*Controller

	result := s.db.Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

func (s *_DBHandler) IsExistController(cid string) bool {
	var controller = Controller{}

	result := s.db.First(&controller, "cid=?", cid)

	return result.Error == nil
}
