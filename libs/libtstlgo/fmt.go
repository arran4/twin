package libtstl

// FmtBase represents a value that can be formatted to a byte slice.
type FmtBase interface {
	WriteTo([]byte) ToCharsResult
	Size() int
}

// FmtVoid outputs nothing.
type FmtVoid struct{}

func (FmtVoid) WriteTo(dst []byte) ToCharsResult { return ToCharsResult{0, SUCCESS} }
func (FmtVoid) Size() int                        { return 0 }

// FmtLong formats signed integers.
type FmtLong struct {
	Val  int64
	Base int
}

func (f FmtLong) WriteTo(dst []byte) ToCharsResult { return ToChars(dst, f.Val, f.Base) }
func (f FmtLong) Size() int                        { return ToCharsLenSigned(f.Val, f.Base) }

// FmtUlong formats unsigned integers.
type FmtUlong struct {
	Val  uint64
	Base int
}

func (f FmtUlong) WriteTo(dst []byte) ToCharsResult { return ToCharsUnsigned(dst, f.Val, f.Base) }
func (f FmtUlong) Size() int                        { return ToCharsLenUnsigned(f.Val, f.Base) }

// FmtBytes copies bytes as-is.
type FmtBytes struct{ Val []byte }

func (f FmtBytes) WriteTo(dst []byte) ToCharsResult { return ToCharsCopy(dst, f.Val) }
func (f FmtBytes) Size() int                        { return len(f.Val) }
