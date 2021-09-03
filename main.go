package main

import (
	"code-wars/kata"
	"fmt"
)

func highAndLow() {
	fmt.Println(kata.HighAndLow("1 2 3 4 5"))                    // "5 1"
	fmt.Println(kata.HighAndLow("1 2 -3 4 5"))                   // "5 -3"
	fmt.Println(kata.HighAndLow("1 9 3 4 -5"))                   // "9 -5"
	fmt.Println(kata.HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")) // "42 -9"
}

func main() {
	//fmt.Println(kata.ReverseWords("The quick brown fox jumps over the lazy dog."))
	//fmt.Println(kata.ReverseWords("double  spaced  words"))
	//fmt.Println(kata.ToWeirdCase("double  spaced  words"))
	//fmt.Println(kata.ToWeirdCase("abc def"))
	fmt.Println(kata.ToWeirdCase("This is a test Looks like you passed"))
	//fmt.Println(kata.Reverser("double  spaced  words"))
	//fmt.Println(strings.Split("double  spaced  words", " "))
}
