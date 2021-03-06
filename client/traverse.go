package client

import (
	"github.com/fishi0x01/vsh/log"
)

func (client *Client) topLevelTraverse(path string) (result []string) {
	for k := range client.KVBackends {
		result = append(result, k)
	}

	return result
}

func (client *Client) lowLevelTraverse(path string) (result []string) {
	s, err := client.Vault.Logical().List(client.getKVMetaDataPath(path))
	if err != nil {
		log.Error("Error traversing path: %v", err)
		return nil
	}

	if s != nil {
		if keysInterface, ok := s.Data["keys"]; ok {
			for _, valInterface := range keysInterface.([]interface{}) {
				val := valInterface.(string)
				result = append(result, client.lowLevelTraverse(path+"/"+val)...)
			}
		}
	} else {
		result = append(result, "/"+path)
	}

	return result
}
