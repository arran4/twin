package libtw

import "fmt"

// Error codes from Twerrno.h
const (
	TW_EBAD_SIZES              = 1
	TW_EBAD_STRUCTS            = 2
	TW_ENO_DISPLAY             = 4
	TW_EBAD_DISPLAY            = 5
	TW_ENO_HOST                = 6
	TW_ENO_AUTH                = 7
	TW_ESYS_NO_MEM             = 8
	TW_ESYS_NO_SOCKET          = 9
	TW_ESYS_CANNOT_CONNECT     = 10
	TW_ESYS_CANNOT_WRITE       = 11
	TW_ESERVER_BAD_VERSION     = 12
	TW_ESERVER_BAD_PROTOCOL    = 13
	TW_ESERVER_BAD_ENDIAN      = 14
	TW_ESERVER_BAD_SIZES       = 15
	TW_ESERVER_LOST_CONNECT    = 16
	TW_ESERVER_ALREADY_CONNECT = 17
	TW_ESERVER_DENIED_CONNECT  = 18
	TW_ESERVER_NO_FUNCTION     = 19
	TW_ESERVER_BAD_FUNCTION    = 20
	TW_ESERVER_BAD_RETURN      = 21
	TW_EGZIP_BAD_PROTOCOL      = 22
	TW_EGZIP_INTERNAL          = 23
	TW_ECALL_BAD               = 24
	TW_ECALL_BAD_ARG           = 25
	TW_ESERVER_READ_TIMEOUT    = 26
)

const (
	TW_EDETAIL_NO_MODULE = 1
)

type TwErrno struct {
	E uint32
	S uint32
}

var commonErrno TwErrno

type TwDisplay struct {
	errno           TwErrno
	listeners       map[uint64][]*TwListener
	DefaultListener func(msg *TwMsg, arg interface{})
	DefaultArg      interface{}
}

func (d *TwDisplay) init() {
	if d.listeners == nil {
		d.listeners = make(map[uint64][]*TwListener)
	}
}

func Tw_ErrnoLocation(d *TwDisplay) *TwErrno {
	if d == nil {
		return &commonErrno
	}
	return &d.errno
}

func Tw_StrError(d *TwDisplay, e uint32) string {
	switch e {
	case 0:
		return "success"
	case TW_ESERVER_BAD_ENDIAN:
		return "server has reversed endianity, impossible to connect"
	case TW_ESERVER_BAD_SIZES:
		return "server has different data sizes, impossible to connect"
	case TW_EBAD_SIZES:
		return "compiled data sizes are incompatible with libtw now in use!"
	case TW_EBAD_STRUCTS:
		return "internal error: structs are not packed! Please contact the author."
	case TW_ENO_DISPLAY:
		return "TWDISPLAY is not set"
	case TW_EBAD_DISPLAY:
		return "badly formed TWDISPLAY"
	case TW_ENO_AUTH:
		return "bad or missing authorization file ~/.TwinAuth, cannot connect"
	case TW_ESYS_CANNOT_CONNECT:
		return "failed to connect: "
	case TW_ESYS_NO_MEM:
		return "out of memory!"
	case TW_ESYS_CANNOT_WRITE:
		return "failed to send data to server: "
	case TW_ESYS_NO_SOCKET:
		return "failed to create socket: "
	case TW_ESERVER_LOST_CONNECT:
		return "connection lost "
	case TW_ESERVER_ALREADY_CONNECT:
		return "already connected"
	case TW_ESERVER_BAD_PROTOCOL:
		return "got invalid data from server, protocol violated"
	case TW_ESERVER_NO_FUNCTION:
		return "function not supported by server: "
	case TW_ESERVER_BAD_FUNCTION:
		return "function is not a possible server function"
	case TW_ESERVER_DENIED_CONNECT:
		return "server denied permission to connect, file ~/.TwinAuth may be wrong"
	case TW_EGZIP_BAD_PROTOCOL:
		return "got invalid data from server, gzip format violated"
	case TW_EGZIP_INTERNAL:
		return "internal gzip error, panic!"
	case TW_ENO_HOST:
		return "unknown host in TWDISPLAY: "
	case TW_ESERVER_BAD_VERSION:
		return "server has incompatible protocol version, impossible to connect"
	case TW_ESERVER_BAD_RETURN:
		return "server function call returned strange data, wrong data sizes? : "
	case TW_ECALL_BAD:
		return "function call rejected by server, wrong data sizes? : "
	case TW_ECALL_BAD_ARG:
		return "function call rejected by server, invalid arguments? : "
	case TW_ESERVER_READ_TIMEOUT:
		return "failed to receive data from server: read timeout"
	default:
		return "unknown error"
	}
}

func Tw_StrErrorDetail(d *TwDisplay, E, S uint32) string {
	switch E {
	case TW_ESERVER_LOST_CONNECT:
		switch S {
		case TW_EDETAIL_NO_MODULE:
			return "(socket module may be not running on server)"
		default:
			break
		}
		return "(explicit kill or server shutdown)"
	case TW_ESYS_CANNOT_CONNECT, TW_ESYS_CANNOT_WRITE, TW_ESYS_NO_SOCKET:
		return fmt.Sprintf("%v", S)
	case TW_ENO_HOST:
		return fmt.Sprintf("%v", S)
	case TW_ESERVER_NO_FUNCTION, TW_ESERVER_BAD_RETURN, TW_ECALL_BAD, TW_ECALL_BAD_ARG:
		return fmt.Sprintf("%d", S)
	}
	return ""
}
