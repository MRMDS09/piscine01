package piscine

func Abort(a, b, c, d, e int) int {
	az := make([]int, 5)
	// az2 := make([]int, 5)
	az[0] = a
	az[1] = b
	az[2] = c
	az[3] = d
	az[4] = e
	cnt := 0
	w := 0
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if az[i] > az[j] {
				w = az[j]
				az[j] = az[i]
				az[i] = w
				cnt++
			}
		}
	}
	return az[2]
}
