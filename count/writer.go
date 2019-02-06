package count

import (
	"io"
	"sync/atomic"
)

// Writer is an io.Writer that forwards writes to another writer, while
// maintaining a count of the number of bytes written.
type Writer struct {
	w io.Writer
	n int64 // atomic
}

// NewWriter returns an io.Writer that forwards writes to another writer, while
// maintaining a count of the number of bytes written.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Count returns the number of bytes written so far.
func (w *Writer) Count() int {
	return int(w.Count64())
}

// Count64 returns the number of bytes written so far as an int64.
func (w *Writer) Count64() int64 {
	return atomic.LoadInt64(&w.n)
}

func (w *Writer) Write(buf []byte) (int, error) {
	n, err := w.w.Write(buf)
	atomic.AddInt64(&w.n, int64(n))
	return n, err
}
