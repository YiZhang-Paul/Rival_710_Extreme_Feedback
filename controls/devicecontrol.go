package controls

// DeviceHandler contains basic meta information to control RGB devices
type DeviceHandler struct {
	Type    string      `json:"device-type"`
	Zone    string      `json:"zone"`
	Mode    string      `json:"mode"`
	Color   interface{} `json:"color,omitempty"`
	Datas   interface{} `json:"datas,omitempty"`
	Pattern interface{} `json:"pattern,omitempty"`
	Rate    interface{} `json:"rate,omitempty"`
}

// NewDeviceHandler creates a basic DeviceHandler to control RGB devices
func NewDeviceHandler(deviceType, zone, mode string) *DeviceHandler {
	return &DeviceHandler{
		Type: deviceType,
		Zone: zone,
		Mode: mode,
	}
}
