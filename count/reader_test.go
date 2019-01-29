package count_test

import (
	"strings"

	. "github.com/dogmatiq/iago/count"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Reader", func() {
	Describe("func Read", func() {
		It("reads from the source reader", func() {
			s := strings.NewReader("foo")
			r := NewReader(s)

			buf := make([]byte, 10)

			n, err := r.Read(buf)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(n).To(Equal(3))

		})
	})

	Describe("func Count", func() {
		It("returns the number of bytes read", func() {
			s := strings.NewReader("foo")
			r := NewReader(s)

			buf := make([]byte, 10)

			_, err := r.Read(buf)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(r.Count()).To(Equal(3))
		})
	})
})
