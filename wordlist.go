// Package wordlist contains only its deflate method.
package wordlist

import (
	"bytes"
	"io"

	_ "embed"

	"github.com/ulikunitz/xz"
)

//go:embed wordlist.txt.xz
var wordlist []byte

// Deflate returns the uncompressed embed wordlist.
func Deflate() ([]byte, error) {
	r, err := xz.NewReader(bytes.NewReader(wordlist))

	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}
