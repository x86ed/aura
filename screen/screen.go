package screen

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	prefix      = string("\u001b[=")
	basePrefix  = string("\u001b[")
	setSuffix   = "h"
	resetSuffix = "l"
	save        = "?47h"
	restore     = "?47l"
	startAlt    = "?1049h"
	endAlt      = "?1049l"
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

// Set40x25Mono sets the screen to 40 x 25 monochrome (text)
func Set40x25Mono() string {
	return prefix + "0" + setSuffix
}

// Unset40x25Mono unsets the screen to 40 x 25 monochrome (text)
func Unset40x25Mono() string {
	return prefix + "0" + resetSuffix
}

// Set40x25Color sets the screen to 40 x 25 color (text)
func Set40x25Color() string {
	return prefix + "1" + setSuffix
}

// Unset40x25Color unsets the screen to 40 x 25 color (text)
func Unset40x25Color() string {
	return prefix + "1" + resetSuffix
}

// Set80x25Mono sets the screen to 80 x 25 monochrome (text)
func Set80x25Mono() string {
	return prefix + "2" + setSuffix
}

// Unset80x25Mono unsets the screen to 80 x 25 monochrome (text)
func Unset80x25Mono() string {
	return prefix + "2" + resetSuffix
}

// Set80x25Color sets the screen to 80 x 25 color (text)
func Set80x25Color() string {
	return prefix + "3" + setSuffix
}

// Unset80x25Color unsets the screen to 80 x 25 color (text)
func Unset80x25Color() string {
	return prefix + "3" + resetSuffix
}

// Set4Color320x200 sets the screen to 320 x 200 4-color (graphics)
func Set4Color320x200() string {
	return prefix + "4" + setSuffix
}

// Unset4Color320x200 unsets the screen to 320 x 200 4-color (graphics)
func Unset4Color320x200() string {
	return prefix + "4" + resetSuffix
}

// Set320x200Mono sets the screen to 320 x 200 monochrome (graphics)
func Set320x200Mono() string {
	return prefix + "5" + setSuffix
}

// Unset320x200Mono unsets the screen to 320 x 200 monochrome (graphics)
func Unset320x200Mono() string {
	return prefix + "5" + resetSuffix
}

// Set600x200Mono sets the screen to 640 x 200 monochrome (graphics)
func Set600x200Mono() string {
	return prefix + "6" + setSuffix
}

// Unset600x200Mono unsets the screen to 640 x 200 monochrome (graphics)
func Unset600x200Mono() string {
	return prefix + "6" + resetSuffix
}

// Set320x200Color sets the screen to 320 x 200 color (graphics)
func Set320x200Color() string {
	return prefix + "13" + setSuffix
}

// Unset320x200Color unsets the screen to 320 x 200 color (graphics)
func Unset320x200Color() string {
	return prefix + "13" + resetSuffix
}

// Set16Color640x200 sets the screen to 640 x 200 color (16-color graphics)
func Set16Color640x200() string {
	return prefix + "14" + setSuffix
}

// Unset16Color640x200 unsets the screen to 640 x 200 color (16-color graphics)
func Unset16Color640x200() string {
	return prefix + "14" + resetSuffix
}

// Set640x350Mono sets the screen to 640 x 350 monochrome (2-color graphics)
func Set640x350Mono() string {
	return prefix + "15" + setSuffix
}

// Unset640x350Mono unsets the screen to 640 x 350 monochrome (2-color graphics)
func Unset640x350Mono() string {
	return prefix + "15" + resetSuffix
}

// Set16Color640x350 sets the screen to 640 x 350 color (16-color graphics)
func Set16Color640x350() string {
	return prefix + "16" + setSuffix
}

// Unset16Color640x350 unsets the screen to 640 x 350 color (16-color graphics)
func Unset16Color640x350() string {
	return prefix + "16" + resetSuffix
}

// Set640x480mono sets the screen to 640 x 480 monochrome (2-color graphics)
func Set640x480mono() string {
	return prefix + "17" + setSuffix
}

// Unset640x480mono unsets the screen to 640 x 480 monochrome (2-color graphics)
func Unset640x480mono() string {
	return prefix + "17" + resetSuffix
}

// Set16Color640x480 sets the screen to 640 x 480 color (16-color graphics)
func Set16Color640x480() string {
	return prefix + "18" + setSuffix
}

// Unset16Color640x480 unsets the screen to 640 x 480 color (16-color graphics)
func Unset16Color640x480() string {
	return prefix + "18" + resetSuffix
}

// Set256Color320x200 sets the screen to 320 x 200 color (256-color graphics)
func Set256Color320x200() string {
	return prefix + "19" + setSuffix
}

// Unset256Color320x200 unsets the screen to 320 x 200 color (256-color graphics)
func Unset256Color320x200() string {
	return prefix + "19" + resetSuffix
}

// SetLineWrapping Enables line wrapping
func SetLineWrapping() string {
	return prefix + "7" + setSuffix
}

// UnsetLineWrapping disables line wrapping
func UnsetLineWrapping() string {
	return prefix + "7" + resetSuffix
}

// SaveScreen saves the contents of the screen
func SaveScreen() string {
	return basePrefix + save
}

// RestoreScreen restores the screen with the contents of the last saved screen
func RestoreScreen() string {
	return basePrefix + restore
}

// EnableAltBuffer Enables the alternate buffer
// for more information on what the alternative buffer is
// and why you would need it read here:
// https://jameshfisher.com/2017/12/04/how-less-works/
func EnableAltBuffer() string {
	return basePrefix + startAlt
}

// DisableAltBuffer disables the alternate buffer
func DisableAltBuffer() string {
	return basePrefix + endAlt
}
