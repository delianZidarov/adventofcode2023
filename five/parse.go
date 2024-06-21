package main

import (
	"bytes"
	"fmt"
	"strconv"
)

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
				last := 0
				if i+1 == len(m) {
					last = len(m)
				} else {
					last = i
				}
				dest, err := strconv.ParseInt(string(m[lastNewLine+1:spaces[0]]), 10, 64)
				if err != nil {
					fmt.Println("Parse Maps dest: ", string(m[lastNewLine+1:spaces[0]]))
					return nil, err
				}
				source, err := strconv.ParseInt(string(m[spaces[0]+1:spaces[1]]), 10, 64)
				if err != nil {
					fmt.Println("Parse Maps source: ", string(m[spaces[0]+1:spaces[1]]))
					return nil, err
				}
				mod, err := strconv.ParseInt(string(m[spaces[1]+1:last]), 10, 64)
				if err != nil {
					fmt.Println("Parse Maps mod: ", string(m[spaces[1]+1:last]), i, spaces)
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
