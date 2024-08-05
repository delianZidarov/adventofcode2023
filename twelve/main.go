package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"regexp"
	//"strconv"
)

func main() {
	// TEST
	/*
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
			sum += comb(in, spec)
		}
	*/
	fmt.Println("TEST")
	//.#.#.###.
	comb("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6})
}

// the numbers provided in the input become the state machine
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

func comb(in string, spec []int) int {
	headList := head{current: 0, next: nil, prev: nil}
	c := &headList

	machine := automata(spec)
	// for each character in the input
	for _, ch := range in {
		// for each head in the headlist
			fmt.Println("Heads: ", headList)
		for c != nil {
			// if the input matches the current state advance the head
			if byte(ch) == machine[c.current] {
				c.current += 1
			}
			// if the input is indetermined add a head to the start of the list
			// with an advanced position. this covers both possibilities of it matching
			// and not matching
			if byte(ch) == '?' {
				headList = *addHead(&headList, c.current+1)
				fmt.Println("WHATS happening", headList, headList.next)
			}
			// after finishing this head set c to the next one in the list
			fmt.Println("this is c:", c, "next", *c.next, "prev", c.prev, c == &headList)
			c = c.next
		}
		// reset c to start of the headList after completing a single inpyt
		c = &headList
	}

	return 0
}

type head struct {
	current int
	next    *head
	prev    *head
}

func (h *head) Delete() {
	h.prev.next = h.next
}

func addHead(h *head, current int) *head {
	nHead := head{current: current, next: h}
	h.prev = &nHead
	fmt.Println("Created new head", nHead, nHead.next)
	return &nHead
}
