package gostl

import "testing"

func TestStringFormat(t *testing.T) {
	s := NewString(0)
	s.Formatf("%d-%s", 10, "test")
	if s.String() != "10-test" {
		t.Fatalf("unexpected formatted string %q", s.String())
	}
}

func TestStringAppendRune(t *testing.T) {
	s := NewString(0)
	s.AppendRune('Æ')
	s.MakeCString()
	if got := string(s.Data()); got != "Æ" {
		t.Fatalf("expected Æ, got %q", got)
	}
}
