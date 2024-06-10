package main

import (
	"bufio"
	"fmt"
)

func partTwo(s *bufio.Scanner) {
	sum := 0
	current := 0
	cards := make([]int, 206)
	for s.Scan() {
		// add the card value itself
		cards[current] += 1
		v := value(s.Text())
		for i := 1; i <= v; i++ {
			cards[current+i] += cards[current]
		}
		sum += cards[current]
		current += 1
	}
	fmt.Println("TOTAL PART 2: ", sum)
}
