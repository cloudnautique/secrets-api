package service

import "encoding/json"

func MapStringInterfaceToString(dataMap map[string]interface{}) (string, error) {
	dataMapBytes, err := json.Marshal(dataMap)
	return string(dataMapBytes), err
}
