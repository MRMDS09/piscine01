package piscine

func Atoi(s string) int {
	a := 1
	r := 0
	for i, v := range s {
		if (i == 0) && ((v == ('-')) || (v == ('+'))) {
			if v == ('-') {
				a = -1
			}
		} else {
			if v >= '0' && v <= '9' {
				r = (r * 10) + int(v-48)
			} else {
				return 0
			}
		}
	}
	l := r * a
	return l
}
