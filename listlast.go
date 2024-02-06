package piscine

func ListLast(l *List) interface{} {
	a := l.Head
	if a == nil {
		return nil
	}
	for a.Next != nil {
		a = a.Next
	}
	return a.Data
}
