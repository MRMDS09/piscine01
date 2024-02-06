package main

import (
	"os"

	"github.com/01-edu/z01"
)

// import ""

func main() {
	args := os.Args[0]
	for _, v := range args {
		if v == '/' || v == '.' {
			continue
		} else {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
