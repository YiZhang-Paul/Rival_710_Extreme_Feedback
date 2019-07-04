package utils

import "fmt"

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

// IsSame checks if given color has same RGB values
func (rgb *RGB) IsSame(other *RGB) bool {
	return rgb.String() == other.String()
}

func (rgb *RGB) String() string {
	return fmt.Sprintf("R:%d;G:%d;B:%d", rgb.R, rgb.G, rgb.B)
}

// BlackRGB color
func BlackRGB() *RGB {
	return NewRGB(0, 0, 0)
}

// WhiteRGB color
func WhiteRGB() *RGB {
	return NewRGB(255, 255, 255)
}

// RedRGB color
func RedRGB() *RGB {
	return NewRGB(255, 0, 0)
}

// GreenRGB color
func GreenRGB() *RGB {
	return NewRGB(0, 255, 0)
}

// BlueRGB color
func BlueRGB() *RGB {
	return NewRGB(0, 0, 255)
}

// YellowRGB color
func YellowRGB() *RGB {
	return NewRGB(255, 255, 0)
}

// PinkRGB color
func PinkRGB() *RGB {
	return NewRGB(255, 0, 255)
}
