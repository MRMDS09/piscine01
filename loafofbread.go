package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		s := ""
		i := 0
		for _, e := range str {
			if i%6 != 5 && e == ' ' {
				continue
			}
			if i%6 == 5 {
				s += " "
			} else {
				s += string(e)
			}
			i++
		}
		// remove trailing spaces
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != ' ' {
				s = s[:i+1]
				break
			}
		}
		return s + "\n"
	}
}
