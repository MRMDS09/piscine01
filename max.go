package piscine

func Max(a []int) int {
	c := a[0]
	for i := 1; i < len(a); i++ {
		if c < a[i] {
			c = a[i]
		}
	}
	return c
}
