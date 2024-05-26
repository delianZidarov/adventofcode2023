package main

import (
	"fmt"
	"os"
)

func main() {
	part, path, pool, err := parseInput(os.Args[1:])
	sum := 0
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
		if part == 1 {
			n, err := Score(ln, &pool)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sum += n
		}
		if part == 2 {
			n, err := Power(ln)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sum += n
		}
	}
	fmt.Println(sum)
}
