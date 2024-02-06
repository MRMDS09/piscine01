package piscine

func Join(strs []string, sep string) string {
	ls := ""
	for i := 0; i < len(strs)-1; i++ {
		ls += strs[i] + sep
	}
	ls += strs[len(strs)-1]
	return ls
}
