package count_test

import (
	"io/ioutil"
	"strings"

	"github.com/dogmatiq/iago"
	. "github.com/dogmatiq/iago/count"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Writer", func() {
	Describe("func Write", func() {
		It("forwards to the target writer", func() {
			b := &strings.Builder{}
			w := NewWriter(b)

			n := iago.MustWriteString(w, "foo")

			Expect(b.String()).To(Equal("foo"))
			Expect(n).To(Equal(3))
		})
	})

	Describe("func Count", func() {
		It("returns the number of bytes written", func() {
			w := NewWriter(ioutil.Discard)

			iago.MustWriteString(w, "foo")

			Expect(w.Count()).To(Equal(3))
		})
	})
})
