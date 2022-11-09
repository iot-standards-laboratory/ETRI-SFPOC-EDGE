package dbstorage

import (
	"encoding/json"
	"etri-sfpoc-edge/v2/model"
	"fmt"
	"io"

	"github.com/google/uuid"
)

func (s *_DBHandler) AddControllerWithJsonReader(r io.Reader) (*model.Controller, error) {
	decoder := json.NewDecoder(r)
	var controller = &model.Controller{}

	err := decoder.Decode(controller)
	if err != nil {
		return nil, err
	}

	controller.CID = uuid.NewString()
	// controller.Key = controller.CID

	tx := s.db.Create(controller)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx.First(&controller, "cid=?", controller.CID)
	fmt.Println(controller)
	return controller, nil
}

func (s *_DBHandler) GetController(cid string) (*model.Controller, error) {
	var ctrl model.Controller
	result := s.db.First(&ctrl, "cid=?", cid)

	if result.Error != nil {
		return nil, result.Error
	}

	return &ctrl, nil
}

func (s *_DBHandler) GetControllers() ([]*model.Controller, error) {
	var list []*model.Controller

	result := s.db.Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

func (s *_DBHandler) IsExistController(cid string) bool {
	var controller = model.Controller{}

	result := s.db.First(&controller, "cid=?", cid)

	return result.Error == nil
}
