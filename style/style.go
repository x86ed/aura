package style

import "strings"

const (
	Prefix   = string("\u001b[")
	Bld      = "1"
	Lite     = "2"
	Ital     = "3"
	Underln  = "4"
	Blnk     = "5"
	Bldd     = "6"
	Inv      = "7"
	Invz     = "8"
	Crs      = "9"
	DUnder   = "21"
	Normal   = "22"
	XItalic  = "23"
	XUnder   = "24"
	Steady   = "25"
	Positive = "27"
	Visible  = "28"
	XCross   = "29"
	Suffix   = "m"
	Rst      = "0"
)

// Bold makes text bold
func Bold(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Bld + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Bld + Suffix
}

// Light makes text Light
func Light(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Lite + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Lite + Suffix
}

// Italic makes text Italic
func Italic(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Ital + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Ital + Suffix
}

// Underline makes text Underline
func Underline(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Underln + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Underln + Suffix
}

// Blink makes text Blink
func Blink(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Blnk + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Blnk + Suffix
}

// Bold2 makes text Bold2
func Bold2(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Bldd + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Bldd + Suffix
}

// Invisible makes text Invisible
func Invisible(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Inv + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Inv + Suffix
}

// Invisible2 makes text Invisible2
func Invisible2(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Invz + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Invz + Suffix
}

// Cross makes text Cross
func Cross(s ...string) string {
	if len(s) > 0 {
		out := Prefix + Crs + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + Crs + Suffix
}

// DoubleUnderline makes text DoubleUnderline
func DoubleUnderline(s ...string) string {
	if len(s) > 0 {
		out := Prefix + DUnder + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + DUnder + Suffix
}
