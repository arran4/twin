package libltdl

import "testing"

func TestLoaderRegistration(t *testing.T) {
	DlExit() // ensure clean state if previously initialized
	DlInit()
	defer DlExit()

	if DlloaderFind("lt_dlopen") == nil {
		t.Fatalf("default loader not registered")
	}

	vt := &Dlvtable{
		Name:        "test_loader",
		ModuleOpen:  func(data interface{}, filename string, advise interface{}) interface{} { return nil },
		ModuleClose: func(data interface{}, module interface{}) int { return 0 },
		FindSym:     func(data interface{}, module interface{}, symbolname string) interface{} { return nil },
		Priority:    DlloaderAppend,
	}
	if rc := DlloaderAdd(vt); rc != 0 {
		t.Fatalf("DlloaderAdd returned %d", rc)
	}
	if DlloaderFind("test_loader") != vt {
		t.Fatalf("DlloaderFind failed")
	}

	idx := DlloaderNext(-1)
	if idx != 0 {
		t.Fatalf("DlloaderNext(-1) = %d", idx)
	}
	idx = DlloaderNext(idx)
	if idx != 1 {
		t.Fatalf("DlloaderNext after first = %d", idx)
	}
	if DlloaderGet(idx) != vt {
		t.Fatalf("DlloaderGet did not return added loader")
	}

	if removed := DlloaderRemove("test_loader"); removed != vt {
		t.Fatalf("DlloaderRemove failed")
	}
	if DlloaderFind("test_loader") != nil {
		t.Fatalf("loader still present after remove")
	}
	if DlloaderNext(0) != -1 {
		t.Fatalf("unexpected next index after removal")
	}
}
