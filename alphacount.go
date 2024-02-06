package piscine

func AlphaCount(s string) int {
	a := 0
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		for j := 'a'; j <= 'z'; j++ {
			if str[i] == j {
				a++
			}
		}
		for j := 'A'; j <= 'Z'; j++ {
			if str[i] == j {
				a++
			}
		}
	}
	return a
}
