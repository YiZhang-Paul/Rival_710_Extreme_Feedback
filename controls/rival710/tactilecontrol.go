package rival710

import (
	"strconv"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

// TactileControl controls device tactile feedbacks
type TactileControl struct{}

func (tc TactileControl) bindTactile(meta controls.TactileMeta, event int) bool {
	var (
		rate = map[string]interface{}{
			"frequency":    meta.Frequency,
			"repeat_limit": meta.Limit,
		}
		pattern  = map[string]interface{}{"type": meta.Type}
		patterns = []map[string]interface{}{pattern}
		data     = tc.newBindMeta(meta.Game, event, patterns, rate)
	)
	return utils.PostJSON(meta.BindAPI, data)
}

func (tc TactileControl) applyTactile(meta controls.TactileMeta, event int) bool {
	data := controls.NewTriggerMeta(meta.Game, strconv.Itoa(event), 1)
	return utils.PostJSON(meta.TriggerAPI, data)
}

func (tc TactileControl) newBindMeta(game string, event int, pattern, rate interface{}) controls.BindMeta {
	var (
		meta    = controls.NewBindMeta(game, strconv.Itoa(event))
		handler = controls.NewDeviceHandler("tactile", "one", "vibrate")
	)
	handler.Pattern = pattern
	handler.Rate = rate
	meta.Handlers = append(meta.Handlers, *handler)
	return *meta
}
