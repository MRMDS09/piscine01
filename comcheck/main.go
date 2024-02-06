package main

import (
	"fmt"
	"os"
)

func main() {
	osargs := os.Args[0:]
	a := 0
	for i := 0; i < len(osargs); i++ {
		for j := 0; j < len(osargs[i]); j++ {
			if string(osargs[i]) == "galaxy" || string(osargs[i]) == "01" || string(osargs[i]) == "galaxy 01" {
				a++
				break
			}
		}
	}
	if a > 0 {
		fmt.Printf("Alert!!!")
		fmt.Printf("\n")
	}
}
