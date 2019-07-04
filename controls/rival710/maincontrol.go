package rival710

// Controller is the main control for device illumination, OLED screen and tactile feedback
type Controller struct {
	Color   colorController
	Screen  screenController
	Tactile tactileController
}
