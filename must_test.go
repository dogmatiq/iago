package iago_test

import (
	"bytes"
	"io"
	"testing"

	. "github.com/dogmatiq/iago"
	"github.com/dogmatiq/iago/iotest"
)

func TestMustWrite(t *testing.T) {
	iotest.TestWrite(
		t,
		func(w io.Writer) int {
			return MustWrite(w, []byte("foo"))
		},
		"foo",
	)
}

func TestMustWriteString(t *testing.T) {
	iotest.TestWrite(
		t,
		func(w io.Writer) int {
			return MustWriteString(w, "foo")
		},
		"foo",
	)
}

func TestMustWriteTo(t *testing.T) {
	iotest.TestWrite(
		t,
		func(w io.Writer) int {
			return MustWriteTo(
				w,
				bytes.NewBuffer([]byte("foo")),
			)
		},
		"foo",
	)
}

func TestMustFprintf(t *testing.T) {
	iotest.TestWrite(
		t,
		func(w io.Writer) int {
			return MustFprintf(
				w,
				"foo %s",
				"bar",
			)
		},
		"foo bar",
	)
}

func TestRecover_DoesNotRecoverFromOtherPanics(t *testing.T) {
	var value interface{}

	fn := func() {
		defer func() {
			value = recover()
		}()

		fn := func() (err error) {
			defer Recover(&err)
			panic("<value>")
		}

		fn()
	}

	fn()

	if value != "<value>" {
		t.Fatal("panic value was unexpectedly suppressed")
	}
}

func TestRecover_DoesNotPanicWhenNoPanicOccurs(t *testing.T) {
	fn := func() (err error) {
		defer Recover(&err)
		return nil
	}

	if fn() != nil {
		t.Fatal("unexpected error was returned")
	}
}

func TestRecover_PanicsWhenPassedANilPointer(t *testing.T) {
	var value interface{}

	fn := func() {
		defer func() {
			value = recover()
		}()

		Recover(nil)
	}

	fn()

	if value != "err must be a non-nil pointer" {
		t.Fatal("Recover() unexpectedly accepted a nil pointer")
	}
}
