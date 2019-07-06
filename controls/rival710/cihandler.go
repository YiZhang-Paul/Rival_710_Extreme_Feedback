package rival710

import (
	"fmt"
	"log"
	"strings"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

func (c Controller) executeCi(mode string, data map[string]interface{}) {
	switch mode {
	case "passing":
		c.handlePassing(data)
	case "broken":
		c.handleBroken(data)
	case "building":
		c.handleBuilding(data)
	case "built":
		c.handleBuilt(data)
	case "build-failed":
		c.handleBuildFailed(data)
	}
}

func (c Controller) handlePassing(data map[string]interface{}) {
	var (
		content = []string{"P: ", "M: ", "D: "}
		lists   = []interface{}{data["pull"], data["merge"], data["deploy"]}
	)
	for i, list := range lists {
		content[i] += strings.Join(utils.ParseToStrings(list), "/")
	}
	screenMeta := newScreenMeta(content, 0, "", true)
	c.Screen.applyShift(*screenMeta, 4500, controls.PassingScreen)
	c.Color.applyStatic(*newColorMeta(nil, 0), controls.PassingColor)
}

func (c Controller) handleBroken(data map[string]interface{}) {
	total, ok := utils.FloatFromMap(data, "total")
	if !ok {
		log.Print("Missing total number of broken builds.")
		return
	}
	time, _ := utils.FloatFromMap(data, "time")
	screenMeta := newScreenMeta(nil, 0, fmt.Sprintf("%d|", int(total)), false)
	c.Screen.startTimer(*screenMeta, controls.BrokenScreen, int(time)/1000)
	c.Color.applyBlink(*newColorMeta(nil, 2), controls.BrokenColor)
}

func (c Controller) handleBuilding(data map[string]interface{}) {
	total, ok := utils.FloatFromMap(data, "total")
	if !ok {
		log.Print("Missing total number of broken builds.")
		return
	}
	var (
		time, _  = utils.FloatFromMap(data, "time")
		minutes  = int(time) / 60000
		prefix   = utils.TernaryString(minutes < 1, "<", "~")
		suffix   = utils.TernaryString(minutes > 1, "s", "")
		duration = fmt.Sprintf("%s%d%smin", prefix, minutes, suffix)
		content  = fmt.Sprintf("%d|%s", int(total), duration)
	)
	screenMeta := newScreenMeta([]string{content}, 0, "", false)
	c.Screen.applyStatic(*screenMeta, controls.BuildingScreen)
	colorMeta := newColorMeta(nil, 0)
	c.Color.applyBreath(*colorMeta, controls.BuildingColor)
}

func (c Controller) handleBuilt(data map[string]interface{}) {
	c.handleComplete(data, controls.BuiltScreen, controls.BuiltTactile)
}

func (c Controller) handleBuildFailed(data map[string]interface{}) {
	c.handleComplete(data, controls.BuildFailedScreen, controls.BuildFailedTactile)
}
