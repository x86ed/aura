package modal

import (
	"github.com/x86ed/aura/screen"
	"github.com/x86ed/aura/util"
)

type Modal struct {
	util.RenderObj
	Content    util.Renderable
	Buttons    map[string]func()
	Padding    int
	HasClose   bool
	hasAnimate bool
	GFX        ModalGlyph
}

type ModalGlyph struct {
	LUC, LLC, RUC, RBC, L, U, R, D string
}

func (m *Modal) Draw() {
	Scr := screen.S{}
	Scr.GetDims()

}
