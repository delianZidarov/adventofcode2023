package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	path := args[0]
	part := args[1]

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var games *node
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var value int
		if part == "1" {
			value = handScore(s[0])
		} else {
			value = jackHandScore(s[0])
		}
		bet, err := strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		games = insertNode(games, value, int(bet))
	}
	f.Close()
	i := 1
	sum := 0
	inorderTran(games, &i, &sum)
	fmt.Println("The total is: ", sum)
}
