package piscine

func Compact(ptr *[]string) int {
	a := 0
	m := make([]string, 0)
	for _, v := range *ptr {
		if v != "" { // ila kanet khawya bhal a[1] a[3] a[5] ydkhol yzid whd lblayes li 3amrin f l a bhal a[0] a[2] a[4]
			m = append(m, v) // n3emro lslice m bli a li 3amrin bhal db m[0]=a[0] w m[1]=a[2] w m[2]=a[4]
			a++              // ghadi yb9a ytzad b lwhd whd l conteur li ghadi ykon chhale mn whd 3emer f lm
		}
	}
	*ptr = m // n3ti<w dik l m l *ptr
	return a
}
