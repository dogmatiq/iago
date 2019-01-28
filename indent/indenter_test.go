package indent_test

import (
	"strings"

	"github.com/dogmatiq/iago"
	. "github.com/dogmatiq/iago/indent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Indenter", func() {
	It("writes indented text", func() {
		b := &strings.Builder{}
		w := NewIndenter(b, nil)

		n := iago.MustWriteString(w, "fo")
		n += iago.MustWriteString(w, "o\nb")
		n += iago.MustWriteString(w, "ar\n")
		n += iago.MustWriteString(w, "baz")

		Expect(
			b.String(),
		).To(Equal(
			"    foo\n    bar\n    baz",
		))

		// total length of input above
		Expect(n).To(Equal(11))
	})
})
