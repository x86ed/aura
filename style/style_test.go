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
