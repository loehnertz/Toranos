package main

import (
	"encoding/json"
	"github.com/loehnertz/toranos/commons"
	"github.com/micro/go-log"
	"io"
	"net/http"
)

func deserialize(target interface{}, body io.ReadCloser) (deserialized interface{}, err error) {
	decoder := json.NewDecoder(body)
	decodeError := decoder.Decode(&target)
	if decodeError != nil {
		log.Log(decodeError)
		return nil, commons.UnknownError
	} else {
		return target, nil
	}
}

func respondWithJson(w *http.ResponseWriter, r interface{}) {
	jsonBytes, marshalError := json.Marshal(r)
	if marshalError != nil {
		log.Log(marshalError)
		(*w).Write([]byte(commons.UnknownError.Error()))
	}

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(jsonBytes)
}
