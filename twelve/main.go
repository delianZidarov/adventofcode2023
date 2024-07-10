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
		in := l.FindAllString(scanner.Text(), -1)
		c := n.FindAllString(scanner.Text(), -1)
		spec := make([]int, len(c))
		m := make(map[int]int)

		for i, s := range c {
			num, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println("Trouble parsing input numbers")
				os.Exit(1)
			}
			spec[i] = int(num)
		}

		comb, ok := findComb(in[0], spec, &m)
		fmt.Println(in[0], "----", comb, "-----", ok)

	}
}

func factorial(a int, mem *map[int]int) int {
	// assert that a is less than 21, higher values overflow
	if a > 20 {
		fmt.Println("Passing too big an integer to factorial")
		os.Exit(1)
	}
	if a < 0 {
		fmt.Println("Passing a negative value to factorial")
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

// check to see if the input chunk is all question marks
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

func findComb(in string, spec []int, m *map[int]int) (int, bool) {
	fmt.Println("Working on: ", in, spec)

	bigestIdx := 0
	biggestVal := 0

	if len(in) == 0 {
		if len(spec) == 0 {
			return 0, true
		} else {
			return 0, false
		}
	}

	if len(spec) == 0 {
		return 0, true
	}
	// if the string is all ?
	if allMaybe(in) {
		spaces := len(in) - sequenceLengthModifier(spec) - (len(spec) - 1)
		pieces := len(spec)
		if spaces == pieces {
			return 0, true
		} else if spaces < pieces {
			return 0, false
		} else {
			n := factorial(spaces, m) / (factorial((spaces-pieces), m) * factorial(pieces, m))
			if n > 0 {
				return n, true
			} else {
				return 0, false
			}
		}

	}

	// find the location of the biggest chunk
	for i, v := range spec {
		if v > biggestVal {
			bigestIdx = i
			biggestVal = v
		}
	}

	if len(in) < biggestVal {
		return 0, false
	}

	s := slider{0, biggestVal - 1}
	if len(spec) == 1 {
		solutions := 0
		for i := 0; i < len(in); i++ {
			var valid bool
			if s.lBound == 0 && s.rBound != len(in)-1 {
				valid = validatePlacement("." + in[:s.rBound+1])
			}
			if s.lBound != 0 && s.rBound == len(in)-1 {
				valid = validatePlacement(in[s.rBound:] + ".")
			}
			if s.lBound == 0 && s.rBound == len(in)-1 {
				valid = validatePlacement("." + in + ".")
			}
			if valid {
				solutions += 1
			}
			s.lBound += 1
			s.rBound += 1
		}
		return solutions, true
	} else {
		solutions := 0
		for i := 0; i < len(in); i++ {
			var valid bool
			if s.lBound == 0 && s.rBound != len(in)-1 {
				valid = validatePlacement("." + in[:s.rBound])
			}
			if s.lBound != 0 && s.rBound == len(in)-1 {
				valid = validatePlacement(in[s.rBound:] + ".")
			}
			if s.lBound == 0 && s.rBound == len(in)-1 {
				valid = validatePlacement("." + in + ".")
			}
			if valid {
				fmt.Println("slider: ", s, in[0:0],"Left", in[0:s.lBound], spec[0:bigestIdx], "Right", in[s.rBound:], spec[bigestIdx+1:])
				
				if s.lBound == 0 {
				left, leftOk := findComb("", spec[0:bigestIdx], m)
				right, rightOk := findComb(in[s.rBound:], spec[bigestIdx+1:], m)
					if leftOk && rightOk {
					return left + right, true
				}

				}	
				left, leftOk := findComb(in[0:s.lBound], spec[0:bigestIdx], m)
				right, rightOk := findComb(in[s.rBound + 1:], spec[bigestIdx+1:], m)
				
				if leftOk && rightOk {
					return left + right, true
				}
			}
			s.lBound += 1
			s.rBound += 1
		}

		return solutions, true

	}
}

func validatePlacement(subS string) bool {
	v := true && (subS[0] == '.' || subS[0] == '?') &&
		(subS[len(subS)-1] == '.' || subS[len(subS)-1] == '?')
	if v == true {
		for i := 1; i < len(subS)-1; i++ {
			if subS[i] == '.' {
				v = false
			}
		}
	}
	return v
}

type slider struct {
	lBound int
	rBound int
}
