package piscine

func ConcatParams(args []string) string {
	gr := ""
	for i := 0; i < len(args); i++ {
		gr += args[i]
		if i < len(args)-1 {
			gr += "\n"
		}
	}
	return gr
}
