package utils

import (
	"log"
	"strconv"
)

// NextIndex finds next valid index in given collection.
// It prevents index out of bound by wrapping back to 0
func NextIndex(collection []string, index int) int {
	max := len(collection) - 1
	if index >= max {
		return 0
	}
	return index + 1
}

// TernaryString is a simple implementation of ternary operator.
// Only use this for trivial assignments as all expressions will be evaluated
func TernaryString(condition bool, ifTrue, ifFalse string) string {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// TernaryBool is a simple implementation of ternary operator.
// Only use this for trivial assignments as all expressions will be evaluated
func TernaryBool(condition, ifTrue, ifFalse bool) bool {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// TernaryInt is a simple implementation of ternary operator.
// Only use this for trivial assignments as all expressions will be evaluated
func TernaryInt(condition bool, ifTrue, ifFalse int) int {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// ParseToStrings will try to parse dynamic data into a list of strings.
// It assumes the underlying data to be a list of float64
func ParseToStrings(data interface{}) []string {
	result := make([]string, 0)
	values, ok := data.([]interface{})
	if !ok {
		log.Printf("Failed to parse %v into collection.\n", data)
		return result
	}
	for _, value := range values {
		if parsed, ok := value.(float64); !ok {
			log.Printf("Failed to parse %v into float64.\n", value)
		} else {
			result = append(result, strconv.Itoa(int(parsed)))
		}
	}
	return result
}

// FloatFromMap retrieves value from a map and parse it into float64.
// If the value does not exist or the value is not float64, 0 will be returned
func FloatFromMap(table map[string]interface{}, key string) (float64, bool) {
	if temp, ok := table[key]; !ok {
		return 0, false
	} else if value, ok := temp.(float64); !ok {
		return 0, false
	} else {
		return value, true
	}
}

// StringFromMap retrieves value from a map and parse it into string.
// If the value does not exist or the value is not string, empty string will be returned
func StringFromMap(table map[string]interface{}, key string) (string, bool) {
	if temp, ok := table[key]; !ok {
		return "", false
	} else if value, ok := temp.(string); !ok {
		return "", false
	} else {
		return value, true
	}
}
