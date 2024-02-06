package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	if n <= 0 {
		return 0
	} else {
		if n > len(str) {
			return 0
		} else {
			if n-1 <= len(str) {
				return str[n-1]
			}
		}
	}

	return 0
}
