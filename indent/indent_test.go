package indent_test

import (
	. "github.com/dogmatiq/iago/indent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func String", func() {
	It("returns an indented string", func() {
		Expect(
			String("foo\nbar\nbaz", ""),
		).To(Equal(
			"    foo\n    bar\n    baz",
		))
	})
})

var _ = Describe("func Bytes", func() {
	It("returns an indented byte-slice", func() {
		Expect(
			Bytes([]byte("foo\nbar\nbaz"), nil),
		).To(Equal(
			[]byte("    foo\n    bar\n    baz"),
		))
	})
})
