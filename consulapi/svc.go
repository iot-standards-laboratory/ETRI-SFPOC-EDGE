package consulapi

import "fmt"

func GetSvcAddr(svcName string) (string, error) {
	svc, _, err := client.Agent().Service(svcName, nil)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", svc.Address, svc.Port), nil
}
