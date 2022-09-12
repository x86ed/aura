package cursor

import (
	"strconv"
)

const (
	prefix     = string("\u001b[")
	prefixDEC  = string("\u001b")
	up         = "A"
	down       = "B"
	right      = "C"
	left       = "D"
	beginNext  = "E"
	beginPrev  = "F"
	column     = "G"
	home       = "H"
	position   = "6n"
	move       = "M"
	decSave    = "7"
	decRestore = "8"
	save       = "s"
	restore    = "u"
	invisible  = "?25l"
	visible    = "?25h"
)

// Up moves the cursor up by the number specified
func Up(i int) string {
	return prefix + strconv.Itoa(i) + up
}

// Down moves the cursor down by the number specified
func Down(i int) string {
	return prefix + strconv.Itoa(i) + down
}

// Right moves the cursor right by the number specified
func Right(i int) string {
	return prefix + strconv.Itoa(i) + right
}

// Left moves the cursor left by the number specified
func Left(i int) string {
	return prefix + strconv.Itoa(i) + left
}

// BeginNext moves the cursor to the begining of the next line by the number of lines specified
func BeginNext(i int) string {
	return prefix + strconv.Itoa(i) + beginNext
}

// BeginPrev moves the cursor to the beginning of the previous line by the number of lines specified
func BeginPrev(i int) string {
	return prefix + strconv.Itoa(i) + beginPrev
}

// Column moves the cursor to the column i
func Column(i int) string {
	return prefix + strconv.Itoa(i) + column
}

// Home moves the cursor to 0,0
func Home() string {
	return prefix + home
}

// Position prints the cursor position to stdout
func Position() string {
	return prefix + position
}

// Move moves the cursor one line up
func Move() string {
	return prefixDEC + move
}

// Save saves the cursor position flag[0] toggles dec flag[1] toggles sco
func Save(flag ...bool) string {
	sco := prefix + save
	dec := prefixDEC + decSave
	if len(flag) > 0 && flag[0] == false {
		dec = ""
	}
	if len(flag) > 1 && flag[1] == false {
		sco = ""
	}
	return sco + dec
}

// Restore restores the cursor position flag[0] toggles dec flag[1] toggles sco
func Restore(flag ...bool) string {
	sco := prefix + restore
	dec := prefixDEC + decRestore
	if len(flag) > 0 && flag[0] == false {
		dec = ""
	}
	if len(flag) > 1 && flag[1] == false {
		sco = ""
	}
	return sco + dec
}

func Move2Coord(x, y int, alt ...bool) string {
	if len(alt) > 0 && alt[0] {
		return prefix + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "f"
	}
	return prefix + strconv.Itoa(y) + ";" + strconv.Itoa(x) + home
}

// Invisible makes the cursor invisible
func Invisible() string {
	return prefix + invisible
}

// Visible makes the cursor visible
func Visible() string {
	return prefix + visible
}
