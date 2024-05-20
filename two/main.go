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
	f:= fileReader{P : path}
	f.Open()
	for {
		ln, err := f.NextLn()
		if err!=nil{
		break
		}
		fmt.Println(ln)
	}
	fmt.Println(part,path,pool)
}
