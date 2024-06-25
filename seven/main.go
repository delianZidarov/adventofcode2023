package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("ARGS ", args)
	t := insertNode(nil, 1, 2)
	for i := 2; i < 11; i++{
   t = insertNode(t, i,2)
	}
	t = insertNode(t, 0, 2)
	t = insertNode(t,23,2)
	i :=0
	inorderTran(t, &i)
}
