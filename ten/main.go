package main

import (
	"fmt"
	"os"
)

func main() {
	// TEST
	// m, c := parseInput("/home/d/Documents/test")
	// TEST
	m, c := parseInput("/home/d/Documents/day10")
	h := make([]cord, 0)
	h = append(h, c)
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range dir {
		result, ok := route(cord{row: c.row + d[0], col: c.col + d[1]}, &m, h)
		if ok {
			fmt.Println("Part One: ", len(result)/2)
			// pick's theorem: given the integer cordinates of a shapes border (b)
			// and the number of points inside (i); we can calculate the area of a
			// shape: A = i + (b/2) -1
			// shoelace formula: way to calculate the area of a simple polygon by
			// cross multiplying and subtracting vertices: .5 * x1 * y2 - y1 * x2 ... xn * y1 - yn * x1
			// once the area is calculated we can solve for i, the answer to part 2
			// i = A - (b/2) + 1
			fmt.Println("Part Two: ", lace(result)-(len(result)/2)+1)
			break
		}
	}
}

type cord struct {
	row int
	col int
}

func route(c cord, m *[][]byte, history []cord) (h []cord, ok bool) {
	current := (*m)[c.row][c.col]
	if current == 'S' {
		h = history
		return h, true
	}

	if !validMove(current,
		cord{
			row: c.row - history[len(history)-1].row,
			col: c.col - history[len(history)-1].col,
		}) {
		return h, false
	}

	n := nextCord(current, cord{
		row: c.row - history[len(history)-1].row,
		col: c.col - history[len(history)-1].col,
	})

	if n.row == 0 && n.col == 0 {
		fmt.Println("Got stuck at: ", current, c)
		return h, false
	}

	history = append(history, c)

	return route(cord{row: c.row + n.row, col: c.col + n.col}, m, history)
}

func validMove(c byte, d cord) bool {
	switch {
	// going down
	case d.row == 1 && d.col == 0:
		if c == '|' || c == 'L' || c == 'J' {
			return true
		} else {
			return false
		}
		// going up
	case d.row == -1 && d.col == 0:
		if c == '|' || c == '7' || c == 'F' {
			return true
		} else {
			return false
		}
		// going right
	case d.row == 0 && d.col == 1:
		if c == '-' || c == '7' || c == 'J' {
			return true
		} else {
			return false
		}
		// going left
	case d.row == 0 && d.col == -1:
		if c == '-' || c == 'L' || c == 'F' {
			return true
		} else {
			return false
		}
	}
	return false
}

func nextCord(c byte, d cord) cord {
	switch {
	// going down
	case d.row == 1 && d.col == 0:
		if c == '|' {
			return cord{row: 1, col: 0}
		}
		if c == 'L' {
			return cord{row: 0, col: 1}
		}
		if c == 'J' {
			return cord{row: 0, col: -1}
		}
		// going up
	case d.row == -1 && d.col == 0:
		if c == '|' {
			return cord{row: -1, col: 0}
		}
		if c == '7' {
			return cord{row: 0, col: -1}
		}
		if c == 'F' {
			return cord{row: 0, col: 1}
		}
		// going right
	case d.row == 0 && d.col == 1:
		if c == '-' {
			return cord{row: 0, col: 1}
		}
		if c == '7' {
			return cord{row: 1, col: 0}
		}
		if c == 'J' {
			return cord{row: -1, col: 0}
		}
		// going left
	case d.row == 0 && d.col == -1:
		if c == '-' {
			return cord{row: 0, col: -1}
		}
		if c == 'L' {
			return cord{row: -1, col: 0}
		}
		if c == 'F' {
			return cord{row: 1, col: 0}
		}
	}
	return cord{row: 0, col: 0}
}

func parseInput(s string) ([][]byte, cord) {
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
			row = int(i / lnLength)
			col = int(i % lnLength)
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

	return m, cord{row, col}
}

func lace(vert []cord) int {
	vert = append(vert, vert[0])
	area := 0
	for i := 0; i < len(vert) -1 ; i++ {
		// the last vert is cross multiplied with the first
			area = area + vert[i].col*vert[i+1].row - vert[i].row*vert[i+1].col
	}
	if area < 0 {
		return (area * -1) / 2
	} else {
		return (area) / 2
	}
}
