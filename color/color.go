package color

import (
	"strconv"
	"strings"
)

const (
	Prefix     = string("\u001b[")
	TxtBlack   = "30"
	TxtRed     = "31"
	TxtGreen   = "32"
	TxtYellow  = "33"
	TxtBlue    = "34"
	TxtMagenta = "35"
	TxtCyan    = "36"
	TxtWhite   = "37"
	BkBlack    = "40"
	BkRed      = "41"
	BkGreen    = "42"
	BkYellow   = "43"
	BkBlue     = "44"
	BkMagenta  = "45"
	BkCyan     = "46"
	BkWhite    = "47"
	Bright     = "1"
	Rst        = "0"
	Suffix     = "m"
)

// Black makes the text black or escapes the subsequent text black if called without arguments
func Black(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtBlack + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtBlack + Suffix
}

// Red makes the text black or escapes the subsequent text red if called without arguments
func Red(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtRed + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtRed + Suffix
}

// Green makes the text black or escapes the subsequent text green if called without arguments
func Green(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtGreen + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtGreen + Suffix
}

func Yellow(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtYellow + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtYellow + Suffix
}

func Blue(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtBlue + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtBlue + Suffix
}

func Magenta(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtMagenta + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtMagenta + Suffix
}

func Cyan(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtCyan + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtCyan + Suffix
}

func White(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtWhite + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtWhite + Suffix
}

// Bright

func BrightBlack(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtBlack + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtBlack + ";" + Bright + Suffix
}

func BrightRed(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtRed + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtRed + ";" + Bright + Suffix
}

func BrightGreen(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtGreen + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtGreen + ";" + Bright + Suffix
}

func BrightYellow(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtYellow + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtYellow + ";" + Bright + Suffix
}

func BrightBlue(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtBlue + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtBlue + ";" + Bright + Suffix
}

func BrightMagenta(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtMagenta + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtMagenta + ";" + Bright + Suffix
}

func BrightCyan(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtCyan + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtCyan + ";" + Bright + Suffix
}

func BrightWhite(s ...string) string {
	if len(s) > 0 {
		out := Prefix + TxtWhite + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + TxtWhite + ";" + Bright + Suffix
}

// Background

func BgBlack(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkBlack + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkBlack + Suffix
}

func BgRed(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkRed + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkRed + Suffix
}

func BgGreen(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkGreen + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkGreen + Suffix
}

func BgYellow(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkYellow + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkYellow + Suffix
}

func BgBlue(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkBlue + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkBlue + Suffix
}

func BgMagenta(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkMagenta + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkMagenta + Suffix
}

func BgCyan(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkCyan + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkCyan + Suffix
}

func BgWhite(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkWhite + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkWhite + Suffix
}

// Bright Background

func BgBrightBlack(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkBlack + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkBlack + ";" + Bright + Suffix
}

func BgBrightRed(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkRed + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkRed + ";" + Bright + Suffix
}

func BgBrightGreen(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkGreen + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkGreen + ";" + Bright + Suffix
}

func BgBrightYellow(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkYellow + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkYellow + ";" + Bright + Suffix
}

func BgBrightBlue(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkBlue + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkBlue + ";" + Bright + Suffix
}

func BgBrightMagenta(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkMagenta + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkMagenta + ";" + Bright + Suffix
}

func BgBrightCyan(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkCyan + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkCyan + ";" + Bright + Suffix
}

func BgBrightWhite(s ...string) string {
	if len(s) > 0 {
		out := Prefix + BkWhite + ";" + Bright + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + BkWhite + ";" + Bright + Suffix
}

// Pallettes

func Txt256(i int, s ...string) string {

	if len(s) > 0 {
		out := Prefix + "38;5;" + strconv.Itoa(i) + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + "38;5;" + strconv.Itoa(i) + Suffix
}

func Bg256(i int, s ...string) string {

	if len(s) > 0 {
		out := Prefix + "48;5;" + strconv.Itoa(i) + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + "48;5;" + strconv.Itoa(i) + Suffix
}

func TxtRGB(r, g, b int, s ...string) string {
	if len(s) > 0 {
		out := Prefix + "38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + "38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + Suffix
}

func BgRGB(r, g, b int, s ...string) string {
	if len(s) > 0 {
		out := Prefix + "48;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + Suffix + strings.Join(s, " ") + Prefix + Rst + Suffix
		return out
	}
	return Prefix + "48;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + Suffix
}

func Reset() string {
	return Prefix + Rst + Suffix
}
