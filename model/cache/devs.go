package cache

import (
	"errors"
	"etri-sfpoc-edge/model"
	"strings"
	"sync"
)

var devs map[string][]*model.Device // devs[cid] = dev list
var devsMutex sync.Mutex

func AddDevs(dev *model.Device) error {
	devsMutex.Lock()
	defer devsMutex.Unlock()

	list, ok := devs[dev.CID]
	if !ok {
		list := make([]*model.Device, 0, 10)
		list = append(list, dev)
		devs[dev.CID] = list
		return nil
	}

	for _, e := range list {
		if strings.Compare(e.DID, dev.DID) == 0 {
			return errors.New("already exist device")
		}
	}

	devs[dev.CID] = append(list, dev)

	return nil
}

func removeDevs(cid string) {
	devsMutex.Lock()
	defer devsMutex.Unlock()

	delete(devs, cid)
}

func RemoveDev(dev *model.Device) {
	devsMutex.Lock()
	defer devsMutex.Unlock()

	list, ok := devs[dev.CID]
	if !ok {
		return
	}

	for i, e := range list {
		if strings.Compare(dev.DID, e.DID) == 0 {
			list[i] = list[len(list)-1]
			if len(list)-1 == 0 {
				removeDevs(e.CID)
			} else {
				devs[dev.CID] = list[:len(list)-1]
			}
		}
	}
}

func GetDevList() []*model.Device {
	var devList = make([]*model.Device, 0, 20)

	for _, v := range devs {
		devList = append(devList, v...)
	}

	return devList
}
