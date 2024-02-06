package main

import (
	"fmt"
	"os"
)

func main() {
	length := os.Args[1:]
	lengthOfArg := len(length)
	if lengthOfArg == 0 {
		fmt.Println("File name missing")
	}
	if lengthOfArg == 1 {
		fmt.Println("Almost there!!")
	} else {
		if lengthOfArg > 0 {
			fmt.Println("Too many arguments")
		}
	}
}
