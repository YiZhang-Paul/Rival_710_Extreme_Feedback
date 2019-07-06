package utils

import "testing"

func TestNextIndex(t *testing.T) {
	var (
		texts = []string{"a", "b", "c"}
		cases = []struct {
			list            []string
			index, expected int
		}{
			{texts, 0, 1},
			{texts, 1, 2},
			{texts, 2, 0},
		}
	)
	for _, c := range cases {
		if actual := NextIndex(c.list, c.index); actual != c.expected {
			format := "NextIndex(%v, %d) == %d. Expected %d"
			t.Errorf(format, c.list, c.index, actual, c.expected)
		}
	}
}

func TestStringPrefixSum(t *testing.T) {
	var (
		text     = "text"
		actual   = StringPrefixSum(text)
		expected = []string{"", "t", "te", "tex", "text"}
	)
	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("StringPrefixSum(%s) == %v. Expected %v", text, actual, expected)
			return
		}
	}
}

func TestTernary(t *testing.T) {
	// string
	if TernaryString(true, "a", "b") != "a" {
		t.Error(`TernaryString(true, "a", "b") != "a". Expected "a"`)
	}
	if TernaryString(false, "a", "b") != "b" {
		t.Error(`TernaryString(false, "a", "b") != "b". Expected "b"`)
	}
	// bool
	if TernaryBool(true, false, true) {
		t.Error("TernaryBool(true, false, true) != false. Expected false")
	}
	if TernaryBool(false, true, false) {
		t.Error("TernaryBool(false, true, false) != false. Expected false")
	}
	// int
	if TernaryInt(true, 1, 0) != 1 {
		t.Error("TernaryInt(true, 1, 0) != 1. Expected 1")
	}
	if TernaryInt(false, 1, 0) != 0 {
		t.Error("TernaryInt(false, 1, 0) != 0. Expected 0")
	}
}

func TestMultiplyUint8(t *testing.T) {
	cases := []struct {
		a, b, expected uint8
	}{
		{0, 5, 0},
		{5, 0, 0},
		{10, 20, 200},
		{100, 20, 255},
	}
	for _, c := range cases {
		if actual := MultiplyUint8(c.a, c.b); actual != c.expected {
			t.Errorf("MultiplyUint8(%d, %d) == %d. Expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func TestMinusUint8(t *testing.T) {
	cases := []struct {
		a, b, expected uint8
	}{
		{200, 20, 180},
		{0, 5, 0},
		{5, 5, 0},
	}
	for _, c := range cases {
		if actual := MinusUint8(c.a, c.b); actual != c.expected {
			t.Errorf("MinusUint8(%d, %d) == %d. Expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func TestParseToStrings(t *testing.T) {
	cases := []struct {
		data     interface{}
		expected []string
	}{
		{5, make([]string, 0)},
		{[]interface{}{"a", "b", "c"}, make([]string, 0)},
		{[]interface{}{1.1, 2.5, 3.0}, []string{"1", "2", "3"}},
	}
	for _, c := range cases {
		actual := ParseToStrings(c.data)
		for i, text := range actual {
			if len(actual) != len(c.expected) || text != c.expected[i] {
				t.Errorf("ParseToStrings(%v) == %v. Expected %v", c.data, actual, c.expected)
				break
			}
		}
	}
}

func TestFloatFromMap(t *testing.T) {
	cases := []struct {
		table    map[string]interface{}
		key      string
		expected float64
		success  bool
	}{
		{map[string]interface{}{"a": 1.0}, "b", 0, false},
		{map[string]interface{}{"a": "text"}, "a", 0, false},
		{map[string]interface{}{"a": 1.0}, "a", 1.0, true},
	}
	for _, c := range cases {
		actual, isSuccess := FloatFromMap(c.table, c.key)
		if actual != c.expected || isSuccess != c.success {
			format := "FloatFromMap(%v, %s) == %f, %t. Expected %f, %t"
			t.Errorf(format, c.table, c.key, actual, isSuccess, c.expected, c.success)
		}
	}
}

func TestStringFromMap(t *testing.T) {
	cases := []struct {
		table    map[string]interface{}
		key      string
		expected string
		success  bool
	}{
		{map[string]interface{}{"a": "text"}, "b", "", false},
		{map[string]interface{}{"a": 2}, "a", "", false},
		{map[string]interface{}{"a": "text"}, "a", "text", true},
	}
	for _, c := range cases {
		actual, isSuccess := StringFromMap(c.table, c.key)
		if actual != c.expected || isSuccess != c.success {
			format := "StringFromMap(%v, %s) == %s, %t. Expected %s, %t"
			t.Errorf(format, c.table, c.key, actual, isSuccess, c.expected, c.success)
		}
	}
}
