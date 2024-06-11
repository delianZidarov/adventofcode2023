package main

import (
	"bytes"
	"fmt"
	"os"
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
	blocks(buf)

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
		if len(mid[i]) > 0{
			chunks = append(chunks, mid[i])
		 fmt.Println(string(mid[i]))
		}
	}
	return chunks
}
