package utils

import "testing"

func TestFormatTime(t *testing.T) {
	cases := []struct {
		seconds  int
		expected string
	}{
		{8, "00:00:08"},
		{59, "00:00:59"},
		{60 * 3, "00:03:00"},
		{60 * 31, "00:31:00"},
		{3600, "01:00:00"},
		{3600 * 12, "12:00:00"},
	}
	for _, c := range cases {
		if actual := FormatTime(c.seconds); actual != c.expected {
			t.Errorf("FormatTime(%d) == %s. Expected %s", c.seconds, actual, c.expected)
		}
	}
}
