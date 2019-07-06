package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func subscribe(target, callback string) (string, error) {
	var (
		payload   = map[string]interface{}{"callbackUrl": callback}
		data, err = json.Marshal(payload)
	)
	if err != nil {
		return "", err
	}
	res, err := http.Post(target, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if id, err := ioutil.ReadAll(res.Body); err != nil {
		return "", err
	} else {
		return string(id), nil
	}
}
