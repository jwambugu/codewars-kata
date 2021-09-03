package kata_test

import (
	. "code-wars/kata"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	_ "github.com/onsi/gomega"
)

var _ = Describe("High And Low", func() {
	It("should work for sample test cases", func() {
		gomega.Expect(HighAndLow("1 2 3 4 5")).To(gomega.Equal("5 1"))
		gomega.Expect(HighAndLow("1 2 -3 4 5")).To(gomega.Equal("5 -3"))
		gomega.Expect(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")).To(gomega.Equal("42 -9"))
	})
})

var _ = Describe("Reverse Words", func() {
	It("should work for sample test cases", func() {
		gomega.Expect(ReverseWords("The quick brown fox jumps over the lazy dog.")).
			To(gomega.Equal("ehT kciuq nworb xof spmuj revo eht yzal .god"))

		gomega.Expect(ReverseWords("apple")).To(gomega.Equal("elppa"))
		gomega.Expect(ReverseWords("a b c d")).To(gomega.Equal("a b c d"))
		gomega.Expect(ReverseWords("double  spaced  words")).To(gomega.Equal("elbuod  decaps  sdrow"))
		gomega.Expect(ReverseWords("stressed desserts")).To(gomega.Equal("desserts stressed"))
		gomega.Expect(ReverseWords("hello hello")).To(gomega.Equal("olleh olleh"))
	})
})
