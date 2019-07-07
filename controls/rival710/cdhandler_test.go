package rival710

import (
	"testing"

	"github.com/yi-zhang/rival-710-extreme-feedback/controls"
)

func TestHandleDeploying(t *testing.T) {
	setupMockController()
	notification := controls.NotificationMeta{
		Event: "cd",
		Mode:  "deploying",
		Data:  map[string]interface{}{"branch": "UAT"},
	}
	controller.Execute(notification)
	var (
		content       = []string{"", "U", "UA", "UAT"}
		actualContent = screenControl.applyShiftSpy.meta.Content
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
	if colorControl.applyBreathSpy.count != 1 {
		t.Errorf("Expected applyBreath() to be called once, called %d times instead.", colorControl.applyBreathSpy.count)
	}
	if screenControl.applyShiftSpy.count != 1 {
		t.Errorf("Expected applyShift() to be called once, called %d times instead.", screenControl.applyShiftSpy.count)
	}
}

func TestHandlePending(t *testing.T) {
	setupMockController()
	notification := controls.NotificationMeta{
		Event: "cd",
		Mode:  "pending",
		Data:  map[string]interface{}{"branch": "DEV"},
	}
	controller.Execute(notification)
	var (
		content                 = []string{"DEV", ""}
		actualContent           = screenControl.applyShiftSpy.meta.Content
		frequency       float64 = 1
		actualFrequency         = colorControl.applyBlinkSpy.meta.Frequency
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
	if frequency != actualFrequency {
		t.Errorf("Expected frequency to be %f, got %f", frequency, actualFrequency)
	}
	if colorControl.applyBlinkSpy.count != 1 {
		t.Errorf("Expected applyBlink() to be called once, called %d times instead.", colorControl.applyBlinkSpy.count)
	}
	if screenControl.applyShiftSpy.count != 1 {
		t.Errorf("Expected applyShift() to be called once, called %d times instead.", screenControl.applyShiftSpy.count)
	}
}

func TestHandleDeployBroken(t *testing.T) {
	setupMockController()
	notification := controls.NotificationMeta{
		Event: "cd",
		Mode:  "deploy-broken",
		Data:  map[string]interface{}{"branch": "PROD"},
	}
	controller.Execute(notification)
	var (
		content                 = []string{"PROD", ""}
		actualContent           = screenControl.applyShiftSpy.meta.Content
		frequency       float64 = 2
		actualFrequency         = colorControl.applyBlinkSpy.meta.Frequency
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
	if frequency != actualFrequency {
		t.Errorf("Expected frequency to be %f, got %f", frequency, actualFrequency)
	}
	if colorControl.applyBlinkSpy.count != 1 {
		t.Errorf("Expected applyBlink() to be called once, called %d times instead.", colorControl.applyBlinkSpy.count)
	}
	if screenControl.applyShiftSpy.count != 1 {
		t.Errorf("Expected applyShift() to be called once, called %d times instead.", screenControl.applyShiftSpy.count)
	}
}

func TestHandleDeployed(t *testing.T) {
	setupMockController()
	notification := controls.NotificationMeta{
		Event: "cd",
		Mode:  "deployed",
		Data:  map[string]interface{}{"branch": "UAT"},
	}
	controller.Execute(notification)
	var (
		content       = []string{"UAT"}
		actualContent = screenControl.applyStaticSpy.meta.Content
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
}

func TestHandleDeployFailed(t *testing.T) {
	setupMockController()
	notification := controls.NotificationMeta{
		Event: "cd",
		Mode:  "deploy-failed",
		Data:  map[string]interface{}{"branch": "DEV"},
	}
	controller.Execute(notification)
	var (
		content       = []string{"DEV"}
		actualContent = screenControl.applyStaticSpy.meta.Content
	)
	for i, text := range content {
		if text != actualContent[i] {
			t.Errorf("Expected content to be %v, got %v", content, actualContent)
			break
		}
	}
}
