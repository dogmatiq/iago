package iotest_test

import (
	"testing"

	. "github.com/dogmatiq/iago/iotest"
)

func TestFailer_Read(t *testing.T) {

	f := NewFailer(nil, nil)
	n, err := f.Read(nil)

	if err != ErrRead {
		t.Fatal("expected error did not occur")
	}

	if n != 0 {
		t.Fatalf("expected number of bytes to be zero, not %d", n)
	}
}

func TestFailer_Write(t *testing.T) {
	f := NewFailer(nil, nil)
	n, err := f.Write(nil)

	if err != ErrWrite {
		t.Fatal("expected error did not occur")
	}

	if n != 0 {
		t.Fatalf("expected number of bytes to be zero, not %d", n)
	}
}
