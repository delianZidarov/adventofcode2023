package main 

import (
 "fmt"
	"os"
	"bufio"
)

func main () {
 arg := os.Args [1:]
   fp:= arg[0]
	  var p int
 if len(arg) == 1 {
		p=1
	}else {
    p= 2
	}
	
	f, err := os.Open(fp)
	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if p==1{
  partOne(scanner)
	}

	partTwo(scanner)
}
