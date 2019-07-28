package utils

import (
	"fmt"
	"strings"
)

// RGB color
type RGB struct {
	R uint8 `json:"red"`
	G uint8 `json:"green"`
	B uint8 `json:"blue"`
}

// NewRGB creates new RGB color
func NewRGB(r, g, b uint8) *RGB {
	return &RGB{r, g, b}
}

// ParseRGB takes a comma separated string of RGB values and creates new RGB color
func ParseRGB(rgb string) *RGB {
	rgbStrings := strings.Split(rgb, ",")
	if len(rgbStrings) != 3 {
		return NewRGB(0, 0, 0)
	}
	rgbUint8s := StringsToUint8s(rgbStrings)
	return NewRGB(rgbUint8s[0], rgbUint8s[1], rgbUint8s[2])
}

func (rgb *RGB) String() string {
	return fmt.Sprintf("R:%d;G:%d;B:%d", rgb.R, rgb.G, rgb.B)
}

// IsSame checks if given color has same RGB values
func (rgb *RGB) IsSame(other *RGB) bool {
	return rgb.String() == other.String()
}

// ReverseRGB returns a new reversed collection of colors
func ReverseRGB(colors []*RGB) []*RGB {
	reversed := make([]*RGB, 0)
	for i := len(colors) - 1; i >= 0; i-- {
		reversed = append(reversed, colors[i])
	}
	return reversed
}
