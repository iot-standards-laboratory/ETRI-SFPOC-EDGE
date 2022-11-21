package dbstorage

import "etri-sfpoc-edge/v2/model"

var DefaultDB model.I_DBHandler = nil

func init() {
	var err error
	DefaultDB, err = NewPostgresqlHandler("localhost", "postgres", "postgres", "postgres")
	if err != nil {
		panic(err)
	}
}
