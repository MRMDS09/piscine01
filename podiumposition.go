package piscine

func PodiumPosition(podium [][]string) [][]string {
	mtx := make([][]string, len(podium))
	j := 0
	for i := 3; i >= 0; i-- {
		mtx[j] = podium[i]
		j++
	}
	return mtx
}
