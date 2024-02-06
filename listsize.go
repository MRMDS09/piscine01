package piscine

func ListSize(l *List) int {
	a := 0

	s := l.Head
	for s != nil {
		a++
		s = s.Next
	}
	return a
}
