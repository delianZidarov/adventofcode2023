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
  p, ok := arg["p"]
	if p=="1" || !ok{
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
	}
	if ok && p=="2"{
		for row := 0; row < len(m); row ++ {
      for column := 0; column < len(m[row]); column ++{
	     if chkbyte.IsAsterix(m[row][column]){
         n := chkbyte.CheckNumberNeighbor(row, column, &m)
					if len(n) == 2 {
						s := 1
						for _, pair := range n{
							n, err := chkbyte.Number(pair[0], pair[1], &m)
							if err != nil {
               fmt.Println(err)
								os.Exit(1)
							}
							s = s * n
						}
						sum += s
					}
				}
		}
		}
	}
	fmt.Println("THE SUM IS: ",sum)
}
