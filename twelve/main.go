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
		sum := 0
		for i, s := range c {
			num, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println("Trouble parsing input numbers")
				os.Exit(1)
			}
			spec[i] = int(num)
		}
		fmt.Println("doing", in, spec)
		sum += comb(in, spec)
	}
	fmt.Println("TEST")
	//.#.#.###.
	comb("?#?#?#?#?#?#?#?", []int{1,3,1,6})
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

func reducer(a []int, state int) int {
	sum := 0
	for _, v := range a {
		if v == state {
			sum += 1
		}
	}
	return sum
}

func delHead(a []int, idx int) []int {
  n := make([]int, 0)
	for i, v := range a {
   if i == idx {
     continue
		}
		n = append(n, v)
	}
	return n
}

func comb(in string, spec []int) int {
	headArray := make([]int, 1)

	machine := automata(spec)
	for _, input := range in {
		 nHeads := make([]int,0)
		 writeNHead := false
		fmt.Println(headArray, "INPUT", string(input), len(headArray), "first location", headArray[0])
		for j:=0; j < len(headArray); j++{
			head := headArray[j]
			if head < len(machine)-1 && j < len(headArray) {
				currentState := machine[head]
				nextState := machine[head+1]
				fmt.Println("Cur", string(currentState),"Next", string(nextState), "In", string(input), "J", j)
				if input == rune(nextState) {
					fmt.Println("Something funky", headArray)
					headArray[j] += 1
					fmt.Println("Something funky 2", headArray)
				}

				// erase heads
				if currentState == '#' &&
					nextState == '.' &&
					input == '#' {
					headArray = delHead(headArray, j)
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
					nHeads =append(nHeads, headArray[j])
					headArray[j] += 1
					writeNHead = true

				}
			}
			// The end is gauranteed to be a "." so it should only accept "." or "#"
			if head == len(machine)-1 &&
				input == '#' {
				headArray = delHead(headArray, j)
			}
		}
    if writeNHead {
      headArray= append(headArray, nHeads ...)
		}

	}
	fmt.Println("Possible answer?", headArray, reducer(headArray, len(machine)-1))

	return reducer(headArray, len(machine)-1)
}
type head struct {
  current int
	next   *head
	prev   *head
}

func (h *head) Delete () {
  h.prev.next = h.next
}

func addHead (h *head, current int) *head {
 nHead := head{current : current, next: h}
 h.prev = &nHead
 return &nHead
}
