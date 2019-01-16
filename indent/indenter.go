package indent

import (
	"bytes"
	"io"

	"github.com/dogmatiq/iago"
)

// DefaultIndent is the default index prefix to use.
var DefaultIndent = []byte("    ")

type indenter struct {
	w        io.Writer
	prefix   []byte
	indented bool
}

// NewIndenter returns an io.Writer that indents each line with a fixed prefix
// before forwarding to w.
//
// If p is empty, it defaults to DefaultIndent.
func NewIndenter(w io.Writer, p []byte) io.Writer {
	if len(p) == 0 {
		p = DefaultIndent
	}

	return &indenter{
		w:      w,
		prefix: p,
	}
}

func (w *indenter) Write(buf []byte) (n int, err error) {
	defer iago.Recover(&err)

	// keep writing so long as there's something in the buffer
	for len(buf) > 0 {
		// indent if we're ready to do so
		if !w.indented {
			n += iago.MustWrite(w.w, w.prefix)
			w.indented = true
		}

		// find the next line break character
		i := bytes.IndexByte(buf, '\n')

		// if there are no more line break characters, simply write the remainder of
		// the buffer and we're done
		if i == -1 {
			n += iago.MustWrite(w.w, buf)
			break
		}

		// otherwise, write the remainder of this line, including the line break
		// character, and trim the beginning of the buffer
		n += iago.MustWrite(w.w, buf[:i+1])
		buf = buf[i+1:]

		// we're ready for another indent if/when there is more content
		w.indented = false
	}

	return
}
