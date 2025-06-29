package server

// Go reimplementation of data structures from data.cpp.

// RGB holds red, green and blue components of a color.
type RGB struct {
	Red, Green, Blue byte
}

// KeyList mirrors the C++ keylist struct.
type KeyList struct {
	Name string
	Key  uint16
	Len  byte
	Seq  string
}

const (
	tblack   = 0
	tblue    = 1
	tgreen   = 2
	tred     = 4
	tcyan    = tblue | tgreen
	tmagenta = tblue | tred
	tyellow  = tgreen | tred
	twhite   = tblue | tgreen | tred
	thigh    = 8
	tmaxcol  = 0xF
)

// default directories used by the server
var (
	Plugindir = "/usr/local/lib/twin"
	Confdir   = "/usr/local/etc/twin"
)

func tcol(fg, bg byte) byte { return fg | bg<<4 }

// Palette contains the default colour table.
var Palette = [tmaxcol + 1]RGB{
	{0, 0, 0}, {0, 0, 0xAA}, {0, 0xAA, 0}, {0, 0xAA, 0xAA},
	{0xAA, 0, 0}, {0xAA, 0, 0xAA}, {0xAA, 0xAA, 0}, {0xAA, 0xAA, 0xAA},
	{0x55, 0x55, 0x55}, {0x55, 0x55, 0xFF}, {0x55, 0xFF, 0x55}, {0x55, 0xFF, 0xFF},
	{0xFF, 0x55, 0x55}, {0xFF, 0x55, 0xFF}, {0xFF, 0xFF, 0x55}, {0xFF, 0xFF, 0xFF},
}

// DefaultPalette duplicates Palette so it can be restored later.
var DefaultPalette = Palette

var (
	GadgetResize = [2]rune{0xCD, 0xBC}
	ScrollBarX   = [3]rune{0xB1, 0x11, 0x10}
	ScrollBarY   = [3]rune{0xB1, 0x1E, 0x1F}
	TabX         = rune(0xDB)
	TabY         = rune(0xDB)
	StdBorder    = [2][9]rune{{0xC9, 0xCD, 0xBB, 0xBA, 0x20, 0xBA, 0xC8, 0xCD, 0xBC},
		{0xDA, 0xC4, 0xBF, 0xB3, 0x20, 0xB3, 0xC0, 0xC4, 0xD9}}
	ScreenBack = [2]rune{0x12, 0x12}
)

var (
	DefaultColGadgets        = tcol(thigh|tyellow, tcyan)
	DefaultColArrows         = tcol(thigh|tgreen, thigh|tblue)
	DefaultColBars           = tcol(twhite, thigh|tblue)
	DefaultColTabs           = tcol(thigh|twhite, thigh|tblue)
	DefaultColBorder         = tcol(thigh|twhite, thigh|tblue)
	DefaultColDisabled       = tcol(thigh|tblack, tblack)
	DefaultColSelectDisabled = tcol(tblack, thigh|tblack)
)

// InitData performs setup similar to the C++ version.
func InitData() bool {
	// TODO: translate runes from CP437 to UTF-32 once charset tables are available.
	return true
}
