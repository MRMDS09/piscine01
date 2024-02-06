package piscine

func BasicAtoi(s string) int {
	k := 0

	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		} else {
			chn := int(v - '0')
			k = k * 10
			k = k + chn
		}
	}
	return k
}
