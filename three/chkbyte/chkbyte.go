package chkbyte

import (
	"bytes"
	"os"
)

//Tries to open a file at f and returns a 2D slice
//seperated by new lines
func BMatrix(f string) (matrix [][]byte, err error) {
	fl, err := os.Open(f)
	defer fl.Close()
	if err != nil {
		return matrix, err
	}
	// 19740 is the size of day3s input
	data := make([]byte, 19740)
	_, err = fl.Read(data)
	if err != nil {
		return matrix, err
	}
	return bytes.Split(data, []byte("\n")), nil
}

//Tests if a byte is a number
func IsNumber(n byte) (isnumber bool) {
	//48 is 0 in ASCII and 57 is 9
	if n > 47 && n < 58 {
		isnumber = true
	}
	return
}

//Tests if a byte is equal to a period
func IsDot(d byte) (isdot bool) {
	if d == '.' {
		isdot = true
	}
	return
}

//Checks all neighbors to position r(ow) c(olumn) in
//provided matrix for bytes that can be considered symbols
//symbols are non-numbers and not a period
func HasSymbolNeighbor(r int, c int, matrix *[][]byte) (foundSymbol bool) {
	directions := [][]int{
		{0, 1},{0, -1},{1, 1},{1, -1},{-1, 1},{-1, -1},{1, 0},{-1, 0},
	}
	for _, d := range directions {
		nR := r + d[1]
		nC := c + d[0]
		if nR < 0 || nR > len(*matrix)-1 {
			continue
		}
		if nC < 0 || nC > len((*matrix)[nR])-1 {
			continue
		}
		if !IsDot((*matrix)[nR][nC]) && !IsNumber((*matrix)[nR][nC]) {
			foundSymbol = true
		}

	}

	return foundSymbol
}
