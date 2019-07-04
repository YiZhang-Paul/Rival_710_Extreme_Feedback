package rival710

import "github.com/yi-zhang/rival-710-extreme-feedback/controls"

type colorController interface {
	bindStatic(controls.ColorMeta, int) bool
	bindBlink(controls.ColorMeta, int) bool
	bindBreath(controls.ColorMeta, int) bool
	applyStatic(controls.ColorMeta, int) bool
	applyBlink(controls.ColorMeta, int) bool
	applyBreath(controls.ColorMeta, int) bool
}

type screenController interface {
	bindText(controls.ScreenMeta, int) bool
	applyStatic(controls.ScreenMeta, int) bool
	applyShift(controls.ScreenMeta, int, int) bool
	continueTimer(controls.ScreenMeta, int, string, int) bool
}

type tactileController interface {
	bindTactile(controls.TactileMeta, int) bool
	applyTactile(controls.TactileMeta, int) bool
}
