package consulapi

import "github.com/hashicorp/consul/api"

func Put(key string, value []byte) error {
	_, err := client.KV().Put(&api.KVPair{
		Key:   key,
		Value: value,
	}, nil)

	return err
}

func Get(key string) ([]byte, error) {
	kvPair, _, err := client.KV().Get(key, nil)

	return kvPair.Value, err
}
