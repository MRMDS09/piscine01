package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < 47; j++ {
			if int(str[i]) == j {
				return false
			}
		}
		for j := 58; j < 64; j++ {
			if int(str[i]) == j {
				return false
			}
		}
		for j := 91; j < 96; j++ {
			if int(str[i]) == j {
				return false
			}
		}
		for j := 123; j < 255; j++ {
			if int(str[i]) == j {
				return false
			}
		}
	}
	return true
}
