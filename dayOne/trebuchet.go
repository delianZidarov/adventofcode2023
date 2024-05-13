package main

import (
	"fmt"
	"os"
	"strconv"
)

type ArgError struct {}

func (e *ArgError) Error() string {
	return fmt.Sprintf("Please specify whether to solve for part one (-p 1) or part two (-p 2)")
}

func argumentsParse (args []string) (bool, error) {
	//Checks the arguments passed into day one 
	//If there no appropriate arguments throws error
	//Returns true if its part two of the daily challenge
	if len(args) != 2 {
		return false, &ArgError{}
	}
	if args[0] != "-p" {
		return false, &ArgError{}
	}
	number, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil{
		return false, &ArgError{}
	}
	if number == 1 {
		return false, nil
	}
	if number == 2 {
		return true, nil
	}
	return false, &ArgError{}
}

func main (){
	args := os.Args[1:]
	isPartTwo, err := argumentsParse(args)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(isPartTwo)

}
