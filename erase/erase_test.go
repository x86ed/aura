package erase

import "testing"

func TestToScreenEnd(t *testing.T) {
	if ToScreenEnd() != prefix+"0"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "0", sSuffix, ToScreenEnd())
	}
}

func TestToScreenBegin(t *testing.T) {
	if ToScreenBegin() != prefix+"1"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", sSuffix, ToScreenBegin())
	}
}

func TestScreen(t *testing.T) {
	if Screen() != prefix+"2"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "2", sSuffix, Screen())
	}
}

func TestSavedLines(t *testing.T) {
	if SavedLines() != prefix+"3"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "3", sSuffix, SavedLines())
	}
}

func TestToLineEnd(t *testing.T) {
	if ToLineEnd() != prefix+"0"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "0", sSuffix, ToLineEnd())
	}
}

func TestToLineBegin(t *testing.T) {
	if ToLineStart() != prefix+"1"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", sSuffix, ToLineStart())
	}
}

func TestLine(t *testing.T) {
	if Line() != prefix+"2"+sSuffix {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "2", sSuffix, Line())
	}
}
