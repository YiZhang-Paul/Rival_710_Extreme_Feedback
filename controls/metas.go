package controls

import (
	"github.com/yi-zhang/rival-710-extreme-feedback/utils"
)

// NotificationMeta contains required information from ci/cd services
type NotificationMeta struct {
	Event string      `json:"event"`
	Mode  string      `json:"mode"`
	Data  interface{} `json:"data,omitempty"`
}

// ControlMeta contains basic event information used to construct other meta
type ControlMeta struct {
	Game       string `json:"game"`
	BindAPI    string `json:"bindApi"`
	TriggerAPI string `json:"triggerApi"`
	RemoveAPI  string `json:"removeApi,omitempty"`
}

// BindMeta contains data to register events and provides default event behaviors
type BindMeta struct {
	Game          string          `json:"game"`
	Event         string          `json:"event"`
	Handlers      []DeviceHandler `json:"handlers,omitempty"`
	Pattern       string          `json:"pattern,omitempty"`
	ValueOptional bool            `json:"value_optional,omitempty"`
}

// TriggerMeta contains data to fire registered events with necessary inputs
type TriggerMeta struct {
	Game  string `json:"game"`
	Event string `json:"event"`
	Data  *struct {
		Value interface{} `json:"value"`
	} `json:"data"`
}

// ColorMeta for device illumination control
type ColorMeta struct {
	ControlMeta
	Color     *utils.RGB `json:"color"`
	Frequency float64    `json:"frequency,omitempty"`
	Limit     int        `json:"limit,omitempty"`
}

// ScreenMeta for OLED screen control
type ScreenMeta struct {
	ControlMeta
	Content []string `json:"content"`
	Icon    int      `json:"icon,omitempty"`
	Prefix  string   `json:"prefix,omitempty"`
	Suffix  string   `json:"suffix,omitempty"`
	Bold    bool     `json:"bold,omitempty"`
}

// TactileMeta for tactile feedback control
type TactileMeta struct {
	ControlMeta
	Type      string `json:"type"`
	Frequency int    `json:"frequency,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}
