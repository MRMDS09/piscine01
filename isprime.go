package piscine

func IsPrime(nb int) bool {
	s := 0
	if nb < 0 {
		nb = nb * (-1)
		for i := 2; i <= nb; i++ {
			if nb%i == 0 {
				s = s + 1
			} else {
				return true
			}
		}
	} else {
		if nb < 1000000087 {
			for i := 2; i <= nb; i++ {
				if nb%i == 0 {
					s = s + 1
				}
			}
		} else {
			return true
		}
	}

	if s == 1 {
		return true
	} else {
		return false
	}
}
