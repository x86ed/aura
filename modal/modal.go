package modal

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/x86ed/aura/cursor"
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
	LUC, LDC, RUC, RDC, L, U, R, D, F string
}

func (m *Modal) Draw(o ...util.Coord) {
	var out string
	var app string
	if len(o) > 0 {
		m.IsOffset = true
		m.Offset.Y += o[0].Y
		m.Offset.X += o[0].X
	}
	if m.IsOffset == true {
		out += cursor.Move2Coord(m.Offset.X, m.Offset.Y)
		app = cursor.Column(m.Offset.X)
	}
	out += m.GFX.LUC + strings.Repeat(m.GFX.U, m.Dims.X-2) + m.GFX.RUC + "\n" + app
	for i := 0; i < m.Dims.Y-2; i++ {
		out += m.GFX.L + strings.Repeat(m.GFX.F, m.Dims.X-2) + m.GFX.R + "\n" + app
	}
	out += m.GFX.LDC + strings.Repeat(m.GFX.D, m.Dims.X-2) + m.GFX.RDC + "\n" + app
	fmt.Print(out)
	newOff := m.Offset
	newOff.X += (utf8.RuneCountInString(m.GFX.L) + m.Padding)
	newOff.Y += (utf8.RuneCountInString(m.GFX.U) + m.Padding)
	newDim := m.Dims
	newDim.X -= (utf8.RuneCountInString(m.GFX.L) + utf8.RuneCountInString(m.GFX.R) + (m.Padding * 2))
	newDim.Y -= (utf8.RuneCountInString(m.GFX.U) + utf8.RuneCountInString(m.GFX.D) + (m.Padding * 2))
	padding := util.Coord{X: (utf8.RuneCountInString(m.GFX.L) + m.Padding), Y: (utf8.RuneCountInString(m.GFX.U) + m.Padding)}
	m.Content.Draw(newOff, newDim, padding)
}

// Hide Renderable interface method
func (m *Modal) Hide(o ...util.Coord) {
	var app string
	if len(o) > 0 {
		m.IsOffset = true
		m.Offset.Y += o[0].Y
		m.Offset.X += o[0].X
	}
	if m.IsOffset {
		fmt.Print(cursor.Move2Coord(m.Offset.X, m.Offset.Y))
		app = cursor.Column(m.Offset.X)
	}
	var p string
	for l := 0; l < m.Dims.Y; l++ {
		for c := 0; c < m.Dims.X; c++ {
			p += " "
		}
		p += "\n" + app
	}
	fmt.Print(p)
}

// Interrupt Renderable interface method
func (m *Modal) Interrupt() error {
	return nil
}
