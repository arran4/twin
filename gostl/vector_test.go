package gostl

import "testing"

func TestVectorReserveResize(t *testing.T) {
	var v Vector[int]
	v.Reserve(5)
	if v.Capacity() < 5 {
		t.Fatalf("expected capacity >=5, got %d", v.Capacity())
	}
	v.Resize(3)
	if v.Size() != 3 {
		t.Fatalf("expected size 3, got %d", v.Size())
	}
	v.Resize(0)
	if !v.Empty() {
		t.Fatalf("expected empty vector")
	}
}
