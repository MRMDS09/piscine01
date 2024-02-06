package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	gr := ""
	for i := 0; i < len(a); i++ {
		gr += a[i]
		if i < len(a)-1 {
			gr += "\n"
		}
	}
	for i := 0; i < len(gr); i++ {
		z01.PrintRune(rune(gr[i]))
	}
	z01.PrintRune('\n')
}
