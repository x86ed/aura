package style

import "testing"

type tst struct {
	input, expected string
}

func TestBold(t *testing.T) {
	tsts := []tst{{"test", prefix + bold + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Bold(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Bold()
	if ttt != prefix+bold+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, bold, suffix, ttt)
	}
}

func TestLight(t *testing.T) {
	tsts := []tst{{"test", prefix + light + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Light(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Light()
	if ttt != prefix+light+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, light, suffix, ttt)
	}
}

func TestItalic(t *testing.T) {
	tsts := []tst{{"test", prefix + italic + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Italic(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Italic()
	if ttt != prefix+italic+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, italic, suffix, ttt)
	}
}

func TestUnderline(t *testing.T) {
	tsts := []tst{{"test", prefix + underline + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Underline(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Underline()
	if ttt != prefix+underline+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, underline, suffix, ttt)
	}
}

func TestBlink(t *testing.T) {
	tsts := []tst{{"test", prefix + blink + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Blink(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Blink()
	if ttt != prefix+blink+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, blink, suffix, ttt)
	}
}

func TestBold2(t *testing.T) {
	tsts := []tst{{"test", prefix + bold2 + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Bold2(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Bold2()
	if ttt != prefix+bold2+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, bold2, suffix, ttt)
	}
}

func TestInvert(t *testing.T) {
	tsts := []tst{{"test", prefix + invert + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Invert(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Invert()
	if ttt != prefix+invert+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, invert, suffix, ttt)
	}
}

func TestInvisible(t *testing.T) {
	tsts := []tst{{"test", prefix + invisible + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Invisible(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Invisible()
	if ttt != prefix+invisible+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, invisible, suffix, ttt)
	}
}

func TestCross(t *testing.T) {
	tsts := []tst{{"test", prefix + cross + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Cross(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Cross()
	if ttt != prefix+cross+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, cross, suffix, ttt)
	}
}

func TestDoubleUnderline(t *testing.T) {
	tsts := []tst{{"test", prefix + doubleUnderline + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := DoubleUnderline(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := DoubleUnderline()
	if ttt != prefix+doubleUnderline+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, doubleUnderline, suffix, ttt)
	}
}

func TestNormal(t *testing.T) {
	tsts := []tst{{"test", prefix + normal + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Normal(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Normal()
	if ttt != prefix+normal+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, normal, suffix, ttt)
	}
}

func TestUnItalic(t *testing.T) {
	tsts := []tst{{"test", prefix + unItalic + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := UnItalic(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := UnItalic()
	if ttt != prefix+unItalic+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, unItalic, suffix, ttt)
	}
}

func TestUnUnderlined(t *testing.T) {
	tsts := []tst{{"test", prefix + unUnderlined + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := UnUnderlined(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := UnUnderlined()
	if ttt != prefix+unUnderlined+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, unUnderlined, suffix, ttt)
	}
}

func TestSteady(t *testing.T) {
	tsts := []tst{{"test", prefix + steady + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Steady(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Steady()
	if ttt != prefix+steady+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, steady, suffix, ttt)
	}
}

func TestPositive(t *testing.T) {
	tsts := []tst{{"test", prefix + positive + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Positive(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Positive()
	if ttt != prefix+positive+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, positive, suffix, ttt)
	}
}

func TestVisible(t *testing.T) {
	tsts := []tst{{"test", prefix + visible + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := Visible(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := Visible()
	if ttt != prefix+visible+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, visible, suffix, ttt)
	}
}

func TestUnCross(t *testing.T) {
	tsts := []tst{{"test", prefix + uncross + suffix + "test" + prefix + reset + suffix}}
	for _, tt := range tsts {
		got := UnCross(tt.input)
		if got != tt.expected {
			t.Errorf("Expected: %s Got: %s", tt.expected, got)
		}
	}
	ttt := UnCross()
	if ttt != prefix+uncross+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, uncross, suffix, ttt)
	}
}

func TestReset(t *testing.T) {
	if Reset() != prefix+reset+suffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, reset, suffix, Reset())
	}
}
