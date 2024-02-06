package piscine

func UltimateDivMod(a *int, b *int) {
	aa := *a / *b
	bb := *a % *b
	*a = aa
	*b = bb
}
