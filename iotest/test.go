package iotest

import (
	"io"
	"strings"

	"github.com/dogmatiq/iago/must"
)

// TestingT is the interface via which Iago's test functions consume Go's
// *testing.T type.
type TestingT interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// TestWriteSuccess is a utility for testing that code that writes to an
// io.Writer, such as an io.WriterTo implementation, produces the correct content
// and byte counts.
//
// fn() is a function that writes to w and returns the number of bytes it writes.
// It can use the iago.MustXXX() functions to simplify error propagation and
// byte counting.
//
// The test fails immediately if the actual content written does not match the
// given lines of text, or the byte count returned by fn() does not equal the
// actual number of bytes written.
func TestWriteSuccess(
	t TestingT,
	fn func(w io.Writer) int,
	lines ...string,
) {
	var w strings.Builder

	n, err := func() (n int, err error) {
		defer must.Recover(&err)
		n = fn(&w)
		return
	}()

	if err != nil {
		t.Fatal(err)
	}

	expect := strings.Join(lines, "\n")
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

// TestWriteFailure is a utility for testing that code that writes to an
// io.Writer, such as an io.WriterTo implementation, propagates errors that are
// caused by that writer.
//
// fn() is a function that writes to w. It can use the iago.MustXXX() propagate
// errors. w is a writer that intentionally fails.
//
// The test fails immediately if fn() does not propagate the error caused by w.
func TestWriteFailure(
	t TestingT,
	fn func(w io.Writer),
) {
	w := NewFailer(nil, nil)

	err := func() (err error) {
		defer must.Recover(&err)
		fn(w)
		return
	}()

	if err != ErrWrite {
		t.Fatal("expected error did not occur")
	}
}

// TestWrite is a convenience method for running a TestWriteSuccess() and
// TestWriteFailure() test against fn.
func TestWrite(
	t TestingT,
	fn func(w io.Writer) int,
	lines ...string,
) {
	TestWriteSuccess(t, fn, lines...)
	TestWriteFailure(t, func(w io.Writer) { fn(w) })
}
