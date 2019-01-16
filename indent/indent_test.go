package indent_test

import (
	"bytes"
	"testing"

	. "github.com/dogmatiq/iago/indent"
)

func TestString(t *testing.T) {
	expected := "    foo\n    bar\n    baz"
	result := String("foo\nbar\nbaz", "")

	if result != expected {
		t.Fatalf(
			"unexpected indentation:\n"+
				" got: %#v\n"+
				"want: %#v",
			result,
			expected,
		)
	}
}

func TestBytes(t *testing.T) {
	expected := []byte("    foo\n    bar\n    baz")
	result := Bytes([]byte("foo\nbar\nbaz"), nil)

	if !bytes.Equal(result, expected) {
		t.Fatalf(
			"unexpected indentation:\n"+
				" got: %#v\n"+
				"want: %#v",
			string(result),
			string(expected),
		)
	}
}
