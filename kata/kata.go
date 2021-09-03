package kata

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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

// ToWeirdCase accepts a string, and returns the same string with all even indexed characters in each word upper cased,
// and all odd indexed characters in each word lower cased.
// The indexing just explained is zero based, so the zero-ith index is even,
// therefore that character should be upper cased.
func ToWeirdCase(str string) string {
	result, index := "", 0

	for _, s := range str {
		if s == ' ' {
			index = 0
			result += " "
			continue
		}

		if index%2 == 0 {
			result += string(unicode.ToUpper(s))
		} else {
			result += string(unicode.ToLower(s))
		}

		index++
	}

	return result

	//var result []string
	//
	//for _, s := range strings.Split(str, " ") {
	//	r := []rune(s)
	//
	//	for i := 0; i < len(s); i++ {
	//		if i%2 == 0 {
	//			r[i] = unicode.ToUpper(r[i])
	//		} else {
	//			r[i] = unicode.ToLower(r[i])
	//		}
	//	}
	//
	//	result = append(result, string(r))
	//}
	//
	//return strings.Join(result, " ")
}

// NbYear returns number of years taken to get x habitats
// Challenge::
// In a small town the population is p0 = 1000 at the beginning of a year.
// The population regularly increases by 2 percent per year and moreover 50 new inhabitants per year come to live
// in the town.
// How many years does the town need to see its population greater or equal to p = 1200 inhabitants?
// At the end of the first year there will be:
// 1000 + 1000 * 0.02 + 50 => 1070 inhabitants
//
// At the end of the 2nd year there will be:
// 1070 + 1070 * 0.02 + 50 => 1141 inhabitants (** number of inhabitants is an integer **)
//
// At the end of the 3rd year there will be:
// 1141 + 1141 * 0.02 + 50 => 1213
//
// It will need 3 entire years.
//
// More generally given parameters:
// p0, percent, aug (inhabitants coming or leaving each year), p (population to surpass)
//
// the function nb_year should return n number of entire years needed to get a population greater or equal to p.
//
// aug is an integer, percent a positive or null floating number, p0 and p are positive integers (> 0)

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0

	if p0 >= p {
		return 0
	}

	newPopulation := p0 + int(float64(p0)*percent/100) + aug
	years++

	return years + NbYear(newPopulation, percent, aug, p)
}
