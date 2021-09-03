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

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values", func() {
		gomega.Expect(ToWeirdCase("abc def")).To(gomega.Equal("AbC DeF"))
		gomega.Expect(ToWeirdCase("ABC")).To(gomega.Equal("AbC"))
		gomega.Expect(ToWeirdCase("This is a test Looks like you passed")).To(gomega.Equal("ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"))
	})
})

var _ = Describe("NbYear", func() {
	It("fixed tests", func() {
		gomega.Expect(NbYear(1500, 5, 100, 5000)).To(gomega.Equal(15))
		gomega.Expect(NbYear(1500000, 2.5, 10000, 2000000)).To(gomega.Equal(10))
		gomega.Expect(NbYear(1500000, 0.25, 1000, 2000000)).To(gomega.Equal(94))
		gomega.Expect(NbYear(1500000, 0.25, -1000, 2000000)).To(gomega.Equal(151))
	})
})
