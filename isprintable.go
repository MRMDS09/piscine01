package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		// 0-33 code ascii dial non printable https://www.commentcamarche.net/informatique/technologies/1589-code-ascii/
		for j := 0; j < 30; j++ {
			// convert l str l int bach t9der tcampari j m3a code ascii dial str
			if int(str[i]) == j {
				return false
			}
		}
	}
	return true
}
