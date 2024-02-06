package piscine

func ShoppingListSort(slice []string) []string {
	// sl := []string{}
	w := ""
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice)-1; j++ {
			if slice[i] > slice[j] {
				w = slice[j]
				slice[j] = slice[i]
				slice[i] = w
			}
		}
		//	sl[i] = slice[i]
	}
	return slice
}
