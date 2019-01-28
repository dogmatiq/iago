package iago_test

import (
	"bytes"
	"io"

	. "github.com/dogmatiq/iago"
	"github.com/dogmatiq/iago/iotest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func MustWrite", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return MustWrite(w, []byte("foo"))
			},
			"foo",
		)
	})
})

var _ = Describe("func MustWriteByte", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return MustWriteByte(w, 'f')
			},
			"f",
		)
	})
})

var _ = Describe("func MustWriteString", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return MustWriteString(w, "foo")
			},
			"foo",
		)
	})
})

var _ = Describe("func MustWriteTo", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return MustWriteTo(
					w,
					bytes.NewBuffer([]byte("foo")),
				)
			},
			"foo",
		)
	})
})

var _ = Describe("func MustFprintf", func() {
	It("passes the IO test", func() {
		iotest.TestWrite(
			GinkgoT(),
			func(w io.Writer) int {
				return MustFprintf(
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
