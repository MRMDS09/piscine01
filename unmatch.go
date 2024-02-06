package piscine

func Unmatch(a []int) int {
	as := len(a) / 2
	// s := 0
	az := make([]int, as)
	cnt := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				az[cnt] = a[cnt]
				cnt++
				break
			} else {
				continue
			}
		}
	}

	return cnt
}
