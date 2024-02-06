package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var ar string

	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				ar = args[i]
				args[i] = args[j]
				args[j] = ar
			}
		}
	}
	for _, v := range args {
		for _, s := range v {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
