package piscine

func TrimAtoi(s string) int {
	d := 0
	a := 0
	c := 0
	for _, v := range s {
		if v == '-' && a == 0 {
			c = -1
		} else {
			if v >= '0' && v <= '9' {
				d = d*10 + int(v-'0')
				a++
			}
		}
	}
	if c == -1 {
		d *= -1
	}
	return d
}
