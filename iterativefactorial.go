package piscine

func IterativeFactorial(nb int) int {
	s := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb <= 21 && nb > 1 {
		for i := 1; i <= nb; i++ {
			s = s * i
		}
	} else {
		return 0
	}
	return s
}
