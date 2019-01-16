package indent

import (
	"bytes"
	"strings"

	"github.com/dogmatiq/iago"
)

// String returns a copy of s with each line indented by p.
// If p is empty, it defaults to DefaultIndent.
func String(s, p string) string {
	var b strings.Builder
	w := NewIndenter(&b, []byte(p))
	iago.MustWriteString(w, s)
	return b.String()
}

// Bytes returns a copy of buf with each line indented by p.
// If p is empty, it defaults to DefaultIndent.
func Bytes(buf, p []byte) []byte {
	var b bytes.Buffer
	w := NewIndenter(&b, p)
	iago.MustWrite(w, buf)
	return b.Bytes()
}
