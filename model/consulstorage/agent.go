package consulstorage

import (
	"encoding/json"
	"etri-sfpoc-edge/consulapi"
	"etri-sfpoc-edge/model"
	"fmt"
	"io"

	"github.com/google/uuid"
)

func _getAgentKey(id string) string {
	return fmt.Sprintf("agents/metadata/%s", id)
}

func (s *_consulStorage) AddAgentWithJsonReader(r io.Reader) (*model.Agent, error) {
	dec := json.NewDecoder(r)
	var agent = &model.Agent{}
	err := dec.Decode(&agent)
	if err != nil {
		return nil, err
	}

	agent.ID = uuid.NewString()

	b, err := json.Marshal(agent)
	if err != nil {
		return nil, err
	}

	err = consulapi.Put(_getAgentKey(agent.ID), b)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (s *_consulStorage) GetAgent(id string) (*model.Agent, error) {
	b, err := consulapi.Get(_getAgentKey(id))
	if err != nil {
		return nil, err
	}

	var agent model.Agent
	err = json.Unmarshal(b, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

func (s *_consulStorage) DeleteAgent(id string) error {
	return consulapi.Delete(_getAgentKey(id))
}

func (s *_consulStorage) GetAgents() ([]*model.Agent, error) {
	pairs, err := consulapi.GetPairs("agents/metadata")
	if err != nil {
		return nil, err
	}

	list := make([]*model.Agent, 0, len(pairs))
	for _, pair := range pairs {
		agent := model.Agent{}
		err := json.Unmarshal(pair.Value, &agent)
		if err != nil {
			continue
		}

		list = append(list, &agent)
	}

	// result := s.db.Find(&list)

	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	return list, nil
}

// func (s *_DBHandler) IsExistAgent(id string) bool {
// 	var controller = model.Controller{}

// 	result := s.db.First(&controller, "cid=?", cid)

// 	return result.Error == nil
// }
