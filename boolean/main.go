package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	Msg1 := "I have an "
	Msg2 := " number of arguments"
	for _, r := range Msg1 {
		z01.PrintRune(r)
	}
	for _, r := range s {
		z01.PrintRune(r)
	}
	for _, r := range Msg2 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) == 0 {
		return true
	} else {
		return false
	}
}

func even(nbr int) int {
	return nbr % 2
}

func main() {
	length := os.Args[1:]
	lengthOfArg := len(length)
	if isEven(lengthOfArg) {
		printStr("even")
	} else {
		printStr("odd")
	}
}
