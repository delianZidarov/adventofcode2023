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
	sr, _ := regexp.Compile("[A-Z][A-Z]A")
	er, _ := regexp.Compile("[A-Z][A-Z]Z")

	m := make(map[string][]string)
	start := make([]string, 0)

	for scanner.Scan() {
		node := r.FindAllString(scanner.Text(), 3)
		m[node[0]] = node[1:]
		if sr.MatchString(node[0]) {
			start = append(start, node[0])
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
	fmt.Println("Part 1: ", steps)

	steps = 0
	mult := make([]int, len(start))
	writeCount := 0
	for true {
		// Check if we have found the end for all nodes and exit
		// loop if yes
		for i, node := range start {
			if er.MatchString(node) && mult[i] == 0 {
				mult[i] = steps
				writeCount += 1
			}
		}

		if writeCount == len(start) {
			break
		}

		// update the nodes that will be checked
		for i, node := range start {
			if directions[steps%len(directions)] == 'L' {
				start[i] = m[node][0]
			} else {
				start[i] = m[node][1]
			}
		}
		// increase steps
		steps += 1
	}
	fmt.Println("Part 2: ", mult)
}
