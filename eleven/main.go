package main

import (
	"fmt"
	"os"
)

func main() {
	points, rowGaps, columnGaps := parse("/home/d/Documents/day11")

	fmt.Println(rowGaps)
	fmt.Println(columnGaps)
	fmt.Println("Star locations: ", points)
}

type ord struct {
	row int
	col int
}

func parse(s string) ([]ord, []int, []int) {
	f, err := os.Open(s)
	defer f.Close()
	if err != nil {
		fmt.Println("Can't open file")
		os.Exit(1)
	}
	buf := make([]byte, 19750)
	end, _ := f.Read(buf)

	rowGaps := make([]int, 0)
	columnGaps := make([]int, 0)
	points := make([]ord, 0)
	lnLength := 0
	metStar := 0
	lastWrite := 0
	m := make([][]byte, 0)

	for i := 0; i < end; i++ {
		if buf[i] == '\n' {
			if lnLength == 0 {
				lnLength = i + 1
			}
			if metStar == 0 {
				ln := 0
				if i > lnLength && lnLength > 0 {
					ln = i / lnLength
				}
				rowGaps = append(rowGaps, ln)
			}
			m = append(m, buf[lastWrite:i])
			lastWrite = i + 1
			metStar = 0
		}
		if buf[i] == '#' {
			row := 0
			if lnLength > 0 && i > lnLength {
        row = i / lnLength
			}
			column := i
			if lnLength > 0 {
      column = i % lnLength
			}
			points = append(points, ord{row: row, col: column})
			metStar += 1
		}
	}
	for i := 0; i < len(m[0]); i++ {
		metStar = 0
		for j := 0; j < len(m)-1; j++ {
			if m[j][i] == '#' {
				metStar += 1
			}
		}
		if metStar == 0 {
			columnGaps = append(columnGaps, i)
		}
	}

	return points, rowGaps, columnGaps
}
