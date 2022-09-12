package progbar

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/x86ed/aura/color"
	"github.com/x86ed/aura/cursor"
	"github.com/x86ed/aura/erase"
	"github.com/x86ed/aura/screen"
	"github.com/x86ed/aura/util"
)

type BasicProg struct {
	ProgBar
}

func drawBar(b, c string, l, f int) string {
	out, off := util.EscOffset(color.TxtRGB(0xde, 0xfa, 0xce))
	// var out string
	for i := 0; i < l; i++ {
		out += b
	}
	if c != " " {
		out += c
	}
	for utf8.RuneCountInString(out)-off < f {
		out += " "
	}
	out += color.Reset()
	return out
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func humanFileSize(size float64) string {
	var suffixes [5]string
	suffixes[0] = "b"
	suffixes[1] = "kb"
	suffixes[2] = "mb"
	suffixes[3] = "gb"
	suffixes[4] = "tb"

	base := math.Log(size) / math.Log(1024)
	getSize := round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func countRate(c int, t time.Time) string {
	elapsed := time.Since(t) / time.Second
	amount := float64(c) / float64(elapsed)
	if amount <= 0 {
		return ""
	}
	return humanFileSize(amount)
}

func (p *BasicProg) basicDraw(m []uint) string {
	out := "*"
	lCap := " |"
	rCap := "| "
	bar := "█"
	endBar := []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
	spin := Spinners["Block"]

	for _, v := range m {
		switch v {
		case Percent:
			out += fmt.Sprintf("| %-6.2f %%", float32(p.Count)/float32(p.Capacity)*100)
		case Fraction:
			out += fmt.Sprintf("| %d/%d ", p.Count, p.Capacity)
		case Count:
			out += fmt.Sprintf("| %d of %d ", p.Count, p.Capacity)
		case Label:
			out = p.Label + out
		case Rate:
			out += fmt.Sprintf("| %s sec", countRate(p.Count, p.Started))
		}
	}
	bl := p.Dims.X - (utf8.RuneCountInString(out) - 1 + utf8.RuneCountInString(lCap) + utf8.RuneCountInString(rCap))

	if bl < 100 || p.Spinner {
		p.Spinner = true
		out := strings.Replace(out, "*", spin[p.Count%len(spin)], -1)
		return out

	}

	b := p.Count * bl / p.Capacity
	t := int((float64(p.Count)*float64(bl)/float64(p.Capacity) - float64(b)) * float64(len(endBar)))
	out = strings.Replace(out, "*", lCap+drawBar(bar, endBar[t], b, bl)+rCap, -1)
	return out
}

func (p *BasicProg) basicErase() string {
	out := cursor.Up(2) + erase.Line()
	if p.IsOffset {
		out += cursor.Move2Coord(p.Offset.X, p.Offset.Y)
	}
	return out
}

func (b *BasicProg) New() *BasicProg {
	var scr screen.S

	scr.GetDims()
	b.Modes = []uint{Percent}
	b.Dims.X = scr.Width
	b.Dims.Y = 1
	b.DrawBar = b.basicDraw
	b.Erase = b.basicErase
	b.Start()
	return b
}
