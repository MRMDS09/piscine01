package piscine

func RockAndRoll(n int) string {
	rock := 0
	roll := 0
	if n < 0 {
		return "error: number is negative\n"
	}
	if n%2 == 0 {
		rock++
	}
	if n%3 == 0 {
		roll++
	}
	if rock != 0 && roll != 0 {
		return "rock and roll\n"
	} else {
		if roll != 0 {
			return "roll\n"
		} else {
			if rock != 0 {
				return "rock\n"
			}
		}
	}
	return "error: non divisible\n"
}
