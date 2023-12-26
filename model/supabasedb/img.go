package supabasedb

import (
	"encoding/json"
	"errors"
	"etri-sfpoc-edge/model"

	"github.com/google/uuid"
)

func (db *_supabaseDB) SelectImg(id uuid.UUID) (*model.SG_Image, error) {
	d, _, err := client.From("etri_list_imgs").Select("*", "", false).Eq("id", id.String()).Execute()
	if err != nil {
		return nil, err
	}

	return parseImg(d)
}

func parseImg(d []byte) (*model.SG_Image, error) {
	img := make([]model.SG_Image, 1)
	err := json.Unmarshal(d, &img)
	if err != nil {
		return nil, err
	}

	if len(img) != 1 {
		return nil, errors.New("does not exist error")
	}

	return &img[0], err
}
