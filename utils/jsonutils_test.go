package utils

import (
	"bytes"
	"io/ioutil"
	"testing"
)

const testJSON = "../testdata/testjson.json"

func TestParseJSON(t *testing.T) {
	file, err := ioutil.ReadFile(testJSON)
	if err != nil {
		t.Error("Failed to read test data.")
		return
	}
	parsed, ok := ParseJSON(bytes.NewBuffer(file))
	if !ok || parsed["a"] != float64(1) || parsed["b"] != "some text" {
		t.Errorf("ParseJSON(%v) does not return expected JSON.", testJSON)
	}
}

func TestParseJSONFile(t *testing.T) {
	cases := []struct {
		path       string
		shouldFail bool
	}{
		{"invalid_path/invalid.json", true},
		{testJSON, false},
	}
	for _, c := range cases {
		parsed, ok := ParseJSONFile(c.path)
		if c.shouldFail {
			if ok {
				t.Errorf("ParseJSONFile(%v) should fail instead.", c.path)
			}
			continue
		}
		if parsed["a"] != float64(1) || parsed["b"] != "some text" {
			t.Errorf("ParseJSONFile(%v) does not return expected JSON.", c.path)
		}
	}
}
