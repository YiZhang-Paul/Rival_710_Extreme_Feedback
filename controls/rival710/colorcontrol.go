package rival710

import (
	"log"
	"math"
	"strconv"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

// ColorControl controls device illuminations
type ColorControl struct {
	event int
}

func (cc ColorControl) bindStatic(meta controls.ColorMeta, event int) bool {
	data := cc.newBindMeta(meta.Game, strconv.Itoa(event), meta.Color)
	return utils.PostJSON(meta.BindAPI, data)
}

func (cc ColorControl) bindBlink(meta controls.ColorMeta, event int) bool {
	var (
		on = map[string]interface{}{
			"low":   1,
			"high":  1,
			"color": meta.Color,
		}
		off = map[string]interface{}{
			"low":   0,
			"high":  0,
			"color": utils.BlackRGB(),
		}
		colors = []interface{}{on, off}
		data   = cc.newBindMeta(meta.Game, strconv.Itoa(event), colors)
	)
	return utils.PostJSON(meta.BindAPI, data)
}

func (cc ColorControl) bindBreath(meta controls.ColorMeta, event int) bool {
	if meta.Frequency < 1 || meta.Frequency > 10 {
		log.Println("Frequency value out of range.")
		return false
	}
	var (
		colors = getBreathColorsInRanges(meta.Color, meta.Frequency)
		data   = cc.newBindMeta(meta.Game, strconv.Itoa(event), colors)
	)
	return utils.PostJSON(meta.BindAPI, data)
}

func getBreathColorsInRanges(start *utils.RGB, frequency float64) []map[string]interface{} {
	var (
		colors = getBreathColors(start, frequency)
		steps  = 100 / len(colors)
	)
	if steps > 1 {
		steps--
	}
	ranges := make([]map[string]interface{}, 0)
	for i, color := range colors {
		low := (steps + 1) * i
		newRange := map[string]interface{}{
			"low":   low,
			"color": color,
			"high":  100,
		}
		if i != len(colors)-1 {
			newRange["high"] = utils.MaxInt(100, low+steps)
		}
		ranges = append(ranges, newRange)
	}
	return ranges
}

func getBreathColors(start *utils.RGB, frequency float64) []*utils.RGB {
	var (
		colors = make([]*utils.RGB, 0)
		rgb    = []uint8{start.R, start.G, start.B}
		ticks  = math.Floor(50 / frequency)
		deltas = make([]int, len(rgb))
	)
	for i, value := range rgb {
		deltas[i] = int(math.Ceil(float64(value) / ticks))
	}
	for i := 0; i < int(ticks); i++ {
		newRgb := make([]uint8, len(rgb))
		for j, value := range rgb {
			newRgb[j] = utils.MaxUint8(0, value-uint8(deltas[j]*i))
		}
		colors = append(colors, utils.NewRGB(newRgb[0], newRgb[1], newRgb[2]))
	}
	reversed := utils.ReverseRGB(colors)
	if lastRGB := colors[len(colors)-1]; !lastRGB.IsSame(utils.BlackRGB()) {
		colors = append(colors, utils.BlackRGB())
	}
	return append(colors, reversed...)
}

func (cc *ColorControl) newBindMeta(game, event string, color interface{}) controls.BindMeta {
	var (
		meta         = controls.NewBindMeta(game, event)
		logoHandler  = controls.NewDeviceHandler("mouse", "logo", "color")
		wheelHandler = controls.NewDeviceHandler("mouse", "wheel", "color")
	)
	meta.Handlers = append(meta.Handlers, *logoHandler, *wheelHandler)
	for i := range meta.Handlers {
		meta.Handlers[i].Color = color
	}
	return *meta
}
