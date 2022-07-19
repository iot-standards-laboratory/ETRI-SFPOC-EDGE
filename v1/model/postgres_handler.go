package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlHandler(path string) (DBHandlerI, error) {
	dsn := "host=localhost user=user password=user_password dbname=godopudb port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Device{})
	db.AutoMigrate(&Controller{})
	db.AutoMigrate(&Service{})

	return &_DBHandler{db}, nil
}
