package cursor

import "testing"

func TestUp(t *testing.T) {
	if Up(1) != prefix+"1"+up {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", up, Up(1))
	}
}

func TestDown(t *testing.T) {
	if Down(1) != prefix+"1"+down {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", down, Down(1))
	}
}

func TestRight(t *testing.T) {
	if Right(1) != prefix+"1"+right {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", right, Right(1))
	}
}

func TestLeft(t *testing.T) {
	if Left(1) != prefix+"1"+left {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", left, Left(1))
	}
}

func TestBeginNext(t *testing.T) {
	if BeginNext(1) != prefix+"1"+beginNext {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", beginNext, BeginNext(1))
	}
}

func TestBeginPrev(t *testing.T) {
	if BeginPrev(1) != prefix+"1"+beginPrev {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", beginPrev, BeginPrev(1))
	}
}

func TestColumn(t *testing.T) {
	if Column(1) != prefix+"1"+column {
		t.Errorf("Expected: %s%s%s Got: %s", prefix, "1", column, Column(1))
	}
}

func TestHome(t *testing.T) {
	if Home() != prefix+home {
		t.Errorf("Expected: %s%s Got: %s", prefix, home, Home())
	}
}

func TestPosition(t *testing.T) {
	if Position() != prefix+position {
		t.Errorf("Expected: %s%s Got: %s", prefix, position, Position())
	}
}

func TestMove(t *testing.T) {
	if Move() != prefix+move {
		t.Errorf("Expected: %s%s Got: %s", prefix, move, Move())
	}
}

func TestSave(t *testing.T) {
	sco := prefix + save
	dec := prefixDEC + decSave

	ttt := Save(false)
	if ttt != sco {
		t.Errorf("Expected: %s Got: %s", ttt, sco)
	}
	ttt = Save(false, false)
	if ttt != "" {
		t.Errorf("Expected: %s Got: %s", ttt, "")
	}
	ttt = Save(true, false)
	if ttt != sco {
		t.Errorf("Expected: %s Got: %s", ttt, dec)
	}
	ttt = Save()
	if ttt != sco+dec {
		t.Errorf("Expected: %s Got: %s", ttt, sco+dec)
	}
}

func TestRestore(t *testing.T) {
	sco := prefix + restore
	dec := prefixDEC + decRestore

	ttt := Restore(false)
	if ttt != sco {
		t.Errorf("Expected: %s Got: %s", ttt, sco)
	}
	ttt = Restore(false, false)
	if ttt != "" {
		t.Errorf("Expected: %s Got: %s", ttt, "")
	}
	ttt = Restore(true, false)
	if ttt != dec {
		t.Errorf("Expected: %s Got: %s", ttt, dec)
	}
	ttt = Restore()
	if ttt != sco+dec {
		t.Errorf("Expected: %s Got: %s", ttt, sco+dec)
	}
}

func TestMove2Coord(t *testing.T) {
	sco := prefix + "2;2" + "f"
	dec := prefix + "2;2" + home

	ttt := Move2Coord(2, 2)
	if ttt != sco {
		t.Errorf("Expected: %s Got: %s", ttt, dec)
	}
	ttt = Move2Coord(2, 2, true)
	if ttt != "" {
		t.Errorf("Expected: %s Got: %s", ttt, sco)
	}
}

func TestInvisible(t *testing.T) {
	if Invisible() != prefix+invisible {
		t.Errorf("Expected: %s%s Got: %s", prefix, invisible, Invisible())
	}
}

func TestVisible(t *testing.T) {
	if Visible() != prefix+visible {
		t.Errorf("Expected: %s%s Got: %s", prefix, visible, Visible())
	}
}
