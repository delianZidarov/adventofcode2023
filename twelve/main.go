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
	t, ok = combinations([]string{"????"}, []int{1, 1}, &m)
	fmt.Println(t, ok, "Expected 3 - 1")
	t, ok = combinations([]string{"????"}, []int{2, 1}, &m)
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

func comb(in string, specs []int, m *map[int]int) (int, bool) {
	if len(specs) == 0 && len(in) == 0 {
		return 0
	}
	if len(specs) == 1 {
	}
}

func findComb(in string, spec []int, m *map[int]int) (int, []string, bool) {
	bigestIdx:= 0
	biggestVal :=0

   // if the string is all ?
   if allMaybe(in) {
     spaces := len(in) - sequenceLengthModifier(spec) - (len(spec) - 1 )
		 pieces := len(spec)
		 if spaces == pieces {
     return 0, []string{""}, true 
		 } else {
      n := factorial(spaces, m) / (factorial((spaces - pieces), m) * factorial(pieces, m))
			if n > 0 {
			return n, []string{""},true
			} else {
       return 0, []string{""},false
			}
		}
 
  	}



	//find the location of the biggest chunk
	for i, v := range spec {
    if v > biggestVal {
      bigestIdx = i
			biggestVal = v
		}
	}
 
	// construct the slider
  var brokenPieceLocation int
	for i,c := range in {
   if c == '.' {
     brokenPieceLocation = i
		}
	}

}

func validatePlacement(subS string) bool {
	v := true && (subS[0] == '.' || subS[0] == '?') &&
		(subS[len(subS)-1] == '.' || subS[len(subS)-1] == '?')
	return v
}

type slider struct {
	rBound int
	lBound int
}
