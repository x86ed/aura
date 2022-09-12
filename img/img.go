package img

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/anthonynsimon/bild/transform"
	acolor "github.com/x86ed/aura/color"
	"github.com/x86ed/aura/cursor"
	"github.com/x86ed/aura/erase"
	"github.com/x86ed/aura/util"
)

type FileType uint

const (
	JPG FileType = iota
	GIF
	PNG
)

type Scale uint

const (
	Fit Scale = iota
	FitHeight
	FitWidth
	Stretch
)

// Image is an image file to be rendered to the Screen
type Img struct {
	util.RenderObj
	File     FileType
	Animated bool
	Alpha    bool
	FilePath string
	Pixels   [][]Pixel
	Image    image.Image
	Fit      Scale
}

// Draw Renderable interface method
func (i *Img) Draw() {
	i.Hide()
	e := i.load()
	if e != nil {
		fmt.Println(e)
	}
	e = i.scale2Dims(i.Fit)
	if e != nil {
		fmt.Println(e)
	}
	i.print2Screen()
}

// Hide Renderable interface method
func (i *Img) Hide() {
	out := erase.ToScreenEnd()
	if i.IsOffset {
		out = cursor.Move2Coord(i.Offset.X, i.Offset.Y) + out
	}
	fmt.Print(out)
}

// Interrupt Renderable interface method
func (i *Img) Interrupt() error {
	return nil
}

// load loads pixel data for an image
func (i *Img) load() error {
	file, err := os.OpenFile(i.FilePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if i.File == PNG {
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	}
	if i.File == JPG {
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	}
	if i.File == GIF {
		image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	}
	ii, _, err := image.Decode(file)
	i.Image = ii
	if err != nil {
		return err
	}
	bounds := ii.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(ii.At(x, y).RGBA()))
		}
		i.Pixels = append(i.Pixels, row)
	}

	return nil
}

// scale2Dims resizes an image to the
func (i *Img) scale2Dims(s Scale) error {

	var sx, sy int
	switch s {
	case Fit:
		sx, sy = i.calcFit()
	case FitWidth:
		sx, sy = i.calcFit(true)
	case FitHeight:
		sx, sy = i.calcFit(false)
	default:
		sx, sy = i.Dims.X, i.Dims.Y
	}
	ir := transform.Resize(i.Image, sx, sy, transform.Linear)
	bnd := ir.Bounds()
	ox := (bnd.Max.X - i.Dims.X) / 2
	if ox < 0 {
		ox = 0
	}
	oy := (bnd.Max.Y - i.Dims.Y) / 2
	if oy < 0 {
		oy = 0
	}
	ir = transform.Crop(ir, image.Rect(ox, oy, i.Dims.X+ox, i.Dims.Y+oy))
	i.Pixels = [][]Pixel{}
	bounds := ir.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(ir.At(x, y).RGBA()))
		}
		i.Pixels = append(i.Pixels, row)
	}
	i.Pixels = fixCrop(i.Pixels, oy)
	return nil
}

// fixCrop helper function that removes black dead space from the top of cropped images
func fixCrop(a [][]Pixel, lim int) [][]Pixel {
	return a[lim:]
}

func (i *Img) calcFit(b ...bool) (int, int) {
	// get xy values
	ib := i.Image.Bounds()
	ix, iy := ib.Max.X, ib.Max.Y
	bx, by := i.Dims.X, i.Dims.Y
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

func (i *Img) RGBArray2Ascii(arr [][]Pixel) string {
	var out string
	for _, v := range arr {
		for _, vv := range v {
			out += acolor.BgRGB(vv.R, vv.G, vv.B) + " "
		}
		out += acolor.Reset() + "\n"
		if i.IsOffset {
			out += cursor.Column(i.Offset.X)
		}
	}
	return out
}

// Print2Screen prints
func (i *Img) print2Screen() {
	out := i.RGBArray2Ascii(i.Pixels)
	fmt.Print(out)
}
