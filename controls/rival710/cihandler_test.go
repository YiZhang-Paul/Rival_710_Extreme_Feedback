package rival710

import "testing"

var (
	colorControl   *mockColorController
	screenControl  *mockScreenController
	tactileControl *mockTactileController
	controller     *Controller
)

func setup() {
	colorControl = newMockColorController()
	screenControl = newMockScreenController()
	tactileControl = newMockTactileController()
	controller = &Controller{colorControl, screenControl, tactileControl}
}

func TestHandlePassing(t *testing.T) {
	setup()
	data := map[string]interface{}{
		"pull":   []interface{}{float64(2), float64(6), float64(8)},
		"merge":  []interface{}{float64(0), float64(3), float64(4)},
		"deploy": []interface{}{float64(0), float64(3), float64(3)},
	}
	controller.executeCi("passing", data)
	var (
		content       = []string{"P: 2/6/8", "M: 0/3/4", "D: 0/3/3"}
		actualContent = screenControl.applyShiftSpy.meta.Content
	)
	for i, text := range actualContent {
		if text != content[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
	if colorControl.applyStaticSpy.count != 1 {
		t.Errorf("Expected applyStatic() to be called once, called %d times instead.", colorControl.applyStaticSpy.count)
	}
	if screenControl.applyShiftSpy.count != 1 {
		t.Errorf("Expected applyShift() to be called once, called %d times instead.", screenControl.applyShiftSpy.count)
	}
}

func TestHandleBroken(t *testing.T) {
	setup()
	data := map[string]interface{}{
		"total": float64(2),
		"time":  float64(3600000),
	}
	controller.executeCi("broken", data)
	var (
		prefix                  = "2|"
		actualPrefix            = screenControl.startTimerSpy.meta.Prefix
		frequency       float64 = 2
		actualFrequency         = colorControl.applyBlinkSpy.meta.Frequency
	)
	if prefix != actualPrefix {
		t.Errorf("Expected prefix to be %s, got %s", prefix, actualPrefix)
	}
	if frequency != actualFrequency {
		t.Errorf("Expected frequency to be %f, got %f", frequency, actualFrequency)
	}
	if colorControl.applyBlinkSpy.count != 1 {
		t.Errorf("Expected applyBlink() to be called once, called %d times instead.", colorControl.applyBlinkSpy.count)
	}
	if screenControl.startTimerSpy.count != 1 {
		t.Errorf("Expected startTimer() to be called once, called %d times instead.", screenControl.startTimerSpy.count)
	}
}

func TestHandleBuildingOneMinute(t *testing.T) {
	setup()
	data := map[string]interface{}{
		"total": float64(2),
		"time":  float64(60000),
	}
	controller.executeCi("building", data)
	var (
		content       = "2|~1min"
		actualContent = screenControl.applyStaticSpy.meta.Content[0]
	)
	if content != actualContent {
		t.Errorf("Expected content to be %s, got %s", content, actualContent)
	}
	if colorControl.applyBreathSpy.count != 1 {
		t.Errorf("Expected applyBreath() to be called once, called %d times instead.", colorControl.applyBreathSpy.count)
	}
	if screenControl.applyStaticSpy.count != 1 {
		t.Errorf("Expected applyStatic() to be called once, called %d times instead.", screenControl.applyStaticSpy.count)
	}
}

func TestHandleBuildingLessThanOneMinute(t *testing.T) {
	setup()
	data := map[string]interface{}{
		"total": float64(2),
		"time":  float64(59999),
	}
	controller.executeCi("building", data)
	var (
		content       = "2|<1min"
		actualContent = screenControl.applyStaticSpy.meta.Content[0]
	)
	if content != actualContent {
		t.Errorf("Expected content to be %s, got %s", content, actualContent)
	}
}

func TestHandleBuildingMoreThanOneMinute(t *testing.T) {
	setup()
	data := map[string]interface{}{
		"total": float64(2),
		"time":  float64(180000),
	}
	controller.executeCi("building", data)
	var (
		content       = "2|~3mins"
		actualContent = screenControl.applyStaticSpy.meta.Content[0]
	)
	if content != actualContent {
		t.Errorf("Expected content to be %s, got %s", content, actualContent)
	}
}

func TestHandleBuilt(t *testing.T) {
	setup()
	controller.executeCi("built", map[string]interface{}{"branch": "DEV"})
	var (
		content       = []string{"DEV", ""}
		actualContent = screenControl.applyStaticSpy.meta.Content
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
}

func TestHandleBuildFailed(t *testing.T) {
	setup()
	controller.executeCi("build-failed", map[string]interface{}{"branch": "PROD"})
	var (
		content       = []string{"PROD", ""}
		actualContent = screenControl.applyStaticSpy.meta.Content
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
}
