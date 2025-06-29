package libltdl

import (
	"os"
	"reflect"
	"testing"
)

func TestArgzHelpers(t *testing.T) {
	if v := ArgzCreateSep("a:b::c", ':'); !reflect.DeepEqual(v, []string{"a", "b", "c"}) {
		t.Fatalf("ArgzCreateSep unexpected %v", v)
	}

	argz := []string{"a", "b"}
	ArgzAppend(&argz, "c")
	if !reflect.DeepEqual(argz, []string{"a", "b", "c"}) {
		t.Fatalf("ArgzAppend result %v", argz)
	}

	ArgzInsert(&argz, "b", "x")
	if !reflect.DeepEqual(argz, []string{"a", "x", "b", "c"}) {
		t.Fatalf("ArgzInsert result %v", argz)
	}

	ArgzInsert(&argz, "missing", "y")
	if !reflect.DeepEqual(argz, []string{"a", "x", "b", "c", "y"}) {
		t.Fatalf("ArgzInsert append %v", argz)
	}

	if next := ArgzNext(argz, "a"); next != "x" {
		t.Fatalf("ArgzNext returned %q", next)
	}
	if first := ArgzNext(argz, ""); first != "a" {
		t.Fatalf("ArgzNext first %q", first)
	}

	if s := ArgzStringify(argz, ':'); s != "a:x:b:c:y" {
		t.Fatalf("ArgzStringify %q", s)
	}

	ordered := []string{"b", "d"}
	ArgzInsertInOrder(&ordered, "c")
	ArgzInsertInOrder(&ordered, "d")
	ArgzInsertInOrder(&ordered, "a")
	if !reflect.DeepEqual(ordered, []string{"a", "b", "c", "d"}) {
		t.Fatalf("ArgzInsertInOrder %v", ordered)
	}

	dir := t.TempDir()
	os.WriteFile(dir+"/b.so", []byte{}, 0o644)
	os.WriteFile(dir+"/a1.so", []byte{}, 0o644)
	os.WriteFile(dir+"/.hidden", []byte{}, 0o644)
	os.WriteFile(dir+"/c.txt", []byte{}, 0o644)
	var list []string
	if rc := ArgzInsertDir(&list, dir); rc != 0 {
		t.Fatalf("ArgzInsertDir rc=%d", rc)
	}
	if !reflect.DeepEqual(list, []string{dir + "/a", dir + "/b", dir + "/c"}) {
		t.Fatalf("ArgzInsertDir %v", list)
	}
}
