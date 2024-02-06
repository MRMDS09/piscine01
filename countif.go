package piscine

func CountIf(f func(string) bool, tab []string) int {
	ea := 0
	for i := 0; i < len(tab); i++ {
		ez := f(tab[i])
		if ez == true {
			ea++
		}

	}
	return ea
}
