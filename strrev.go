package piscine

func StrRev(s string) string {
	chn := ""
	for i := len(s) - 1; i >= 0; i-- {
		chn = chn + string(s[i])
	}
	return chn
}
