package libltdl

var (
	lastError        string
	userErrorStrings []string
)

var errorStrings = []string{
	"unknown error",
	"dlopen support not available",
	"invalid loader",
	"loader initialization failed",
	"loader removal failed",
	"file not found",
	"dependency library not found",
	"no symbols defined",
	"can't open the module",
	"can't close the module",
	"symbol not found",
	"not enough memory",
	"invalid module handle",
	"internal buffer overflow",
	"invalid errorcode",
	"library already shutdown",
	"can't close resident module",
	"internal error (code withdrawn)",
	"invalid search path insert position",
	"symbol visibility can be global or local",
}

// LtDladdError records a new error diagnostic and returns its index.
func LtDladdError(diagnostic string) int {
	userErrorStrings = append(userErrorStrings, diagnostic)
	return len(errorStrings) + len(userErrorStrings) - 1
}

// LtDlsetError sets lastError from the given error code.
func LtDlsetError(errindex int) int {
	if errindex < 0 || errindex >= len(errorStrings)+len(userErrorStrings) {
		lastError = "invalid errorcode"
		return 1
	}
	if errindex < len(errorStrings) {
		lastError = errorStrings[errindex]
	} else {
		lastError = userErrorStrings[errindex-len(errorStrings)]
	}
	return 0
}

// LtErrorString returns the fixed error string for errorcode.
func LtErrorString(errorcode int) string {
	if errorcode < 0 || errorcode >= len(errorStrings) {
		return ""
	}
	return errorStrings[errorcode]
}

// LtGetLastError returns the last error message.
func LtGetLastError() string { return lastError }

// LtSetLastError stores errormsg as the last error and returns it.
func LtSetLastError(errormsg string) string {
	lastError = errormsg
	return lastError
}

// DlError returns the most recent loader error message.
func DlError() string { return lastError }
