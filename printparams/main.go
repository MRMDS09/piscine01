package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	// os.Args[1:] bach tbedi mn hadi choummi is the best cat madkholch hadi /tmp/go-build2786375066/b001/exe/printparams hit hadi hya 0 fin kyn db w ila dert 1 bohdo
	// bla slice ghadi y3tini bli hadik raha position dialha f code ascii

	for _, v := range args {
		for _, s := range v {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')

	}
}
