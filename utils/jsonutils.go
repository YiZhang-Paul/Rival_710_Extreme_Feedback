package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ParseJSON parses json string into a map
func ParseJSON(input io.Reader) (map[string]interface{}, bool) {
	parsed := make(map[string]interface{})
	content, err := ioutil.ReadAll(input)
	if err != nil || json.Unmarshal(content, &parsed) != nil {
		return parsed, false
	}
	return parsed, true
}
