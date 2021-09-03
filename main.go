package main

import (
	"code-wars/kataa"
	"fmt"
)

func highAndLow() {
	fmt.Println(kataa.HighAndLow("1 2 3 4 5"))                    // "5 1"
	fmt.Println(kataa.HighAndLow("1 2 -3 4 5"))                   // "5 -3"
	fmt.Println(kataa.HighAndLow("1 9 3 4 -5"))                   // "9 -5"
	fmt.Println(kataa.HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")) // "42 -9"
}

func main() {
	//fmt.Println(kataa.ReverseWords("The quick brown fox jumps over the lazy dog."))
	//fmt.Println(kataa.ReverseWords("double  spaced  words"))
	//fmt.Println(kataa.ToWeirdCase("double  spaced  words"))
	//fmt.Println(kataa.ToWeirdCase("abc def"))
	//fmt.Println(kataa.ToWeirdCase("This is a test Looks like you passed"))
	//fmt.Println(kataa.Reverser("double  spaced  words"))
	//fmt.Println(kataa.NbYear(1000, 2, 50, 1200)) // 3
	//fmt.Println(kataa.NbYear(1500, 5, 100, 5000))           // 15
	//fmt.Println(kataa.NbYear(1500000, 2.5, 10000, 2000000)) // 10
}
