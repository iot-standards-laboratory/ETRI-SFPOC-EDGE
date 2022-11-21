package dbstorage

import (
	"etri-sfpoc-edge/v2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type _DBHandler struct {
	db *gorm.DB
}

func NewPostgresqlHandler(endpoint, user, pwd, database string) (model.I_DBHandler, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=Asia/Seoul", endpoint, user, pwd, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Controller{})
	db.AutoMigrate(&model.Agent{})
	// db.AutoMigrate(&model.Device{})
	// db.AutoMigrate(&model.Service{})

	return &_DBHandler{db}, nil
}
