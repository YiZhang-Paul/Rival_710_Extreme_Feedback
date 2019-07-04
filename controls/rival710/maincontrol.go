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
