package main

import (
	"fmt"
	"os"
)

func main() {
	part, path, pool, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f := fileReader{P: path}
	f.Open()
	for {
		ln, err := f.NextLn()
		if err != nil {
			break
		}
		Score(ln, &pool)
	}
	fmt.Println(part, path, pool)
}
