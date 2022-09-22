package img

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"time"

	"github.com/anthonynsimon/bild/transform"
	acolor "github.com/x86ed/aura/color"
	"github.com/x86ed/aura/cursor"
	"github.com/x86ed/aura/util"
)

type Gif struct {
	util.RenderObj
	Animated bool
	Alpha    bool
	FilePath string
	Pixels   [][][]Pixel
	Fit      Scale
	Frames   []*image.Paletted
	Index    int
	Delay    []int
}

// Draw Renderable interface method
func (g *Gif) Draw(o ...util.Coord) {
	if len(o) > 0 {
		g.IsOffset = true
		g.Offset = o[0]
	}
	if g.IsOffset {
		fmt.Print(cursor.Move2Coord(g.Offset.X, g.Offset.Y))
	}
	e := g.load()
	if e != nil {
		fmt.Println(e)
	}
	e = g.scale2Dims(g.Fit)
	if e != nil {
		fmt.Println(e)
	}
	g.print2Screen()
	g.Index = (g.Index + 1) % len(g.Frames)
	time.Sleep(time.Millisecond * time.Duration(10*g.Delay[g.Index]))
	g.Draw(g.Offset)
}

// Hide Renderable interface method
func (g *Gif) Hide(o ...util.Coord) {
	if len(o) > 0 {
		g.IsOffset = true
		g.Offset = o[0]
	}
	if len(o) > 1 {
		g.Dims.Y += o[0].Y
		g.Dims.X += o[0].X
	}
	var app string
	if g.IsOffset {
		fmt.Print(cursor.Move2Coord(g.Offset.X, g.Offset.Y))
		app = cursor.Column(g.Offset.X)
	}
	var p string
	for l := 0; l < g.Dims.Y; l++ {
		for c := 0; c < g.Dims.X; c++ {
			p += "  "
		}
		p += "\n" + app
	}
	fmt.Print(p)
}

// Interrupt Renderable interface method
func (g *Gif) Interrupt() error {
	return nil
}

// load loads pixel data for an image
func (g *Gif) load() error {
	file, err := os.OpenFile(g.FilePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	ii, err := gif.DecodeAll(file)
	g.Frames = ii.Image
	g.Delay = ii.Delay
	if err != nil {
		return err
	}
	bounds := g.Frames[g.Index].Rect
	width, height := bounds.Max.X, bounds.Max.Y
	for z := 0; z < len(g.Frames); z++ {
		var frame [][]Pixel
		for y := 0; y < height; y++ {
			var row []Pixel
			for x := 0; x < width; x++ {
				row = append(row, rgbaToPixel(g.Frames[g.Index].At(x, y).RGBA()))
			}
			frame = append(frame, row)
		}
		g.Pixels = append(g.Pixels, frame)
	}

	return nil
}

// scale2Dims resizes an image to the
func (g *Gif) scale2Dims(s Scale) error {

	var sx, sy int
	switch s {
	case Fit:
		sx, sy = g.calcFit()
	case FitWidth:
		sx, sy = g.calcFit(true)
	case FitHeight:
		sx, sy = g.calcFit(false)
	default:
		sx, sy = g.Dims.X, g.Dims.Y
	}
	ir := transform.Resize(g.Frames[g.Index], sx, sy, transform.Linear)
	bnd := ir.Bounds()
	ox := (bnd.Max.X - g.Dims.X) / 2
	if ox < 0 {
		ox = 0
	}
	oy := (bnd.Max.Y - g.Dims.Y) / 2
	if oy < 0 {
		oy = 0
	}
	ir = transform.Crop(ir, image.Rect(ox, oy, g.Dims.X+ox, g.Dims.Y+oy))
	g.Pixels = [][][]Pixel{}
	bounds := ir.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for z := 0; z < len(g.Frames); z++ {
		var frame [][]Pixel
		for y := 0; y < height; y++ {
			var row []Pixel
			for x := 0; x < width; x++ {
				row = append(row, rgbaToPixel(ir.At(x, y).RGBA()))
			}
			frame = append(frame, row)
		}
		frame = fixCrop(frame, oy)
		g.Pixels = append(g.Pixels, frame)
	}
	return nil
}

func (g *Gif) calcFit(b ...bool) (int, int) {
	// get xy values
	ib := g.Frames[g.Index].Rect
	ix, iy := ib.Max.X, ib.Max.Y
	bx, by := g.Dims.X, g.Dims.Y
	// x/y rations
	ir := float32(ix) / float32(iy)
	br := float32(bx) / float32(by)

	x2x := float32(bx) / float32(ix)
	y2y := float32(by) / float32(iy)

	// if ratio is equal return box size
	if ir == br {
		return bx, by
	}

	// FitWidth
	if len(b) > 0 && b[0] {
		nx := bx
		ny := float32(iy) * x2x
		return nx, int(ny)
	}

	// FitHeight
	if len(b) > 0 && !b[0] {
		ny := by
		nx := float32(ix) * y2y
		return int(nx), ny
	}

	if by > bx {
		ny := by
		nx := float32(ix) * y2y
		return int(nx), ny
	}

	nx := bx
	ny := float32(iy) * x2x
	return nx, int(ny)
}

func (g *Gif) RGBArray2Ascii(arr [][][]Pixel) string {
	var out string
	iarr := arr[g.Index]
	for _, v := range iarr {
		for _, vv := range v {
			out += acolor.BgRGB(vv.R, vv.G, vv.B) + "  "
		}
		out += acolor.Reset() + "\n"
		if g.IsOffset {
			out += cursor.Column(g.Offset.X)
		}
	}
	return out
}

// Print2Screen prints
func (g *Gif) print2Screen() {
	out := g.RGBArray2Ascii(g.Pixels)
	fmt.Print(out)
}
