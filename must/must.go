package must

import (
	"fmt"
	"io"
)

// Write calls w.Write(buf) and returns the number of bytes written, or
// panics if an error occurs.
func Write(w io.Writer, buf []byte) int {
	return Must(w.Write(buf))
}

// WriteByte calls w.Write([]byte{b}) and returns the number of bytes
// written or panics if an error occurs.
func WriteByte(w io.Writer, b byte) int {
	return Must(w.Write([]byte{b}))
}

// WriteString calls io.WriteString(w, s) and returns the number of bytes
// written or panics if an error occurs.
func WriteString(w io.Writer, s string) int {
	return Must(io.WriteString(w, s))
}

// WriteTo calls s.WriteTo(w) and returns the numner of bytes written, or
// panics if an error occurs.
func WriteTo(w io.Writer, s io.WriterTo) int {
	return Must64(s.WriteTo(w))
}

// Fprintf calls fmt.Fprintf(w, f, v...) and returns the numner of bytes
// written, or panics if an error occurs.
func Fprintf(w io.Writer, f string, v ...interface{}) int {
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
	case PanicSentinel:
		*err = v.Cause
	case nil:
		return
	default:
		panic(v)
	}
}

// PanicSentinel is a wrapper value used to identify panic's that are caused
// by one of the MustXXX() functions.
type PanicSentinel struct {
	// Cause is the error that caused the panic.
	Cause error
}

// Must panics if err is non-nil, otherwise it returns n.
func Must(n int, err error) int {
	if err != nil {
		panic(PanicSentinel{err})
	}

	return n
}

// Must64 panics if err is non-nil, otherwise it returns n.
func Must64(n int64, err error) int {
	if err != nil {
		panic(PanicSentinel{err})
	}

	return int(n)
}
