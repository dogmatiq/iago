package count_test

import (
	"io/ioutil"
	"strings"

	"github.com/dogmatiq/iago/must"

	. "github.com/dogmatiq/iago/count"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Writer", func() {
	Describe("func Write", func() {
		It("forwards to the target writer", func() {
			b := &strings.Builder{}
			w := NewWriter(b)

			n := must.WriteString(w, "foo")

			Expect(b.String()).To(Equal("foo"))
			Expect(n).To(Equal(3))
		})
	})

	Describe("func Count", func() {
		It("returns the number of bytes written", func() {
			w := NewWriter(ioutil.Discard)

			must.WriteString(w, "foo")

			Expect(w.Count()).To(Equal(3))
		})
	})

	Describe("func Count64", func() {
		It("returns the number of bytes written", func() {
			w := NewWriter(ioutil.Discard)

			must.WriteString(w, "foo")

			Expect(w.Count64()).To(Equal(int64(3)))
		})
	})
})
