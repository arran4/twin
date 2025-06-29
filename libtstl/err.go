package libtstl

import "errors"

// ErrNum represents a common error code from the C++ library.
type ErrNum int

// Known error numbers corresponding to libs/libtstl/err.cpp.
const (
	Success ErrNum = iota
	NoMemory
	NoTables
	DlError
	SysError
	InvalError
)

var (
	lastErrNum ErrNum
	lastErrStr string
)

// SetError records an error number and returns it as a Go error.
func SetError(n ErrNum) error {
	lastErrNum = n
	switch n {
	case Success:
		lastErrStr = ""
	case NoMemory:
		lastErrStr = "Out of memory!"
	case NoTables:
		lastErrStr = "Internal tables full!"
	case DlError:
		lastErrStr = "Error loading dynamic library"
	case SysError:
		lastErrStr = "system error"
	case InvalError:
		lastErrStr = "Invalid value"
	default:
		lastErrStr = ""
	}
	if lastErrStr == "" {
		return nil
	}
	return errors.New(lastErrStr)
}

// SetSysError stores a system error and returns it as a Go error.
func SetSysError(err error) error {
	lastErrNum = SysError
	if err != nil {
		lastErrStr = err.Error()
	} else {
		lastErrStr = "system error"
	}
	return errors.New(lastErrStr)
}

// LastError returns the last recorded error string.
func LastError() string { return lastErrStr }

// LastErrNum returns the last recorded error code.
func LastErrNum() ErrNum { return lastErrNum }
