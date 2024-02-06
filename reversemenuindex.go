package piscine

func ReverseMenuIndex(menu []string) []string {
	a := make([]string, len(menu))
	j := 0
	for i := len(menu) - 1; i >= 0; i-- {
		a[j] = menu[i]
		j++
	}
	return a
}
