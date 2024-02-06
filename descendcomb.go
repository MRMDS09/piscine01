package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for z := '9'; z >= '0'; z-- {
				for a := '9'; a >= '0'; a-- {
					if (i > z) || (j > a && i == z) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(a)
						if !(i == '0' && j == '1' && z == '0' && a == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
