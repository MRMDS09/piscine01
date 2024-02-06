package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune('n')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var result []rune
	for n > 0 {
		digit := n % 10
		result = append(result, rune(digit+'0'))
		n /= 10
	}
	var a rune
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				a = rune(result[i])
				result[i] = result[j]
				result[j] = a
			}
		}
	}
	for i := 0; i < len((result)); i++ {
		z01.PrintRune(result[i])
	}
}
