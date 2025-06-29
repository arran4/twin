package libtutf

import "testing"

func TestASCII(t *testing.T) {
	if asciiToUTF32['A'] != 'A' {
		t.Fatalf("A mapping wrong: %x", asciiToUTF32['A'])
	}
	if UTF32ToASCII('A') != 'A' {
		t.Fatalf("UTF32ToASCII failed")
	}
}

func TestCP437(t *testing.T) {
	if cp437ToUTF32[0x80] != 0x00C7 {
		t.Fatalf("cp437 mapping wrong")
	}
	if UTF32ToCP437(0x00C7) != 0x80 {
		t.Fatalf("reverse mapping wrong")
	}
}

func TestISO8859_1(t *testing.T) {
	if iso8859_1ToUTF32[0xA3] != 0x00A3 {
		t.Fatalf("iso mapping wrong")
	}
	if UTF32ToISO8859_1('£') != 0xA3 {
		t.Fatalf("reverse iso wrong")
	}
}
