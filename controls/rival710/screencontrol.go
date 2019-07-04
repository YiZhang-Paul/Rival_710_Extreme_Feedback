package rival710

import (
	"strconv"

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
