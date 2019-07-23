package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ParseJSON parses json string into a map
func ParseJSON(input io.Reader) (map[string]interface{}, bool) {
	content, err := ioutil.ReadAll(input)
	if err != nil {
		return make(map[string]interface{}), false
	}
	return tryParseJSON(content)
}

// ParseJSONFile takes full path of a json file and parse it into a map
func ParseJSONFile(path string) (map[string]interface{}, bool) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return make(map[string]interface{}), false
	}
	return tryParseJSON(file)
}

func tryParseJSON(data []byte) (map[string]interface{}, bool) {
	parsed := make(map[string]interface{})
	if json.Unmarshal(data, &parsed) != nil {
		return parsed, false
	}
	return parsed, true
}

// PostJSON wraps http.Post() to reduce boilerplate code.
// Only use this when POST response is not needed
func PostJSON(url string, data interface{}) bool {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return false
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	} else {
		defer response.Body.Close()
	}
	return err == nil
}
