package main

import (
	"fmt"
	"os"
)

func main() {
	p := os.Args[1]

	switch p {
	case "1":
		t := []int{52, 94, 75, 94}
		d := []int{426, 1374, 1279, 1216}
		score := 1
		for i := 0; i < len(t); i++ {
			a, b := minMult(t[i], d[i])
			score = score * (b-a+1)
		}
		fmt.Println("Part 1: ", score)
	case "2":
		t := 52947594
		d := 426137412791216
		a, b := minMult(t, d)
		fmt.Println("Part 2: ", b-a+1)
	}
}

func minMult(t, d int) (l, up int) {
	start := 0
	end := t
	for start-end != -1 {
		mid := (end-start)/2 + start
		if mid*(t-mid) > d {
			end = mid
		} else {
			start = mid
		}
	}
	if start*(t-start) > d {
		l = start
		up = t - start
	} else {
		l = end
		up = t - end
	}
	return l, up
}
