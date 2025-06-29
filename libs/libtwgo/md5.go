package libtw

import (
	"crypto/md5"
	"hash"
)

// MD5Uint32 is a 32-bit unsigned integer type used by the MD5 implementation.
type MD5Uint32 = uint32

// MD5Context holds the context for MD5 operations.
type MD5Context struct {
	h hash.Hash
}

// MD5_CTX is kept for compatibility with the original C API.
type MD5_CTX = MD5Context

// MD5Init initializes the given context for a new MD5 computation.
func MD5Init(ctx *MD5Context) {
	ctx.h = md5.New()
}

// MD5Update adds data to the MD5 context.
func MD5Update(ctx *MD5Context, data []byte) {
	if ctx.h == nil {
		ctx.h = md5.New()
	}
	ctx.h.Write(data)
}

// MD5Final returns the final MD5 checksum and resets the context.
func MD5Final(digest *[16]byte, ctx *MD5Context) {
	if ctx.h == nil {
		ctx.h = md5.New()
	}
	sum := ctx.h.Sum(nil)
	copy(digest[:], sum)
	*ctx = MD5Context{}
}
