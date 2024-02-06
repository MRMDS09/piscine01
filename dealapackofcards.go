package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	cnt := 0
	fmt.Printf("Player 1:")
	for i := 0; i < 3; i++ {
		if cnt == 2 {
			fmt.Printf(" %d", deck[i])
			cnt = 0
			break
		} else {
			fmt.Printf(" %d,", deck[i])
			cnt++
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Player 2:")
	for i := 3; i < 6; i++ {
		if cnt == 2 {
			fmt.Printf(" %d", deck[i])
			cnt = 0
			break
		} else {
			fmt.Printf(" %d,", deck[i])
			cnt++
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Player 3:")
	for i := 6; i < 9; i++ {
		if cnt == 2 {
			fmt.Printf(" %d", deck[i])
			cnt = 0
			break
		} else {
			fmt.Printf(" %d,", deck[i])
			cnt++
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Player 4:")
	for i := 9; i < 12; i++ {
		if cnt == 2 {
			fmt.Printf(" %d", deck[i])
			cnt = 0
			break
		} else {
			fmt.Printf(" %d,", deck[i])
			cnt++
		}
	}
	fmt.Printf("\n")
}
