package block

import (
	"fmt"
	"strings"

	"github.com/x86ed/aura/cursor"
	"github.com/x86ed/aura/util"
)

// Block is a linked list of renderable strings.
// The first in the list renders below the others.
type Block struct {
	util.RenderObj
	Content string
	Next    *Block
}

// Draw Renderable method
func (b *Block) Draw(o ...util.Coord) {
	var out string
	var app string
	var paddingX string
	var paddingY string
	if len(o) > 0 {
		b.IsOffset = true
		b.Offset.Y += o[0].Y
		b.Offset.X += o[0].X
	}
	if len(o) > 1 {
		b.Dims.Y += o[1].Y
		b.Dims.X += o[1].X
	}
	if len(o) > 2 {
		paddingX = strings.Repeat(" ", o[2].X)
		paddingY = strings.Repeat("\n", o[2].Y)
	}
	if b.IsOffset == true {
		out += cursor.Move2Coord(b.Offset.X, b.Offset.Y)
		app = cursor.Column(b.Offset.X)
	}
	cs := strings.Split(b.Content, "\n")
	for len(cs) < b.Dims.Y+3 {
		cs = append(cs, "")
	}
	for py := 0; py < b.Dims.Y; py++ {
		if len(cs[py]) > b.Dims.X {
			out += cs[py][:b.Dims.X] + "\n" + app
		} else {
			out += cs[py] + "\n" + app
		}
	}
	fmt.Print(out + paddingX + paddingY)
	if b.Next != nil {
		b.Next.Draw(o...)
	}
}

// Hide Renderable Method
func (b *Block) Hide(o ...util.Coord) {
	if len(o) > 0 {
		b.IsOffset = true
		b.Offset.Y += o[0].Y
		b.Offset.X += o[0].X
	}
	if len(o) > 1 {
		b.Dims.Y = o[1].Y
		b.Dims.X = o[1].X
	}
	var app string
	if b.IsOffset {
		fmt.Print(cursor.Move2Coord(b.Offset.X, b.Offset.Y))
		app = cursor.Column(b.Offset.X)
	}
	var p string
	for l := 0; l < b.Dims.Y; l++ {
		for c := 0; c < b.Dims.X; c++ {
			p += " "
		}
		p += "\n" + app
	}
	fmt.Print(p)
}

// Interrupt Renderable interface method
func (b *Block) Interrupt() error {
	return nil
}
