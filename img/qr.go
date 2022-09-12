package img

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	qrcode "github.com/skip2/go-qrcode"
	acolor "github.com/x86ed/aura/color"
	"github.com/x86ed/aura/cursor"
	"github.com/x86ed/aura/util"
)

// Pixel is an RGBA pixel value with a string
type Pixel struct {
	R   int
	G   int
	B   int
	A   int
	Str string
}

// QR is the type used by the QRFromxxx functions.
// for colors
type QR struct {
	util.RenderObj
	Recovery qrcode.RecoveryLevel
	Size     int
	FGColor  string
	BGColor  string
	Value    string
	Content  string
}

func getPixels(file []byte) ([][]Pixel, error) {
	img, _, err := image.Decode(bytes.NewReader(file))

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257), ""}
}

func (q *QR) BnWArray2Ascii(arr [][]Pixel, fg, bg string) (string, int, int) {
	var out string
	for _, v := range arr {
		for _, vv := range v {
			if vv.R > 0 {
				out += bg + "  "
			} else {
				out += fg + "  "
			}
		}
		out += acolor.Reset() + "\n"
		if q.IsOffset {
			out += cursor.Column(q.Offset.X)
		}
	}
	return out, len(arr[0]) * 2, len(arr)
}

func (q *QR) QRFromText(s string) error {
	q.Value = s
	raw, err := qrcode.Encode(q.Value, q.Recovery, q.Size)
	if err != nil {
		return err
	}
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	pixels, err := getPixels(raw)
	if err != nil {
		return err
	}
	draw, x, y := q.BnWArray2Ascii(pixels, q.FGColor, q.BGColor)
	q.Dims.X = x
	q.Dims.Y = y
	fmt.Print(draw)
	return nil
}

func (q *QR) Hide() {
	out := cursor.Home()
	if q.IsOffset {
		out += cursor.Move2Coord(q.Offset.X, q.Offset.Y)
	}
	fmt.Print(out)
}

func (q *QR) Interrupt() error {
	return nil
}

func (q *QR) Draw() {
	q.Hide()
	raw, _ := qrcode.Encode(q.Value, q.Recovery, q.Size)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	pixels, _ := getPixels(raw)
	draw, x, y := q.BnWArray2Ascii(pixels, q.FGColor, q.BGColor)
	q.Dims.X = x
	q.Dims.Y = y
	fmt.Print(draw)
}
