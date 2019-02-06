package count

import (
	"io"
	"sync/atomic"
)

// Reader is an io.Reader that forwards reads to another Reader, while
// maintaining a count of the number of bytes read.
type Reader struct {
	r io.Reader
	n int64 // atomic
}

// NewReader returns an io.Reader that forwards reads to another Reader, while
// maintaining a count of the number of bytes read.
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// Count returns the number of bytes read so far.
func (r *Reader) Count() int {
	return int(r.Count64())
}

// Count64 returns the number of bytes read so far as an int64.
func (r *Reader) Count64() int64 {
	return atomic.LoadInt64(&r.n)
}

func (r *Reader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	atomic.AddInt64(&r.n, int64(n))
	return n, err
}
