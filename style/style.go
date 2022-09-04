package style

import "strings"

const (
	prefix          = string("\u001b[")
	bold            = "1"
	light           = "2"
	italic          = "3"
	underline       = "4"
	blink           = "5"
	bold2           = "6"
	invert          = "7"
	invisible       = "8"
	cross           = "9"
	doubleUnderline = "21"
	normal          = "22"
	unItalic        = "23"
	unUnderlined    = "24"
	steady          = "25"
	positive        = "27"
	visible         = "28"
	uncross         = "29"
	suffix          = "m"
	reset           = "0"
)

// Bold makes text bold
func Bold(s ...string) string {
	if len(s) > 0 {
		out := prefix + bold + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + bold + suffix
}

// Light makes text Light
func Light(s ...string) string {
	if len(s) > 0 {
		out := prefix + light + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + light + suffix
}

// Italic makes text Italic
func Italic(s ...string) string {
	if len(s) > 0 {
		out := prefix + italic + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + italic + suffix
}

// Underline makes text Underline
func Underline(s ...string) string {
	if len(s) > 0 {
		out := prefix + underline + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + underline + suffix
}

// Blink makes text Blink
func Blink(s ...string) string {
	if len(s) > 0 {
		out := prefix + blink + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + blink + suffix
}

// Bold2 makes text Bold2
func Bold2(s ...string) string {
	if len(s) > 0 {
		out := prefix + bold2 + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + bold2 + suffix
}

// Invert makes text inverted in apearance
func Invert(s ...string) string {
	if len(s) > 0 {
		out := prefix + invert + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + invert + suffix
}

// Invisible makes text Invisible
func Invisible(s ...string) string {
	if len(s) > 0 {
		out := prefix + invisible + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + invisible + suffix
}

// Cross makes text Cross
func Cross(s ...string) string {
	if len(s) > 0 {
		out := prefix + cross + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + cross + suffix
}

// DoubleUnderline makes text DoubleUnderline
func DoubleUnderline(s ...string) string {
	if len(s) > 0 {
		out := prefix + doubleUnderline + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + doubleUnderline + suffix
}

// Normal makes text Normal
func Normal(s ...string) string {
	if len(s) > 0 {
		out := prefix + normal + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + normal + suffix
}

// Unitalic makes text Unitalic
func UnItalic(s ...string) string {
	if len(s) > 0 {
		out := prefix + unItalic + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + unItalic + suffix
}

// UnUnderlined makes text UnUnderlined
func UnUnderlined(s ...string) string {
	if len(s) > 0 {
		out := prefix + unUnderlined + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + unUnderlined + suffix
}

// Steady makes text Steady
func Steady(s ...string) string {
	if len(s) > 0 {
		out := prefix + steady + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + steady + suffix
}

// Positive makes text Positive
func Positive(s ...string) string {
	if len(s) > 0 {
		out := prefix + positive + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + positive + suffix
}

// Visible makes text Visible
func Visible(s ...string) string {
	if len(s) > 0 {
		out := prefix + visible + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + visible + suffix
}

// UnCross makes text UnCross
func UnCross(s ...string) string {
	if len(s) > 0 {
		out := prefix + uncross + suffix + strings.Join(s, " ") + prefix + reset + suffix
		return out
	}
	return prefix + uncross + suffix
}

func Reset() string {
	return prefix + reset + suffix
}
