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
	comb(".??..?##?", []int{1, 3})
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
	return state
}

func comb(in string, spec []int) int {
	stateMachine := automata(spec)

	// we start with one default state, the 0
	// since this a NFA we need to make sure to keep track
	// of branching. This can be done recursively making 2
	// branches or by keeping track of states. As we proggress
	// through the machine the state array gets larger. At each
	// input we assume we are at every already visited state
	visitedStates := make([]int,0)
	visitedStates = append(visitedStates, 1)
	// go through every input for the state Machine
	for _, c := range in {
		fmt.Println("Checking char", string(c))
		for i := range visitedStates {
			if transition(i, stateMachine, byte(c)) {
				lastVisited := len(visitedStates) - 1
				if lastVisited <= i {
					visitedStates = append(visitedStates, 1)
				} else {
					visitedStates[i] += 1
				}
			}
		}
		fmt.Println("State after checking the char", visitedStates, "first",visitedStates[0])
	}
	fmt.Println("State Machine:", stateMachine, "Intput", in)
	fmt.Println("How is it looking", visitedStates, "last position", len(stateMachine)-1)
	return 0
}

func transition(curLoc int, state string, in byte) bool {
	lastPosition := len(state) - 1
	if curLoc >= lastPosition {
		return false
	}
	if in == state[curLoc+1] || in == '?' {
		return true
	}
	return false
}
