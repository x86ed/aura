package progbar

import "testing"

func TestNew(t *testing.T) {
	b := BasicProg{}
	b.New()
	if b.Dims.Y != 1 {
		t.Errorf("Expect")
	}
}
