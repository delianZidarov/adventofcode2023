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
	// f, _ := os.Open("/home/d/Documents/test")
	// Test
	f, _ := os.Open("/home/d/Documents/day9")
	scanner := bufio.NewScanner(f)

	r, _ := regexp.Compile("(-?)[0-9]+")

	sum := int64(0)
  prevSum := int64(0)

	for scanner.Scan() {
		s := intMap(r.FindAllString(scanner.Text(), -1))
		sum += nextValue(s)
		prevSum += prevValue(s)
	}
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", prevSum)
}

func intMap(s []string) []int64 {
	i := make([]int64, 0)
	for _, s := range s {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Println("error parsing")
		}
		i = append(i, v)
	}
	return i
}

func nextValue(v []int64) int64 {
	zeroes := 0
	nextArray := make([]int64, 0)
	for i := 1; i < len(v); i++ {
		if i == 0 {
			zeroes += 1
		}
		nextArray = append(nextArray, v[i]-v[i-1])
	}
	if zeroes == len(v) {
		return 0
	} else {
		return v[len(v)-1] + nextValue(nextArray)
	}
}

func prevValue(v []int64) int64 {
	zeroes := 0
	nextArray := make([]int64, 0)
	for i := 1; i < len(v); i++ {
		if i == 0 {
			zeroes += 1
		}
		nextArray = append(nextArray, v[i]-v[i-1])
	}
	if zeroes == len(v) {
		return 0
	} else {
		return v[0] - prevValue(nextArray)
	}
}
