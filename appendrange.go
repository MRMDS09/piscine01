package piscine

func AppendRange(min, max int) []int {
	gr1 := []int{}
	if min < max {
		for i := min; i < max; i++ {
			gr1 = append(gr1, i)
		}
		return gr1
	}
	return nil
}
