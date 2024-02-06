package piscine

func SortIntegerTable(table []int) {
	se := 0
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				se = table[i]
				table[i] = table[j]
				table[j] = se
			}
		}
	}
}
