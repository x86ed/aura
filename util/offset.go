package util

import (
	"unicode/utf8"
)

//EscOffset returns the offset of an escape string
func EscOffset(s string) (string, int) {
	o := utf8.RuneCountInString(s)
	return s, o
}
