package main

import (
	"fmt"
	"os"
	"strconv"
	"three/arg"
	"three/chkbyte"
)

func main() {
	arg, err := arg.ParseMap(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m, err := chkbyte.BMatrix(arg["f"])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := 0
	for row := 0; row < len(m)-1; row++ {
		// The starting index of consecutive numbers
		nStart := 0
		// Holds whether any of the numbers had a symbol
		// neighbor
		hs := false
		// Holds whether the next number encountered is the
		// start of a new cosecutive number string
		newN := true
		for column := 0; column < len(m[row]); column++ {
			if chkbyte.IsNumber(m[row][column]) && newN {
				newN = false
				nStart = column
				// Propagate if any cosecutive number had a symbol neighbor
				hs = hs || chkbyte.HasSymbolNeighbor(row, column, &m)
			} else if chkbyte.IsNumber(m[row][column]) {
				// Propagate if any cosecutive number had a symbol neighbor
				hs = hs || chkbyte.HasSymbolNeighbor(row, column, &m)
			} else if !chkbyte.IsNumber(m[row][column]) {
				if hs {
					n, err := strconv.ParseInt(string(m[row][nStart:column]), 10, 64)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					sum += int(n)
				}
				newN = true
				hs = false
				// Handle the end of a line
			} 
			//This should run at the end of every line
			if column == len(m[row])-1 {
				if hs {
					n, err := strconv.ParseInt(string(m[row][nStart:]), 10, 64)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					sum += int(n)
				}
				newN = true
				hs = false
			}
		}
	}
  

type numberTest struct {
   row int
	 column int
	 expect int
}

var numberTests = []numberTest{
	{0,1,4},
  {0,5,231},
	{5,4,246},
  {4,0, 2},
}
	testMatrix := [][]byte{
		//0    1    2    3    4    5
		{'.', '4', '.', '2', '3', '1'}, // 0
		{'.', '.', '.', '.', '.', '.'}, // 1
		{'.', '+', '5', '.', '.', '.'}, // 2
		{'.', '.', '.', '.', '.', '.'}, // 3
		{'2', '.', '.', '.', '.', '.'}, // 4
		{'.', '/', '.', '2', '4', '6'}, // 5
	}

	fmt.Println("*********************************************")
	for _,test := range numberTests {
		want := test.expect
		got, err := chkbyte.Number(test.row, test.column, &testMatrix)
		fmt.Println("Am I in an infinite", got)
		if err != nil {
		 fmt.Println("Unexpected error: ", err)
		}
		if want != got {
     fmt.Println("Expected to get number ", want," got ", got)
		}
	}
	fmt.Println("*********************************************")
	fmt.Println("THE SUM IS: ",sum)
}
