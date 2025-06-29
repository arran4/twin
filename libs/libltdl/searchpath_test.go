package libltdl

import "testing"

func TestSearchPathHandling(t *testing.T) {
	DlExit()
	DlInit()
	defer DlExit()

	if sp := DlGetSearchPath(); sp != "" {
		t.Fatalf("expected empty path, got %q", sp)
	}

	DlSetSearchPath("a:b")
	if sp := DlGetSearchPath(); sp != "a:b" {
		t.Fatalf("after set got %q", sp)
	}

	DlAddSearchDir("c")
	if sp := DlGetSearchPath(); sp != "a:b:c" {
		t.Fatalf("after add got %q", sp)
	}

	DlInsertSearchDir("b", "x")
	if sp := DlGetSearchPath(); sp != "a:x:b:c" {
		t.Fatalf("after insert got %q", sp)
	}

	DlInsertSearchDir("missing", "y")
	if sp := DlGetSearchPath(); sp != "a:x:b:c:y" {
		t.Fatalf("after append got %q", sp)
	}

	DlSetSearchPath("")
	if sp := DlGetSearchPath(); sp != "" {
		t.Fatalf("after reset got %q", sp)
	}
}
