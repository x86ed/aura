package screen

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type S struct {
	Width  int
	Height int
}

// GetDims gets the screen's dimensions
func (s *S) GetDims() error {
	var err error
	s.Width, s.Height, err = terminal.GetSize(int(os.Stdin.Fd()))
	return err
}
