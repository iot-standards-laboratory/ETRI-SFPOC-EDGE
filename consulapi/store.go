package consulapi

import (
	"errors"

	"github.com/hashicorp/consul/api"
)

func Put(key string, value []byte) error {
	_, err := client.KV().Put(&api.KVPair{
		Key:   key,
		Value: value,
	}, nil)

	return err
}

func GetKeys(prefix string) ([]string, error) {
	list, _, err := client.KV().Keys(prefix, "", nil)
	return list, err
}

func Get(key string) ([]byte, error) {
	kvPair, _, err := client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	if kvPair == nil {
		return nil, errors.New("empty entity error")
	}

	return kvPair.Value, nil
}

func Delete(key string) error {
	_, err := client.KV().Delete(key, nil)
	return err
}
