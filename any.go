package piscine

func Any(f func(string) bool, a []string) bool {
	var ea bool
	for i := 0; i < len(a); i++ {
		ea = f(a[i])
		if ea == true {
			return true
		}

	}
	return ea
}
