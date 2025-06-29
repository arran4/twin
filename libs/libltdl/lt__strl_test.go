package libltdl

import (
	"bytes"
	"testing"
)

func TestLtStrlcpyDstsizeExceedsSlice(t *testing.T) {
	dst := make([]byte, 5)
	n := LtStrlcpy(dst, "hello", 6)
	if n != 5 {
		t.Fatalf("expected length 5, got %d", n)
	}
	if !bytes.Equal(dst, make([]byte, 5)) {
		t.Fatalf("dst modified when dstsize exceeds slice")
	}
}

func TestLtStrlcatDstsizeExceedsSlice(t *testing.T) {
	dst := []byte{'a', 0}
	original := append([]byte(nil), dst...)
	n := LtStrlcat(dst, "bc", 4)
	if n != 3 {
		t.Fatalf("expected length 3, got %d", n)
	}
	if !bytes.Equal(dst, original) {
		t.Fatalf("dst modified when dstsize exceeds slice")
	}
}

func TestLtStrlcpyNormal(t *testing.T) {
	dst := make([]byte, 6)
	n := LtStrlcpy(dst, "hello", 6)
	if n != 5 {
		t.Fatalf("expected length 5, got %d", n)
	}
	if string(dst[:5]) != "hello" || dst[5] != 0 {
		t.Fatalf("unexpected dst contents: %v", dst)
	}
}

func TestLtStrlcatNormal(t *testing.T) {
	dst := []byte{'a', 0, 0, 0, 0}
	n := LtStrlcat(dst, "bc", 5)
	if n != 3 {
		t.Fatalf("expected length 3, got %d", n)
	}
	want := []byte{'a', 'b', 'c', 0, 0}
	if !bytes.Equal(dst, want) {
		t.Fatalf("unexpected dst contents: %v", dst)
	}
}
