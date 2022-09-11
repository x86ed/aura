package progbar

import "testing"

func TestDraw(t *testing.T) {
	b := BasicProg{}
	b.New()
	b.Modes = []uint{Percent, Fraction, Count, Label}
	b.Draw()
}

func TestInterrupt(t *testing.T) {
	b := BasicProg{}
	b.New()
	b.Modes = []uint{Percent, Fraction, Count, Label}
	b.Draw()
	b.Interrupt()
}
