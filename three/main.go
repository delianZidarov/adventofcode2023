package main

import (
	"os"
	"fmt"

	"three/arg"
)

func main () {
	arg, err := arg.ParseMap(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  fmt.Println(arg["f"], arg["p"])
}
