package rival710

import "github.com/yi-zhang/rival-710-extreme-feedback/controls"

type tactileMetaSpy struct {
	count int
	meta  controls.TactileMeta
}

type mockTactileController struct {
	bindTactileSpy  *tactileMetaSpy
	applyTactileSpy *tactileMetaSpy
}

func newMockTactileController() *mockTactileController {
	return &mockTactileController{
		bindTactileSpy:  &tactileMetaSpy{},
		applyTactileSpy: &tactileMetaSpy{},
	}
}

func (mtc *mockTactileController) bindTactile(meta controls.TactileMeta, event int) bool {
	mtc.bindTactileSpy.count++
	mtc.bindTactileSpy.meta = meta
	return true
}

func (mtc *mockTactileController) applyTactile(meta controls.TactileMeta, event int) bool {
	mtc.applyTactileSpy.count++
	mtc.applyTactileSpy.meta = meta
	return true
}
