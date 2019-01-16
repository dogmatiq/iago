package iotest

import (
	"errors"
	"io"
)

var (
	// ErrRead is the default read error used by NewFailer().
	ErrRead = errors.New("<induced read error>")

	// ErrWrite is the default write error used by NewFailer().
	ErrWrite = errors.New("<induced write error>")
)

// failer is an io.ReadWriter that always fails.
type failer struct {
	read  error
	write error
}

// NewFailer returns an io.ReadWriter that always fails.
//
// r and w are the errors returned by Read() and Write(), respectively.
// If either are nil they default to ErrRead and ErrWrite, respectively.
func NewFailer(r, w error) io.ReadWriter {
	if r == nil {
		r = ErrRead
	}

	if w == nil {
		w = ErrWrite
	}

	return failer{r, w}
}

func (f failer) Read([]byte) (int, error) {
	return 0, f.read
}

func (f failer) Write([]byte) (int, error) {
	return 0, f.write
}
