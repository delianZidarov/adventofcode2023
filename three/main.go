package main

import (
	"os"
	"fmt"

	"three/arg"
	"three/chkbyte"
)

func main () {
	arg, err := arg.ParseMap(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  m, err := chkbyte.BMatrix(arg["f"])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, ln := range m {
		fmt.Println(string(ln))
	} 
	fmt.Println("")
  fmt.Println(arg["f"], arg["p"])
}
