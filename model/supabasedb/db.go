package supabasedb

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/model"
	"fmt"

	"github.com/google/uuid"
)

type _supabaseDB struct{}

func (db *_supabaseDB) InsertCtrl(name string) (*model.Controller, error) {
	d, _, err := client.From("etri_list_ctrls").Insert(map[string]interface{}{
		"name":   name,
		"status": true,
	}, true, "", "representation", "").Execute()
	if err != nil {
		return nil, err
	}

	var ctrl []model.Controller
	err = json.Unmarshal(d, &ctrl)

	return &ctrl[0], err
}

func (db *_supabaseDB) SelectSvc(id uuid.UUID) (*model.Service, error) {
	fmt.Println("Select:", id)
	d, _, err := client.From("etri_list_svcs").Select("*", "", false).Eq("id", id.String()).Execute()
	if err != nil {
		return nil, err
	}

	return parseSvc(d)
}

func (db *_supabaseDB) UpdateSvc(svc *model.Service) (*model.Service, error) {
	d, _, err := client.From("etri_list_svcs").Update(svc, "representation", "").Eq("id", svc.ID.String()).Execute()
	if err != nil {
		return nil, err
	}

	return parseSvc(d)
}

func parseSvc(d []byte) (*model.Service, error) {
	svc := make([]model.Service, 1)
	err := json.Unmarshal(d, &svc)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(svc))
	if len(svc) != 1 {
		return nil, errors.New("does not exist error")
	}

	return &svc[0], err
}
