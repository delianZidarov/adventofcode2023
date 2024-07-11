package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"slices"
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
		//m := make(map[int]int)

		for i, s := range c {
			num, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println("Trouble parsing input numbers")
				os.Exit(1)
			}
			spec[i] = int(num)
		}
   
		comb(in, spec)

	}
}


type car struct {
  front int
	back  int
	track *string
	lastValStop *car
  valStops *[]car
}





func (c *car) run(end int) {
   loc := make([]car, 0)
	for c.front <= end {
    if c.validate() {
     loc = append(loc, *c)
		}
    c.front += 1
		c.back += 1
	}


}

func (c *car) validate() bool {
 r, _ := regexp.Compile("^[.?][?#][.?] $")
 return r.Match([]byte( (*c.track)[c.back : c.front+1]))
}

func newCars (specs []int, t *string) []car {
	cars := make([]car, 0)
	lastFront := 0
	for _, c := range specs {
    //we adjust the length of the car by +2 for the spacers and -1 because arrays 
		//start at zero 
		newFront := lastFront + c + 1
		cars = append(cars, car{front: newFront, back: lastFront, track: t })
    lastFront = newFront
	}
  return cars
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

