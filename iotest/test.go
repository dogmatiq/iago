package iotest

import (
	"io"
	"strings"

	"github.com/dogmatiq/iago"
)

// TestingT is the interface via which Iago's test functions consume Go's
// *testing.T type.
type TestingT interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// TestWrite runs subtests that perform success and failure tests of fn.
func TestWrite(
	t TestingT,
	fn func(w io.Writer) int,
	l ...string,
) {
	TestWriteSuccess(t, fn, l...)
	TestWriteFailure(t, fn)
}

// TestWriteSuccess calls fn() and verifies that it writes the lines of text in l to w.
func TestWriteSuccess(
	t TestingT,
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
	t TestingT,
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
