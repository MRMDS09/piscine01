package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	a := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = a
	} else {
		s := l.Head
		for s.Next != nil {
			s = s.Next
		}

		//	fmt.Println("avant: \n", l.Head.Next.Data)

		s.Next = a
	}
}
