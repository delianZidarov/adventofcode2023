package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	p := 0
	if len(args) == 1 || args[1] == "1" {
		p = 1
	} else if args[1] == "2" {
		p = 2
	}

	f, err := os.Open(args[0])
	defer f.Close()

	buf := make([]byte, 5745)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Read(buf)
	b := blocks(buf)
	fmt.Println("******************")
	fmt.Println(parseSeeds(b[0]))
	fmt.Println("******************")
  fmt.Println(string(b[1]))
	switch p {
	case 1:
		fmt.Println("Part 1")
	case 2:
		fmt.Println("Part 2")
	}
}

func blocks(buf []byte) (chunks [][]byte) {
	mid := bytes.Split(buf, []byte("\n\n"))
	for i := 0; i < len(mid)-1; i++ {
		if len(mid[i]) > 0 {
			chunks = append(chunks, mid[i])
		}
	}
	return chunks
}

func parseSeeds(block []byte) (seeds []int, err error) {
	in := make([]int, 0)
	for i, char := range block {
		if char == ' ' {
			in = append(in, i)
		}
	}

	for i := 0; i < len(in); i++ {
		if i == len(in)-1 {
			n, err := strconv.ParseInt(string(block[in[i]+1:]), 10, 64)
			if err != nil {
				return seeds, err
			}
			seeds = append(seeds, int(n))
		} else {
			n, err := strconv.ParseInt(string(block[in[i]+1:in[i+1]]), 10, 64)
			if err != nil {
				return seeds, err
			}
			seeds = append(seeds, int(n))
		}
	}
	return seeds, nil
}
