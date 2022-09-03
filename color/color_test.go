package color

import "testing"

type tst struct {
	input, expected string
}

func TestBlack(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtBlack + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Black(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Black()
	if ttt != Prefix+TxtBlack+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtBlack, Suffix, ttt)
	}
}

func TestRed(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtRed + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Red(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Red()
	if ttt != Prefix+TxtRed+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtRed, Suffix, ttt)
	}
}

func TestGreen(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtGreen + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Green(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Green()
	if ttt != Prefix+TxtGreen+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtGreen, Suffix, ttt)
	}
}

func TestYellow(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtYellow + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Yellow(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Yellow()
	if ttt != Prefix+TxtYellow+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtYellow, Suffix, ttt)
	}
}

func TestBlue(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtBlue + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Blue(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Blue()
	if ttt != Prefix+TxtBlue+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtBlue, Suffix, ttt)
	}
}

func TestMagenta(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtMagenta + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Magenta(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Magenta()
	if ttt != Prefix+TxtMagenta+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtMagenta, Suffix, ttt)
	}
}

func TestCyan(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtCyan + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Cyan(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Cyan()
	if ttt != Prefix+TxtCyan+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtCyan, Suffix, ttt)
	}
}

func TestWhite(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtWhite + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := White(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := White()
	if ttt != Prefix+TxtWhite+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, TxtWhite, Suffix, ttt)
	}
}

// Bright

func TestBrightBlack(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtBlack + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightBlack(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightBlack()
	if ttt != Prefix+TxtBlack+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtBlack, ";", Bright, Suffix, ttt)
	}
}

func TestBrightRed(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtRed + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightRed(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightRed()
	if ttt != Prefix+TxtRed+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtRed, ";", Bright, Suffix, ttt)
	}
}

func TestBrightGreen(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtGreen + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightGreen(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightGreen()
	if ttt != Prefix+TxtGreen+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtGreen, ";", Bright, Suffix, ttt)
	}
}

func TestBrightYellow(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtYellow + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightYellow(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightYellow()
	if ttt != Prefix+TxtYellow+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtYellow, ";", Bright, Suffix, ttt)
	}
}

func TestBrightBlue(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtBlue + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightBlue(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightBlue()
	if ttt != Prefix+TxtBlue+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtBlue, ";", Bright, Suffix, ttt)
	}
}

func TestBrightMagenta(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtMagenta + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightMagenta(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightMagenta()
	if ttt != Prefix+TxtMagenta+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtMagenta, ";", Bright, Suffix, ttt)
	}
}

func TestBrightCyan(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtCyan + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightCyan(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightCyan()
	if ttt != Prefix+TxtCyan+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtCyan, ";", Bright, Suffix, ttt)
	}
}

func TestBrightWhite(t *testing.T) {
	tsts := []tst{{"test", Prefix + TxtWhite + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BrightWhite(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BrightWhite()
	if ttt != Prefix+TxtWhite+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, TxtWhite, ";", Bright, Suffix, ttt)
	}
}

// Background

func TestBgBlack(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkBlack + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBlack(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBlack()
	if ttt != Prefix+BkBlack+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkBlack, Suffix, ttt)
	}
}

func TestBgRed(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkRed + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgRed(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgRed()
	if ttt != Prefix+BkRed+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkRed, Suffix, ttt)
	}
}

func TestBgGreen(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkGreen + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgGreen(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgGreen()
	if ttt != Prefix+BkGreen+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkGreen, Suffix, ttt)
	}
}

func TestBgYellow(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkYellow + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgYellow(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgYellow()
	if ttt != Prefix+BkYellow+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkYellow, Suffix, ttt)
	}
}

func TestBgBlue(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkBlue + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBlue(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBlue()
	if ttt != Prefix+BkBlue+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkBlue, Suffix, ttt)
	}
}

func TestBgMagenta(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkMagenta + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgMagenta(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgMagenta()
	if ttt != Prefix+BkMagenta+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkMagenta, Suffix, ttt)
	}
}

func TestBgCyan(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkCyan + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgCyan(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgCyan()
	if ttt != Prefix+BkCyan+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkCyan, Suffix, ttt)
	}
}

func TestBgWhite(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkWhite + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgWhite(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgWhite()
	if ttt != Prefix+BkWhite+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, BkWhite, Suffix, ttt)
	}
}

// Bright BG

func TestBgBrightBlack(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkBlack + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightBlack(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightBlack()
	if ttt != Prefix+BkBlack+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkBlack, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightRed(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkRed + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightRed(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightRed()
	if ttt != Prefix+BkRed+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkRed, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightGreen(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkGreen + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightGreen(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightGreen()
	if ttt != Prefix+BkGreen+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkGreen, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightYellow(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkYellow + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightYellow(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightYellow()
	if ttt != Prefix+BkYellow+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkYellow, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightBlue(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkBlue + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightBlue(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightBlue()
	if ttt != Prefix+BkBlue+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkBlue, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightMagenta(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkMagenta + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightMagenta(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightMagenta()
	if ttt != Prefix+BkMagenta+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkMagenta, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightCyan(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkCyan + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightCyan(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightCyan()
	if ttt != Prefix+BkCyan+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkCyan, ";", Bright, Suffix, ttt)
	}
}

func TestBgBrightWhite(t *testing.T) {
	tsts := []tst{{"test", Prefix + BkWhite + ";" + Bright + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgBrightWhite(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgBrightWhite()
	if ttt != Prefix+BkWhite+";"+Bright+Suffix {
		t.Errorf("Expected: %s%s%s%s%s Got: %s", Prefix, BkWhite, ";", Bright, Suffix, ttt)
	}
}

// Pallettes

func TestTxt256(t *testing.T) {
	tsts := []tst{{"test", Prefix + "38;5;" + "255" + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Txt256(255, tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Txt256(255)
	if ttt != Prefix+"38;5;255"+Suffix {
		t.Errorf("Expected: %s38;5;255%s Got: %s", Prefix, Suffix, ttt)
	}
}

func TestBg256(t *testing.T) {
	tsts := []tst{{"test", Prefix + "48;5;" + "255" + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := Bg256(255, tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Bg256(255)
	if ttt != Prefix+"48;5;255"+Suffix {
		t.Errorf("Expected: %s48;5;255%s Got: %s", Prefix, Suffix, ttt)
	}
}

func TestTxtRGB(t *testing.T) {
	tsts := []tst{{"test", Prefix + "38;2;" + "255;255;255" + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := TxtRGB(255, 255, 255, tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := TxtRGB(255, 255, 255)
	if ttt != Prefix+"38;2;255;255;255"+Suffix {
		t.Errorf("Expected: %s38;2;255;255;255%s Got: %s", Prefix, Suffix, ttt)
	}
}

func TestBgRGB(t *testing.T) {
	tsts := []tst{{"test", Prefix + "48;2;" + "255;255;255" + Suffix + "test" + Prefix + Rst + Suffix}}
	for _, tt := range tsts {
		got := BgRGB(255, 255, 255, tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := BgRGB(255, 255, 255)
	if ttt != Prefix+"48;2;255;255;255"+Suffix {
		t.Errorf("Expected: %s48;5;255;255;255%s Got: %s", Prefix, Suffix, ttt)
	}
}

func TestReset(t *testing.T) {
	if Reset() != Prefix+Rst+Suffix {
		t.Errorf("Expected: %s%s%s Got: %s", Prefix, Rst, Suffix, Reset())
	}
}
