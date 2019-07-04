package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
