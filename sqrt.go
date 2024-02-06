package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb > 0 {
		r := 1
		for i := 1; i < nb; i++ {
			r = i * i
			if r == nb {
				return i
			}
		}
		return 0
	} else {
		return 0
	}
}
