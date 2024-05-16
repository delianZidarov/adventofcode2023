package main 

import (
	"os"
	"fmt"
)

func main (){
	part, path, pool, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(part,path,pool)
}