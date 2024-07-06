package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, _ := os.Open("/home/d/Documents/day12")
	scanner := bufio.NewScanner(f)
	l, _ := regexp.Compile("[?#.]+")
	n, _ := regexp.Compile("[0-9]+")
	for scanner.Scan() {
		in := l.FindString(scanner.Text())
		c := n.FindAllString(scanner.Text(), -1)
		fmt.Println("Input: ", in, "Broken locations: ", c)

	}
}
