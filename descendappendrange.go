package piscine

func DescendAppendRange(max, min int) []int {
	sl := []int{}
	if min > max {
		return sl
	}
	for i := max; i > min; i-- {
		sl = append(sl, i)
	}
	return sl
}
