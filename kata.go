package kata

import (
	"fmt"
	"strconv"
	"strings"
)

// HighAndLow returns the highest and lowest number of given a string of space separated numbers
// highAndLow("1 2 3 4 5");  // return "5 1"
// highAndLow("1 2 -3 4 5"); // return "5 -3"
// highAndLow("1 9 3 4 -5"); // return "9 -5"
func HighAndLow(in string) string {
	var high, low int

	for i, s := range strings.Fields(in) {
		n, _ := strconv.Atoi(s)

		if i == 0 {
			high = n
			low = n
		}

		if n > high {
			high = n
		}

		if n < low {
			low = n
		}
	}

	return fmt.Sprintf("%d %d", high, low)
}

// ReverseWords accepts a string parameter, and reverses each word in the string.
// All spaces in the string should be retained.
func ReverseWords(str string) string {
	var reversed []string

	for _, s := range strings.Split(str, " ") {
		r := []rune(s)

		for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}

		reversed = append(reversed, string(r))
	}

	return strings.Join(reversed, " ")
}
