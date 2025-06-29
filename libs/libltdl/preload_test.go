package libltdl

import "testing"

func TestPreloadLogic(t *testing.T) {
	DlExit()
	DlInit()
	defer DlExit()

	defaults := []PreloadSym{{Name: "main"}, {Name: "m1"}, {Name: "m2"}}
	DlPreloadDefault(defaults)
	DlPreload(nil)

	if rc := DlPreloadOpen("main", nil); rc != 2 {
		t.Fatalf("expected 2 errors, got %d", rc)
	}

	extra := []PreloadSym{{Name: "other"}, {Name: "extra"}}
	DlPreload(extra)

	if rc := DlPreloadOpen("", nil); rc != 3 {
		t.Fatalf("expected 3 errors, got %d", rc)
	}
}
