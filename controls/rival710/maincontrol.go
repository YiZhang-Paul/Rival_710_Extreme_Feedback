package rival710

import (
	"log"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

// Controller is the main control for device illumination, OLED screen and tactile feedback
type Controller struct {
	Color   colorController
	Screen  screenController
	Tactile tactileController
}

// NewController will create a new controller and make sure to bind all events
func NewController() *Controller {
	c := &Controller{
		Color:   &ColorControl{},
		Screen:  &ScreenControl{},
		Tactile: &TactileControl{},
	}
	c.bindEvents()
	return c
}

func newControlMeta() *controls.ControlMeta {
	return &controls.ControlMeta{
		Game:       game,
		BindAPI:    bindAPI,
		TriggerAPI: triggerAPI,
		RemoveAPI:  removeEventAPI,
	}
}

func newColorMeta(color *utils.RGB, frequency float64) *controls.ColorMeta {
	return &controls.ColorMeta{
		ControlMeta: *newControlMeta(),
		Color:       color,
		Frequency:   frequency,
	}
}

func newScreenMeta(content []string, icon int, prefix string, bold bool) *controls.ScreenMeta {
	return &controls.ScreenMeta{
		ControlMeta: *newControlMeta(),
		Content:     content,
		Icon:        icon,
		Prefix:      prefix,
		Bold:        bold,
	}
}

func newTactileMeta(name string, frequency, limit int) *controls.TactileMeta {
	return &controls.TactileMeta{
		ControlMeta: *newControlMeta(),
		Type:        name,
		Frequency:   frequency,
		Limit:       limit,
	}
}

func (c Controller) bindEvents() {
	c.bindPassingEvent()
	c.bindBrokenEvent()
	c.bindBuildingEvent()
	c.bindOnCompleteEvent()
	c.bindDeployingEvent()
	c.bindPendingEvent()
	c.bindDeployBrokenEvent()
}

func (c Controller) bindPassingEvent() {
	screenMeta := newScreenMeta(nil, 1, "", false)
	c.Screen.bindText(*screenMeta, controls.PassingScreen)
	colorMeta := newColorMeta(utils.GreenRGB(), 2)
	c.Color.bindStatic(*colorMeta, controls.PassingColor)
}

func (c Controller) bindBrokenEvent() {
	screenMeta := newScreenMeta(nil, 6, "", false)
	c.Screen.bindText(*screenMeta, controls.BrokenScreen)
	colorMeta := newColorMeta(utils.RedRGB(), 2)
	c.Color.bindBlink(*colorMeta, controls.BrokenColor)
}

func (c Controller) bindBuildingEvent() {
	screenMeta := newScreenMeta(nil, 13, "", false)
	c.Screen.bindText(*screenMeta, controls.BuildingScreen)
	colorMeta := newColorMeta(utils.YellowRGB(), 2)
	c.Color.bindBreath(*colorMeta, controls.BuildingColor)
}

func (c Controller) bindOnCompleteEvent() {
	tactileMeta := newTactileMeta("ti_predefined_tripleclick_100", 1, 2)
	c.Tactile.bindTactile(*tactileMeta, controls.BuildFailedTactile)
	c.Tactile.bindTactile(*tactileMeta, controls.DeployFailedTactile)
	tactileMeta = newTactileMeta("ti_predefined_strongclick_100", 2, 3)
	c.Tactile.bindTactile(*tactileMeta, controls.BuiltTactile)
	c.Tactile.bindTactile(*tactileMeta, controls.DeployedTactile)
	screenMeta := newScreenMeta(nil, 8, "", true)
	c.Screen.bindText(*screenMeta, controls.BuiltScreen)
	screenMeta = newScreenMeta(nil, 2, "", true)
	c.Screen.bindText(*screenMeta, controls.DeployedScreen)
	screenMeta = newScreenMeta(nil, 7, "", true)
	c.Screen.bindText(*screenMeta, controls.BuildFailedScreen)
	screenMeta = newScreenMeta(nil, 5, "", true)
	c.Screen.bindText(*screenMeta, controls.DeployFailedScreen)
}

func (c Controller) bindDeployingEvent() {
	screenMeta := newScreenMeta(nil, 16, "", false)
	c.Screen.bindText(*screenMeta, controls.DeployingScreen)
	colorMeta := newColorMeta(utils.BlueRGB(), 2)
	c.Color.bindBreath(*colorMeta, controls.DeployingColor)
}

func (c Controller) bindPendingEvent() {
	tactileMeta := newTactileMeta("ti_predefined_strongbuzz_100", 1, 2)
	c.Tactile.bindTactile(*tactileMeta, controls.PendingTactile)
	screenMeta := newScreenMeta(nil, 15, "", false)
	c.Screen.bindText(*screenMeta, controls.PendingScreen)
	colorMeta := newColorMeta(utils.PinkRGB(), 2)
	c.Color.bindBlink(*colorMeta, controls.PendingColor)
}

func (c Controller) bindDeployBrokenEvent() {
	screenMeta := newScreenMeta(nil, 6, "", false)
	c.Screen.bindText(*screenMeta, controls.DeployBrokenScreen)
	colorMeta := newColorMeta(utils.WhiteRGB(), 2)
	c.Color.bindBlink(*colorMeta, controls.DeployBrokenColor)
}

// RemoveGame de-registers given game from SteelSeries Engine
func (c Controller) RemoveGame(game string) {
	data := map[string]interface{}{"game": game}
	utils.PostJSON(removeGameAPI, data)
}

// Execute receives data from ci/cd services and control devices to provide feedbacks on received data
func (c Controller) Execute(meta controls.NotificationMeta) {
	data, ok := meta.Data.(map[string]interface{})
	if !ok {
		log.Println("Invalid notification format.")
		return
	}
	if meta.Event == "ci" {
		c.executeCi(meta.Mode, data)
	} else if meta.Event == "cd" {
		c.executeCd(meta.Mode, data)
	}
}

func (c Controller) handleComplete(data map[string]interface{}, screenEvent, tactileEvent int) {
	branch, ok := utils.StringFromMap(data, "branch")
	if !ok {
		log.Print("Missing branch information.")
		return
	}
	screenMeta := newScreenMeta([]string{branch, ""}, 0, "", false)
	c.Screen.applyStatic(*screenMeta, screenEvent)
	tactileMeta := newTactileMeta("", 0, 0)
	c.Tactile.applyTactile(*tactileMeta, tactileEvent)
}
