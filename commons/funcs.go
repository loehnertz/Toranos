package commons

import (
	"encoding/json"
	"github.com/micro/go-log"
)

func SliceOfStringsContains(s []string, e string) bool {
	for i := range s {
		if s[i] == e {
			return true
		}
	}
	return false
}

func StringifyIntoJson(structure interface{}) (stringifiedJson string) {
	jsonBytes, jsonMarshalError := json.Marshal(structure)
	if jsonMarshalError != nil {
		log.Log(jsonMarshalError)
		return
	}
	return string(jsonBytes[:])
}
