package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	a := 0
	for i := 0; i < len(s); i++ {
		for j := '0'; j <= '9'; j++ {
			if str[i] == j {
				a++
			}
		}
	}
	if a == len(s) {
		return true
	} else {
		return false
	}
}
