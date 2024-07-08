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
	l, _ := regexp.Compile("[?#]+")
	n, _ := regexp.Compile("[0-9]+")
	for scanner.Scan() {
		in := l.FindAllString(scanner.Text(), -1)
		c := n.FindAllString(scanner.Text(), -1)
		spec := make([]int, len(c))

		for i, s := range c {
			num, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println("Trouble parsing input numbers")
				os.Exit(1)
			}
			spec[i] = int(num)
		}

		fmt.Println("Input: ", in, "Broken locations: ", spec)

	}
	m := make(map[int]int)
	fmt.Println(factorial(5, &m))
	t, ok := combinations([]string{"??"}, []int{1}, &m)
	fmt.Println(t, ok, "Expected  2 - 1")
	t, ok = combinations([]string{"???"}, []int{1, 1}, &m)
	fmt.Println(t, ok, "Expected 1 -1")
	t, ok = combinations([]string{"????"}, []int{1,1}, &m)
	fmt.Println(t, ok, "Expected 3 - 1")
	t, ok = combinations([]string{"????"}, []int{2,1}, &m)
	fmt.Println(t, ok, "Expected 1 - 1")

}


func factorial(a int, mem *map[int]int) int {
	// assert that a is less than 21, higher values overflow
	if a > 20 {
		fmt.Println("Passing too big an integer to factorial")
		os.Exit(1)
	}
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}

	n, ok := (*mem)[a]
	if ok {
		return n
	} else {
		b := a * factorial(a-1, mem)
		(*mem)[a] = b
		return b
	}
}

func combinations(chunk []string, spec []int, mem *map[int]int) (c int, ok bool) {
	if len(chunk) == 0 && len(spec) != 0 {
		c = 0
		ok = false
		return
	}

	if allMaybe(chunk[0]) {
		// if we are trying to fit more than one broken sequence in the space
		// we need to subtract one space for each extra broken sequence
		spacerSpaces := len(spec) - 1
		sequenceLengths := sequenceLengthModifier(spec)
		spaces := len(chunk[0]) - spacerSpaces - sequenceLengths
		// find all possible combinations
		if spaces == len(spec) {
			c = 0
			ok = true
			return
		} else {
			comb := factorial(spaces, mem) / (factorial(spaces-len(spec), mem) * factorial(len(spec), mem))
			c = comb - 1
			ok = true
			return
		}
	}

	return 0, false
}

func allMaybe(s string) bool {
	b := true
	for _, c := range s {
		b = b && (c == '?')
	}
	return b
}

// Calculate how much extra space the sequence of broken springs takes
func sequenceLengthModifier(spec []int) int {
	sum := 0
	for _, n := range spec {
		sum += n - 1
	}
	return sum
}
