package transfer

import (
	"crypto/md5"
	"hash"
	"io"
)

// MD5Reader is responsible for reading from stdin
type MD5Reader struct {
	input   io.Reader
	hashing bool
	md5     hash.Hash
}

// compile time check
var _ io.Reader = &MD5Reader{}

func (i *MD5Reader) Read(p []byte) (n int, err error) {
	n, err = i.input.Read(p)
	if i.hashing {
		i.md5.Write(p)
	}
	return n, err
}

// Sum returns the md5 of all congested data
func (i *MD5Reader) Sum() []byte {
	return i.md5.Sum(nil)
}

// Hashing tells wether data is hashed
func (i *MD5Reader) Hashing() bool {
	return i.hashing
}

// NewMD5Reader create a new InputReader
func NewMD5Reader(input io.Reader, hash bool) *MD5Reader {
	return &MD5Reader{
		input:   input,
		hashing: hash,
		md5:     md5.New(),
	}
}
