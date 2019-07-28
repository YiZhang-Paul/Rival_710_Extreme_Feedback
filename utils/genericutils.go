package utils

import (
	"log"
	"math"
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

// StringPrefixSum returns a list of prefix sum for string type.
// e.g, "text" -> ["", "t", "te", "tex", "text"]
func StringPrefixSum(text string) []string {
	sum := make([]string, 0)
	for i := range text {
		sum = append(sum, text[0:i])
	}
	return append(sum, text)
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

// MultiplyUint8 multiplies uint8 inputs without wrapping
func MultiplyUint8(a, b uint8) uint8 {
	if a == 0 || b == 0 {
		return 0
	}
	if math.MaxUint8/a >= b {
		return a * b
	}
	return math.MaxUint8
}

// MinusUint8 minuses uint8 inputs without wrapping
func MinusUint8(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return 0
}

// ParseToStrings will try to parse dynamic data into a list of strings.
// It assumes the underlying data to be a list of float64
func ParseToStrings(data interface{}) []string {
	var (
		result     = make([]string, 0)
		values, ok = data.([]interface{})
	)
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

// StringsToUint8s maps a slice of strings into a slice of uint8s
func StringsToUint8s(data []string) []uint8 {
	result := make([]uint8, len(data))
	for i, value := range data {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			result[i] = 0
		} else {
			result[i] = uint8(parsed)
		}
	}
	return result
}
