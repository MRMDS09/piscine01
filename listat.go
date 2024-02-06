package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	a := l
	c := 0
	for a != nil {
		if c == pos {
			return a
		}
		a = a.Next
		c++
	}
	return nil
}
