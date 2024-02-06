package main

import "github.com/01-edu/z01"

const P = "x = 42, y = 21\n"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	pnt := point{}

	setPoint(&pnt)
	for _, v := range P {
		z01.PrintRune(v)
	}
}
