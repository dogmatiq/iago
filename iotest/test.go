package iotest

import (
	"io"
	"strings"
	"testing"

	"github.com/dogmatiq/iago"
)

// TestWrite runs subtests that perform success and failure tests of fn.
func TestWrite(
	t *testing.T,
	fn func(w io.Writer) int,
	l ...string,
) {
	t.Run(
		"success",
		func(t *testing.T) {
			TestWriteSuccess(t, fn, l...)
		},
	)

	t.Run(
		"failure",
		func(t *testing.T) {
			TestWriteFailure(t, fn)
		},
	)
}

// TestWriteSuccess calls fn() and verifies that it writes the lines of text in l to w.
func TestWriteSuccess(
	t *testing.T,
	fn func(w io.Writer) int,
	l ...string,
) {
	var w strings.Builder

	n, err := func() (n int, err error) {
		defer iago.Recover(&err)
		n = fn(&w)
		return
	}()

	if err != nil {
		t.Fatal(err)
	}

	expect := strings.Join(l, "\n")
	actual := w.String()

	if actual != expect {
		t.Fatalf(
			"unexpected content written:\n\n--- got ---\n\n%s\n\n--- want ---\n\n%s\n",
			actual,
			expect,
		)
	}

	if n != len(expect) {
		t.Fatalf(
			"unexpected number of bytes reported (%d), expected %d",
			n,
			len(expect),
		)
	}
}

// TestWriteFailure calls fn() and verifies that it propagates write errors
// caused by the underlying writer.
func TestWriteFailure(
	t *testing.T,
	fn func(w io.Writer) int,
) {
	w := NewFailer(nil, nil)

	err := func() (err error) {
		defer iago.Recover(&err)
		fn(w)
		return
	}()

	if err != ErrWrite {
		t.Fatal("expected error did not occur")
	}
}
