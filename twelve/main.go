package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// TEST
	f, _ := os.Open("/home/d/Documents/test")
	// TEST
	// f, _ := os.Open("/home/d/Documents/day12")
	scanner := bufio.NewScanner(f)
	l, _ := regexp.Compile("[?#.]+")
	n, _ := regexp.Compile("[0-9]+")
	for scanner.Scan() {
		in := l.FindString(scanner.Text())
		c := n.FindAllString(scanner.Text(), -1)
		spec := make([]int, len(c))
		// m := make(map[int]int)

		for i, s := range c {
			num, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println("Trouble parsing input numbers")
				os.Exit(1)
			}
			spec[i] = int(num)
		}
		fmt.Println("doing", in, spec)
	}
	comb("????", []int{1})
}

// the numbers of provided in the input become the state machine
func automata(spec []int) string {
	state := "."
	for _, s := range spec {
		for i := 0; i < s; i++ {
			state += "#"
		}
		state += "."
	}
	return state[:len(state)-1]
}

func reducer (a []int, state int) int {
	sum := 0
	for _, v := range a {
   if v == state {
     sum += 1
		}
	}
 return sum
}

func comb(in string, spec []int) int {
	headArray := make([]int, 1)

	machine := automata(spec)
	for _, input := range in {
		for j, head := range headArray {
			if head < len(machine)-1 {

				currentState := machine[head]
				nextState := machine[head+1]
				fmt.Println("input", string(input), "head", head, "current", string(currentState), "next", string(nextState), "HEADS", headArray)
				if input == rune(nextState) {
					headArray[j] += 1
				}
				// erase heads
				if currentState == '#' &&
					nextState == '.' &&
					input == '#' {
					headArray = append(headArray[:j], headArray[j+1:]...)
				}
				if currentState == '#' &&
					nextState == '#' &&
					input == '?' {
					headArray[j] += 1
				}
				if currentState == '#' &&
					nextState == '.' &&
					input == '?' {
					headArray[j] += 1
				}
				// this is branch in the state machine, a new head is created at the
				// current location and the original head is advanced one position
				if currentState == '.' &&
					input == '?' {
					fmt.Println(head)
					headArray = append(headArray, headArray[j])
					headArray[j] += 1

				}
			}
			// The end is gauranteed to be a "." so it should only accept "." or "#"
			if head == len(machine)-1 &&
				input == '#' {
				headArray = append(headArray[:j], headArray[j+1:]...)
			}
			fmt.Println("End of loops", headArray)

		}
	}
	fmt.Println("Possible answer?", reducer(headArray,len(machine) -1))

	return 0
}
