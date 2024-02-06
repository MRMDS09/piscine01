package piscine

func Split(s, sep string) []string {
	var l []string
	sl := len(sep)
	for {
		a := -1
		n := len(s)

		for i := 0; i < n-sl; i++ {
			if s[i:i+sl] == sep {
				a = i
				break
			}
		}
		if a == -1 {
			l = append(l, s)
			break
		}

		l = append(l, s[:a])
		s = s[a+sl:]
	}
	return l
}
