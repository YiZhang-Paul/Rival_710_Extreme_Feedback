package rival710

import (
	"strconv"
	"time"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

// ScreenControl controls device OLED screen
type ScreenControl struct {
	totalSeconds int
	contents     []string
	event        int
}

func (sc ScreenControl) bindText(meta controls.ScreenMeta, event int) bool {
	data := sc.newBindMeta(meta, strconv.Itoa(event))
	return utils.PostJSON(meta.BindAPI, data)
}

func (sc ScreenControl) applyText(meta controls.ScreenMeta, text string, event int) bool {
	data := controls.NewTriggerMeta(meta.Game, strconv.Itoa(event), text)
	return utils.PostJSON(meta.TriggerAPI, data)
}

func (sc *ScreenControl) applyStatic(meta controls.ScreenMeta, event int) bool {
	sc.event = 0
	return sc.applyText(meta, meta.Content[0], event)
}

func (sc *ScreenControl) applyShift(meta controls.ScreenMeta, interval, event int) bool {
	sc.contents = meta.Content
	if sc.event == event {
		return false
	}
	var (
		loop         func(int)
		milliseconds = time.Duration(interval) * time.Microsecond
	)
	loop = func(index int) {
		if sc.event != event {
			return
		}
		content := sc.contents[utils.NextIndex(sc.contents, index)]
		sc.applyText(meta, content, event)
		select {
		case <-time.After(milliseconds):
			loop(utils.NextIndex(sc.contents, index))
		}
	}
	sc.event = event
	go loop(0)
	return true
}

func (sc ScreenControl) newBindMeta(screenMeta controls.ScreenMeta, event string) controls.BindMeta {
	var (
		meta    = controls.NewBindMeta(screenMeta.Game, event)
		handler = controls.NewDeviceHandler("screened", "one", "screen")
	)
	datas := map[string]interface{}{
		"has-text": true,
		"prefix":   screenMeta.Prefix,
		"suffix":   screenMeta.Suffix,
		"bold":     screenMeta.Bold,
	}
	if screenMeta.Icon != 0 {
		datas["icon-id"] = screenMeta.Icon
	}
	handler.Datas = []map[string]interface{}{datas}
	meta.Handlers = append(meta.Handlers, *handler)
	return *meta
}
