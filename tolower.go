package piscine

func ToLower(s string) string {
	str := []rune(s)
	str2 := []rune(s)
	for i := 0; i < len(str); i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if str[i] == j {
				str2[i] = j + 32
			}
		}
	}
	return string(str2)
}
