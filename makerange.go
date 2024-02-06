package piscine

func MakeRange(min, max int) []int {
	a := max - min
	if a < 0 {
		return nil
	}
	gr1 := make([]int, a)
	if min < max {
		for i := 0; i < a; i++ {
			gr1[i] = min + i
		}
		return gr1
	}
	return nil
}
