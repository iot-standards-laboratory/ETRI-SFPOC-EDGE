package consulstorage

import "etri-sfpoc-edge/model"

var DefaultDB model.I_DBHandler = nil

type _consulStorage struct{}

func init() {
	DefaultDB = &_consulStorage{}
}
