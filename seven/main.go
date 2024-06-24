package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	a := handScore("22222")
	fmt.Println("ARGS ", args)
	fmt.Println(a)
}
