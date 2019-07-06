package utils

import "testing"

func TestRGBString(t *testing.T) {
	cases := []struct {
		rgb      *RGB
		expected string
	}{
		{NewRGB(0, 10, 125), "R:0;G:10;B:125"},
		{NewRGB(15, 50, 215), "R:15;G:50;B:215"},
	}
	for _, c := range cases {
		if toString := c.rgb.String(); toString != c.expected {
			t.Errorf("%v.String() == %s. Expected %s", c.rgb, toString, c.expected)
		}
	}
}

func TestIsSame(t *testing.T) {
	cases := []struct {
		a, b     [3]uint8
		expected bool
	}{
		{[3]uint8{0, 10, 125}, [3]uint8{0, 10, 125}, true},
		{[3]uint8{15, 50, 215}, [3]uint8{15, 50, 215}, true},
		{[3]uint8{155, 0, 150}, [3]uint8{155, 12, 150}, false},
		{[3]uint8{10, 5, 100}, [3]uint8{10, 5, 95}, false},
	}
	for _, c := range cases {
		var (
			a = NewRGB(c.a[0], c.a[1], c.a[2])
			b = NewRGB(c.b[0], c.b[1], c.b[2])
		)
		if a.IsSame(b) != c.expected {
			t.Errorf("%v == %v => %t. Expected %t", a, b, !c.expected, c.expected)
		}
	}
}

func TestReverseRGB(t *testing.T) {
	var (
		colors   = []*RGB{NewRGB(0, 1, 2), NewRGB(5, 0, 3), NewRGB(9, 1, 4)}
		actual   = ReverseRGB(colors)
		expected = []*RGB{NewRGB(9, 1, 4), NewRGB(5, 0, 3), NewRGB(0, 1, 2)}
	)
	for i, color := range actual {
		other := expected[i]
		if color.R != other.R || color.G != other.G || color.B != other.B {
			t.Errorf("Reverse(%v) == %v. Expected %v", colors, actual, expected)
		}
	}
	if &colors == &actual {
		t.Error("Expected Reverse() to return new slice with different array")
	}
}
