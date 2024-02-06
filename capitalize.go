package piscine

func isAlphaNumeric(a rune) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}

func Capitalize(s string) string {
	ar := []rune(s)

	shouldCapitalize := true

	for i := 0; i < len(s); i++ {
		if isAlphaNumeric(ar[i]) && shouldCapitalize {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] -= 'a' - 'A'
			}
			shouldCapitalize = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] += 'a' - 'A'
		} else if !isAlphaNumeric(ar[i]) {
			shouldCapitalize = true
		}
	}

	return string(ar)
}
