package rival710

import "github.com/yi-zhang/rival-710-extreme-feedback/controls"

type screenMetaSpy struct {
	count int
	meta  controls.ScreenMeta
}

type mockScreenController struct {
	bindTextSpy    *screenMetaSpy
	applyStaticSpy *screenMetaSpy
	applyShiftSpy  *screenMetaSpy
	startTimerSpy  *screenMetaSpy
}

func newMockScreenController() *mockScreenController {
	return &mockScreenController{
		bindTextSpy:    &screenMetaSpy{},
		applyStaticSpy: &screenMetaSpy{},
		applyShiftSpy:  &screenMetaSpy{},
		startTimerSpy:  &screenMetaSpy{},
	}
}

func (msc *mockScreenController) bindText(meta controls.ScreenMeta, event int) bool {
	msc.bindTextSpy.count++
	msc.bindTextSpy.meta = meta
	return true
}

func (msc *mockScreenController) applyStatic(meta controls.ScreenMeta, event int) bool {
	msc.applyStaticSpy.count++
	msc.applyStaticSpy.meta = meta
	return true
}

func (msc *mockScreenController) applyShift(meta controls.ScreenMeta, interval, event int) bool {
	msc.applyShiftSpy.count++
	msc.applyShiftSpy.meta = meta
	return true
}

func (msc *mockScreenController) startTimer(meta controls.ScreenMeta, event, seconds int) bool {
	msc.startTimerSpy.count++
	msc.startTimerSpy.meta = meta
	return true
}
