package utils

import "testing"

func TestMaxInt(t *testing.T) {
	cases := []struct {
		values   []int
		expected int
	}{
		{[]int{-1, 0, 2}, 2},
		{[]int{6, -7, 5}, 6},
	}
	for _, c := range cases {
		if actual := MaxInt(c.values...); actual != c.expected {
			t.Errorf("MaxInt(%v) == %d. Expected %d", c.values, actual, c.expected)
		}
	}
}

func TestMinInt(t *testing.T) {
	cases := []struct {
		values   []int
		expected int
	}{
		{[]int{-1, 0, 2}, -1},
		{[]int{6, -7, 5}, -7},
	}
	for _, c := range cases {
		if actual := MinInt(c.values...); actual != c.expected {
			t.Errorf("MinInt(%v) == %d. Expected %d", c.values, actual, c.expected)
		}
	}
}
