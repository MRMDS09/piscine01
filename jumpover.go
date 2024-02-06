package piscine

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}
	st := ""
	for i := 1; i < len(str); i++ {
		if (i+1)%3 == 0 {
			st += string(str[i])
		}
	}
	return st + "\n"
}
