package indent_test

import (
	"strings"
	"testing"

	"github.com/dogmatiq/iago"
	. "github.com/dogmatiq/iago/indent"
)

func TestIndenter(t *testing.T) {
	b := &strings.Builder{}
	w := NewIndenter(b, nil)

	n := iago.MustWriteString(w, "fo")
	n += iago.MustWriteString(w, "o\nb")
	n += iago.MustWriteString(w, "ar\n")
	n += iago.MustWriteString(w, "baz")

	expected := "    foo\n    bar\n    baz"

	result := b.String()
	if result != expected {
		t.Fatalf(
			"unexpected indentation:\n"+
				" got: %#v\n"+
				"want: %#v",
			result,
			expected,
		)
	}

	if n != len(expected) {
		t.Fatalf(
			"unexpected byte count:\n"+
				" got: %d\n"+
				"want: %d",
			n,
			len(expected),
		)
	}
}
