package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	a := 1
	var z []int
	as := 0
	for i := len(nbr) - 1; i >= 0; i-- {
		if nbr[i] != '0' {
			z = append(z, a)
		}
		a *= 2
	}
	for _, v := range z {
		as += v
	}
	s := ""
	for as > 0 {
		di := as % 10
		s = string(di+'0') + s
		as /= 10
	}
	return s
}
