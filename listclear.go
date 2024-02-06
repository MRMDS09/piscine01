package piscine

func ListClear(l *List) {
	a := l.Head
	for a != nil {
		s := a.Next
		a.Next = nil
		a = s
	}
	l.Head = nil
}
