package dbstorage

import (
	"encoding/json"
	"etri-sfpoc-edge/v2/model"
	"fmt"
	"io"

	"github.com/google/uuid"
)

func (s *_DBHandler) AddAgentWithJsonReader(r io.Reader) (*model.Agent, error) {
	decoder := json.NewDecoder(r)
	var agent = &model.Agent{}

	err := decoder.Decode(agent)
	if err != nil {
		return nil, err
	}

	agent.ID = uuid.NewString()
	// controller.Key = controller.CID

	tx := s.db.Create(agent)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx.First(&agent, "id=?", agent.ID)
	fmt.Println(agent)
	return agent, nil
}

func (s *_DBHandler) GetAgent(id string) (*model.Agent, error) {
	var agent model.Agent
	result := s.db.First(&agent, "id=?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &agent, nil
}

func (s *_DBHandler) DeleteAgent(id string) error {
	result := s.db.Delete(&model.Agent{ID: id})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *_DBHandler) GetAgents() ([]*model.Agent, error) {
	var list []*model.Agent

	result := s.db.Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

// func (s *_DBHandler) IsExistAgent(id string) bool {
// 	var controller = model.Controller{}

// 	result := s.db.First(&controller, "cid=?", cid)

// 	return result.Error == nil
// }
