package must_test

import (
	"bytes"
	"io"

	"github.com/dogmatiq/iago/iotest"
	. "github.com/dogmatiq/iago/must"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Write", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return Write(w, []byte("foo"))
			},
			"foo",
		)
	})
})

var _ = Describe("func WriteByte", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return WriteByte(w, 'f')
			},
			"f",
		)
	})
})

var _ = Describe("func WriteString", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return WriteString(w, "foo")
			},
			"foo",
		)
	})
})

var _ = Describe("func WriteTo", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return WriteTo(
					w,
					bytes.NewBuffer([]byte("foo")),
				)
			},
			"foo",
		)
	})
})

var _ = Describe("func Fprintf", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return Fprintf(
					w,
					"foo %s",
					"bar",
				)
			},
			"foo bar",
		)
	})
})

var _ = Describe("func Recover", func() {
	It("does not recover from unrelated panics", func() {
		var value interface{}

		func() {
			defer func() {
				value = recover()
			}()

			func() (err error) {
				defer Recover(&err)
				panic("<value>") // not a MustPanicSentinel
			}()
		}()

		Expect(value).To(Equal("<value>"))
	})

	It("does not panic when no panic occurs", func() {
		err := func() (err error) {
			defer Recover(&err)
			return nil
		}()

		Expect(err).ShouldNot(HaveOccurred())
	})

	It("panics when passed a nil pointer", func() {
		Expect(func() {
			Recover(nil)
		}).To(Panic())
	})
})
