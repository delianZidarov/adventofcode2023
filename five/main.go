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

	buf := make([]byte, 5745)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Read(buf)
	b := blocks(buf)
	f.Close()

	seeds, err := parseSeeds(b[0])
	if err != nil {
		fmt.Println(err)
	}
	seedToSoil, err := parseInputMap(b[1])
	if err != nil {
		fmt.Println(err)
	}

	soilToFert, err := parseInputMap(b[2])
	if err != nil {
		fmt.Println(err)
	}

	fertToWater, err := parseInputMap(b[3])
	if err != nil {
		fmt.Println(err)
	}

	waterToLight, err := parseInputMap(b[4])
	if err != nil {
		fmt.Println(err)
	}

	lightToTemp, err := parseInputMap(b[5])
	if err != nil {
		fmt.Println(err)
	}

	tempToHumid, err := parseInputMap(b[6])
	if err != nil {
		fmt.Println(err)
	}

	humidToLoc, err := parseInputMap(b[7])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(seedToSoil, soilToFert, fertToWater, waterToLight,
		lightToTemp, tempToHumid, humidToLoc, seeds)

	switch p {
	case 1:
		fmt.Println("Part 1")
	case 2:
		fmt.Println("Part 2")
	}
}

func blocks(buf []byte) (chunks [][]byte) {
	mid := bytes.Split(buf, []byte("\n\n"))
	for i := 0; i < len(mid); i++ {
		if len(mid[i]) > 0 {
			chunks = append(chunks, mid[i])
		}
	}
	return chunks
}

func parseInputMap(m []byte) (n *Node, err error) {
	lastNewLine := 0
	spaces := make([]int, 2)
	location := 0
	for i := 0; i < len(m); i++ {
		switch {
		case m[i] == ' ':
			spaces[location%2] = i
			location += 1
		case m[i] == '\n' || i+1 == len(m):
			if lastNewLine > 0 {
				dest, err := strconv.ParseInt(string(m[lastNewLine+1:spaces[0]]), 10, 64)
				if err != nil {
					return nil, err
				}
				source, err := strconv.ParseInt(string(m[spaces[0]+1:spaces[1]]), 10, 64)
				if err != nil {
					return nil, err
				}
				mod, err := strconv.ParseInt(string(m[spaces[1]+1:i]), 10, 64)
				if err != nil {
					return nil, err
				}
				n = insertNode(n, int(dest), int(source), int(source+mod-1))
			}
			lastNewLine = i
			location = 0
		}
	}
	return n, nil
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
