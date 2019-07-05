package rival710

import (
	"log"
	"math"
	"strconv"
	"time"

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
		colors = getBreathColorRanges(meta.Color, meta.Frequency)
		data   = cc.newBindMeta(meta.Game, strconv.Itoa(event), colors)
	)
	return utils.PostJSON(meta.BindAPI, data)
}

func (cc *ColorControl) applyStatic(meta controls.ColorMeta, event int) bool {
	cc.event = 0
	data := controls.NewTriggerMeta(meta.Game, strconv.Itoa(event), 1)
	return utils.PostJSON(meta.TriggerAPI, data)
}

func (cc *ColorControl) applyBlink(meta controls.ColorMeta, event int) bool {
	if cc.event == event {
		return false
	}
	var (
		loop     func(bool)
		interval = time.Duration(1000/meta.Frequency) * time.Millisecond
	)
	loop = func(on bool) {
		if cc.event != event {
			return
		}
		value := utils.TernaryInt(on, 1, 0)
		data := controls.NewTriggerMeta(meta.Game, strconv.Itoa(event), value)
		utils.PostJSON(meta.TriggerAPI, data)
		select {
		case <-time.After(interval):
			loop(!on)
		}
	}
	cc.event = event
	go loop(false)
	return true
}

func (cc *ColorControl) applyBreath(meta controls.ColorMeta, event int) bool {
	if cc.event == event {
		return false
	}
	var loop func(int, bool)
	loop = func(value int, up bool) {
		if cc.event != event {
			return
		}
		data := controls.NewTriggerMeta(meta.Game, strconv.Itoa(event), value)
		utils.PostJSON(meta.TriggerAPI, data)
		up = utils.TernaryBool(up, value != 100, value == 0)
		value = utils.TernaryInt(up, value+1, value-1)
		select {
		case <-time.After(10 * time.Millisecond):
			loop(value, up)
		}
	}
	cc.event = event
	go loop(0, true)
	return true
}

func getBreathColorRanges(start *utils.RGB, frequency float64) []map[string]interface{} {
	var (
		ranges = make([]map[string]interface{}, 0)
		colors = getBreathColors(start, frequency)
		steps  = utils.MinInt(1, 100/len(colors))
	)
	for i, color := range colors {
		low := (steps + 1) * i
		high := utils.MinInt(100, low+steps)
		ranges = append(ranges, map[string]interface{}{
			"low":   low,
			"high":  utils.TernaryInt(i != len(colors)-1, high, 100),
			"color": color,
		})
	}
	return ranges
}

func getBreathColors(base *utils.RGB, frequency float64) []*utils.RGB {
	var (
		colors = make([]*utils.RGB, 0)
		ticks  = math.Floor(50 / frequency)
		deltaR = int(math.Ceil(float64(base.R) / ticks))
		deltaG = int(math.Ceil(float64(base.G) / ticks))
		deltaB = int(math.Ceil(float64(base.B) / ticks))
	)
	for i := 0; i < int(ticks); i++ {
		r := utils.MaxUint8(0, base.R-uint8(deltaR*i))
		g := utils.MaxUint8(0, base.G-uint8(deltaG*i))
		b := utils.MaxUint8(0, base.B-uint8(deltaB*i))
		colors = append(colors, utils.NewRGB(r, g, b))
	}
	colors = append(colors, utils.BlackRGB())
	return append(colors, utils.ReverseRGB(colors)[1:]...)
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
