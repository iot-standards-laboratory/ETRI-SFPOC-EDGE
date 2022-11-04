package consulapi

import (
	"etri-sfpoc-edge/v2/model"
	"time"

	"github.com/golang/glog"
	"github.com/hashicorp/consul/api"
)

const ttl = time.Duration(time.Second * 5)

func RegisterCtrl(ctrl model.Controller, endpoint string) error {
	err := registerEntity(ctrl.CID, endpoint)
	if err != nil {
		panic(err)
	}

	return nil
}

func DeregisterCtrl(name string) {
	client.Agent().ServiceDeregister(name)
	glog.Infof("[ctrl %v] - deregistered.", name)
}

func UpdateTTL(check func() (bool, error), name string) {
	agent := client.Agent()
	update(check, agent, name)
	ticker := time.NewTicker(ttl / 2)

	i := 0
	for range ticker.C {
		update(check, agent, name)
		i++
		if i > 3 {
			return
		}
	}
}

func update(check func() (bool, error), agent *api.Agent, name string) {
	ok, err := check()
	if !ok {
		glog.Errorf("err=\"Check failed\" msg=\"%s\"", err.Error())
		if agentErr := agent.FailTTL("service:"+name, err.Error()); agentErr != nil {
			glog.Error(agentErr)
		}
	} else {
		if agentErr := agent.PassTTL("service:"+name, "healthy"); agentErr != nil {
			glog.Error(agentErr)
		}
	}
}
