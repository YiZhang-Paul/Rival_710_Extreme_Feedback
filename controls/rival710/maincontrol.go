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

func (c Controller) executeCi(mode string, data map[string]interface{}) {
	switch mode {
	}
}

func (c Controller) executeCd(mode string, data map[string]interface{}) {
	switch mode {
	}
}
