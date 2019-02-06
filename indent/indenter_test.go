package indent_test

import (
	"strings"

	. "github.com/dogmatiq/iago/indent"
	"github.com/dogmatiq/iago/must"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Indenter", func() {
	It("writes indented text", func() {
		b := &strings.Builder{}
		w := NewIndenter(b, nil)

		n := must.WriteString(w, "fo")
		n += must.WriteString(w, "o\nb")
		n += must.WriteString(w, "ar\n")
		n += must.WriteString(w, "baz")

		Expect(
			b.String(),
		).To(Equal(
			"    foo\n    bar\n    baz",
		))

		// total length of input above
		Expect(n).To(Equal(11))
	})
})
