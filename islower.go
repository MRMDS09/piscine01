package piscine

func IsLower(s string) bool {
	str := []rune(s)
	a := 0
	for i := 0; i < len(s); i++ {
		for j := 'a'; j <= 'z'; j++ {
			if str[i] == j {
				a++
			}
		}
	}
	if a != len(s) {
		return false
	} else {
		return true
	}
}
