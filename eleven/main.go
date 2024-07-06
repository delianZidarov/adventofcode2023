package main

import (
	"fmt"
	"os"
)

func main() {
	// TEST
	//points, rowGaps, columnGaps := parse("/home/d/Documents/test")
	// TEST

	points, rowGaps, columnGaps := parse("/home/d/Documents/day11")

	expanded := expandOrd(1, points, rowGaps, columnGaps)
	fmt.Println("Part 1: ", sumOfDistance(expanded))
	fmt.Println("Part 2: ", sumOfDistance(expandOrd(1000000, points, rowGaps, columnGaps)))
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
		for j := 0; j < len(m); j++ {
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

func expandOrd(eFactor int, startingPoints []ord, rowGaps []int, columnGaps []int) []ord {
	newPoints := make([]ord, 0)
	for _, p := range startingPoints {
		rowOffset := 0
		columnOffset := 0
		for _, i := range rowGaps {
			if p.row > i {
				rowOffset += 1
			} else {
				break
			}
		}

		for _, i := range columnGaps {
			if p.col > i {
				columnOffset += 1
			} else {
				break
			}
		}
		if eFactor > 1 {
			newPoints = append(newPoints, ord{row: p.row + (rowOffset * (eFactor - 1)), col: p.col + (columnOffset * (eFactor - 1))})
		} else {
			newPoints = append(newPoints, ord{row: p.row + (rowOffset * eFactor), col: p.col + (columnOffset*eFactor - 1)})
		}
	}
	return newPoints
}

func sumOfDistance(points []ord) int {
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			sum = sum + ordDistance(points[i], points[j])
		}
	}
	return sum
}

func ordDistance(a, b ord) int {
	return abs(a.col-b.col) + abs(a.row-b.row)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
