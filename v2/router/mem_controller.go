package router

import (
	"etri-sfpoc-edge/v2/model"
	"etri-sfpoc-edge/v2/model/dbstorage"
)

var db model.I_DBHandler = nil

func init() {
	var err error
	db, err = dbstorage.NewPostgresqlHandler("localhost", "postgres", "postgres", "postgres")
	if err != nil {
		panic(err)
	}
}
