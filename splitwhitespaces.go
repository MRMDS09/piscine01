package piscine

func SplitWhiteSpaces(s string) []string {
	var gr []string
	str := ""

	for _, i := range s {
		if i != ' ' && i != '\t' && i != '\n' {
			str += string(i)
		} else {
			if str != "" {
				gr = append(gr, str)
				str = ""
			}
		}
	}
	gr = append(gr, str)
	return gr
}
