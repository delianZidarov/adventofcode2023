package main

import (
	"fmt"
	"os"
)

func main() {
	m, row, col := parseInput("/home/d/Documents/day10")


	fmt.Println("Is this the start = ", string(m[row][col]))
}

func parseInput(s string) ([][]byte,  int,  int) {
	f, _ := os.Open(s)
	defer f.Close()
	buf := make([]byte, 19750)
	end, _ := f.Read(buf)
	lnLength := 0
	row := 0
	col := 0
	var m [][]byte
	lastWrite := 0
	for i, c := range buf[:end] {
if c == 'S' {
    row = int(i/lnLength)
		col = int(i%lnLength)
		}
		if c == '\n' {
			if lnLength == 0 {
        lnLength = i + 1
			}
			m = append(m, buf[lastWrite:i])
			lastWrite = i + 1
		}
	}
	m = append(m, buf[lastWrite:end])

	return m, row, col
}

