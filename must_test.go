package iago_test

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/dogmatiq/iago"
	"github.com/dogmatiq/iago/iotest"
)

func TestMustWrite(t *testing.T) {
	expected := "foo"
	var b strings.Builder

	n := MustWrite(
		&b,
		[]byte(expected),
	)

	if b.String() != expected {
		t.Fatalf("unexpected content written: %s", b.String())
	}

	if n != len(expected) {
		t.Fatalf("expected number of bytes to be %d, not %d", len(expected), n)
	}
}

func TestMustWrite_PanicsAndRecovers(t *testing.T) {
	fn := func() (err error) {
		defer Recover(&err)

		MustWrite(
			iotest.NewFailer(nil, nil),
			[]byte("foo"),
		)

		return
	}

	if fn() != iotest.ErrWrite {
		t.Fatal("expected error did not occur")
	}
}

func TestMustWriteString(t *testing.T) {
	expected := "foo"
	var b strings.Builder

	n := MustWriteString(
		&b,
		expected,
	)

	if b.String() != expected {
		t.Fatalf("unexpected content written: %s", b.String())
	}

	if n != len(expected) {
		t.Fatalf("expected number of bytes to be %d, not %d", len(expected), n)
	}
}

func TestMustWriteString_PanicsAndRecovers(t *testing.T) {
	fn := func() (err error) {
		defer Recover(&err)

		MustWriteString(
			iotest.NewFailer(nil, nil),
			"foo",
		)

		return
	}

	if fn() != iotest.ErrWrite {
		t.Fatal("expected error did not occur")
	}
}

func TestMustWriteTo(t *testing.T) {
	expected := "foo"
	var b strings.Builder

	n := MustWriteTo(
		&b,
		bytes.NewBuffer([]byte(expected)),
	)

	if b.String() != expected {
		t.Fatalf("unexpected content written: %s", b.String())
	}

	if n != len(expected) {
		t.Fatalf("expected number of bytes to be %d, not %d", len(expected), n)
	}
}

func TestMustWriteTo_PanicsAndRecovers(t *testing.T) {
	fn := func() (err error) {
		defer Recover(&err)

		MustWriteTo(
			iotest.NewFailer(nil, nil),
			bytes.NewBuffer([]byte("foo")),
		)

		return
	}

	if fn() != iotest.ErrWrite {
		t.Fatal("expected error did not occur")
	}
}

func TestMustFprintf_PanicsAndRecovers(t *testing.T) {
	fn := func() (err error) {
		defer Recover(&err)

		MustFprintf(
			iotest.NewFailer(nil, nil),
			"foo %s",
			"bar",
		)

		return
	}

	if fn() != iotest.ErrWrite {
		t.Fatal("expected error did not occur")
	}
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
