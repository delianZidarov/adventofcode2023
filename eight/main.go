package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	directions := "LRLRLRRLRRRLRRRLRRRLRLRRRLRRRLRRRLLRLRRRLRLRRRLLRRRLRRLRRRLRRLRLRRRLRRRLRLRRLRRRLRRLRRRLRRLRLRRLRRRLRLRRLRRRLLRRRLRLRRLLLRLLRLRRLLRRRLLRLLRRLRLRRRLLLRLRRLRLRRLRRRLRRLLRRLLRLRRRLRRRLRLLLLRLLRLRLRLRRRLRRLRRLRLRRRLLRRLRLLRRLRLRRLRLRLRRLRRLLRLRRLLRLLRRRLLLRRRLRRLRLRRRLRRLRRRLRRLLLRRRR"
	// test directions
	// directions := "RL"
	f, _ := os.Open("/home/d/Documents/day8.txt")
	// test file
	// f, _ := os.Open("/home/d/Documents/8test")
	scanner := bufio.NewScanner(f)

	r, _ := regexp.Compile("[A-Z]{3,3}")
	m := make(map[string][]string)
	start := ""

	for scanner.Scan() {
		node := r.FindAllString(scanner.Text(), 3)
		m[node[0]] = node[1:]
		if len(start) == 0 {
			start = node[0]
		}
	}

	f.Close()

	steps := 0
	nextNode := "AAA"
	for nextNode != "ZZZ" {
		if directions[steps%len(directions)] == 'L' {
			nextNode = m[nextNode][0]
		} else {
			nextNode = m[nextNode][1]
		}
		steps += 1
	}
	fmt.Println(steps)
}
