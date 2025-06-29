package libtutf

// UTF32ToASCII approximates a rune using plain ASCII.
func UTF32ToASCII(r rune) byte {
	if r < 0x20 || r > 0x7e {
		r = cp437ToUTF32[UTF32ToCP437(r)]
		return cp437ToASCII[UTF32ToCP437(r)]
	}
	return byte(r)
}

var cp437Reverse map[rune]byte

// UTF32ToCP437 approximates a rune using codepage 437.
func UTF32ToCP437(r rune) byte {
	if r >= ' ' && r <= '~' || (r&^rune(0xff)) == 0xf000 || (r < 0x100 && cp437ToUTF32[r] == r) {
		return byte(r & 0xff)
	}
	if cp437Reverse == nil {
		cp437Reverse = make(map[rune]byte)
		for i, v := range cp437ToUTF32 {
			if rune(i) == v {
				continue
			}
			cp437Reverse[v] = byte(i)
		}
		// map CHECK MARK to SQUARE ROOT
		cp437Reverse[0x2713] = 0xFB
	}
	if b, ok := cp437Reverse[r]; ok {
		return b
	}
	return '?'
}

var cp865Reverse map[rune]byte

// UTF32ToCP865 approximates a rune using codepage 865.
func UTF32ToCP865(r rune) byte {
	if r >= ' ' && r <= '~' || (r&^rune(0xff)) == 0xf000 || (r < 0x100 && cp865ToUTF32[r] == r) {
		return byte(r & 0xff)
	}
	if cp865Reverse == nil {
		cp865Reverse = make(map[rune]byte)
		for i, v := range cp865ToUTF32 {
			if rune(i) == v {
				continue
			}
			cp865Reverse[v] = byte(i)
		}
		cp865Reverse[0x2713] = 0xFB
	}
	if b, ok := cp865Reverse[r]; ok {
		return b
	}
	return UTF32ToASCII(r)
}

// UTF32ToISO8859_1 approximates a rune using ISO-8859-1.
func UTF32ToISO8859_1(r rune) byte {
	if r < 0x100 || (r&^rune(0xff)) == 0xf000 {
		return byte(r & 0xff)
	}
	return UTF32ToASCII(r)
}
