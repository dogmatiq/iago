package iago

import (
	"fmt"
	"io"
)

// MustWrite calls w.Write(buf) and returns the number of bytes written, or
// panics if an error occurs.
func MustWrite(w io.Writer, buf []byte) int {
	return Must(w.Write(buf))
}

// MustWriteByte calls w.Write([]byte{b}) and returns the number of bytes
// written or panics if an error occurs.
func MustWriteByte(w io.Writer, b byte) int {
	return Must(w.Write([]byte{b}))
}

// MustWriteString calls io.WriteString(w, s) and returns the number of bytes
// written or panics if an error occurs.
func MustWriteString(w io.Writer, s string) int {
	return Must(io.WriteString(w, s))
}

// MustWriteTo calls s.WriteTo(w) and returns the numner of bytes written, or
// panics if an error occurs.
func MustWriteTo(w io.Writer, s io.WriterTo) int {
	return Must64(s.WriteTo(w))
}

// MustFprintf calls fmt.Fprintf(w, f, v...) and returns the numner of bytes
// written, or panics if an error occurs.
func MustFprintf(w io.Writer, f string, v ...interface{}) int {
	return Must(fmt.Fprintf(w, f, v...))
}

// Recover recovers from a panic caused by one of the MustXXX() functions.
//
// It is intended to be used in a defer statement. The error that caused the
// panic is assigned to *err.
func Recover(err *error) {
	if err == nil {
		panic("err must be a non-nil pointer")
	}

	switch v := recover().(type) {
	case MustPanicSentinel:
		*err = v.Cause
	case nil:
		return
	default:
		panic(v)
	}
}

// MustPanicSentinel is a wrapper value used to identify panic's that are caused
// by one of the MustXXX() functions.
type MustPanicSentinel struct {
	// Cause is the error that caused the panic.
	Cause error
}

// Must panics if err is non-nil, otherwise it returns n.
func Must(n int, err error) int {
	if err != nil {
		panic(MustPanicSentinel{err})
	}

	return n
}

// Must64 panics if err is non-nil, otherwise it returns n.
func Must64(n int64, err error) int {
	if err != nil {
		panic(MustPanicSentinel{err})
	}

	return int(n)
}
