package cache

import (
	"errors"
	"etri-sfpoc-edge/model"
	"strings"
	"sync"
)

var ctrls = make([]*model.Controller, 0, 10)
var ctrlsMutex sync.Mutex

func AddCtrls(ctrl *model.Controller) error {
	// set mutex
	ctrlsMutex.Lock()
	defer ctrlsMutex.Unlock()

	// check whether the ctrl with cid exist
	for _, e := range ctrls {
		if strings.Compare(e.CID, ctrl.CID) == 0 {
			return errors.New("already exist controller")
		}
	}

	// add ctrl to table
	ctrls = append(ctrls, ctrl)

	return nil
}

func GetCtrls() []*model.Controller {
	return ctrls
}

func RemoveCtrl(cid string) {
	// set mutex
	ctrlsMutex.Lock()
	defer ctrlsMutex.Unlock()

	// delete controller
	for i, e := range ctrls {
		if strings.Compare(e.CID, cid) == 0 {
			ctrls[i] = ctrls[len(ctrls)-1]
			ctrls = ctrls[:len(ctrls)-1]
		}
	}
	// delete device related to the controller
	removeDevs(cid)
}
