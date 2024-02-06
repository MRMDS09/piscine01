package piscine

func ListPushFront(l *List, data interface{}) {
	// fmt.Println(data)
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		l.Tail = l.Head
		// fmt.Println("head = nil: \n", l.Head.Data)
	} else {
		// creation d'un objet a de type Nodel qui ne sont cree
		a := &NodeL{Data: data}
		// decaler la valeur qui est dans le head ca cera dans le next et laisser la valeur head vide
		a.Next = l.Head
		// fmt.Println("apres:a.next \n", a.Next.Data)
		// ajouter la valeur data dans l'obejt a nous cree et ajouetr dans le head de list l
		l.Head = a
		// fmt.Println("apres: \n", a.Data)
	}
}
