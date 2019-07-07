package rival710

import "github.com/yi-zhang/rival-710-extreme-feedback/controls"

type colorMetaSpy struct {
	count int
	meta  controls.ColorMeta
}

type mockColorController struct {
	bindStaticSpy  *colorMetaSpy
	bindBlinkSpy   *colorMetaSpy
	bindBreathSpy  *colorMetaSpy
	applyStaticSpy *colorMetaSpy
	applyBlinkSpy  *colorMetaSpy
	applyBreathSpy *colorMetaSpy
}

func newMockColorController() *mockColorController {
	return &mockColorController{
		bindBlinkSpy:   &colorMetaSpy{},
		bindBreathSpy:  &colorMetaSpy{},
		bindStaticSpy:  &colorMetaSpy{},
		applyBlinkSpy:  &colorMetaSpy{},
		applyBreathSpy: &colorMetaSpy{},
		applyStaticSpy: &colorMetaSpy{},
	}
}

func (mcc *mockColorController) bindStatic(meta controls.ColorMeta, event int) bool {
	mcc.bindStaticSpy.count++
	mcc.bindStaticSpy.meta = meta
	return true
}

func (mcc *mockColorController) bindBlink(meta controls.ColorMeta, event int) bool {
	mcc.bindBlinkSpy.count++
	mcc.bindBlinkSpy.meta = meta
	return true
}

func (mcc *mockColorController) bindBreath(meta controls.ColorMeta, event int) bool {
	mcc.bindBreathSpy.count++
	mcc.bindBreathSpy.meta = meta
	return true
}

func (mcc *mockColorController) applyStatic(meta controls.ColorMeta, event int) bool {
	mcc.applyStaticSpy.count++
	mcc.applyStaticSpy.meta = meta
	return true
}

func (mcc *mockColorController) applyBlink(meta controls.ColorMeta, event int) bool {
	mcc.applyBlinkSpy.count++
	mcc.applyBlinkSpy.meta = meta
	return true
}

func (mcc *mockColorController) applyBreath(meta controls.ColorMeta, event int) bool {
	mcc.applyBreathSpy.count++
	mcc.applyBreathSpy.meta = meta
	return true
}
