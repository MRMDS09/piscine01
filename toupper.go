package piscine

func ToUpper(s string) string {
	str := []rune(s)
	str2 := []rune(s)
	for i := 0; i < len(str); i++ {
		for j := 'a'; j <= 'z'; j++ {
			// recherche wach kyn chi lettre munuscil
			if j == str[i] {
				// ila l9ah ghadi j = 3adad ghadi dial l herf f code ascii ila n9enaleh 32 g
				// ye3tiina herf li bhalo walakin majuscile
				str2[i] = j - 32
			}
		}
	}
	return string(str2)
}
