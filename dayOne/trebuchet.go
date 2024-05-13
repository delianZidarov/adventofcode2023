package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)


type ArgError struct {}

func (e *ArgError) Error() string {
	return fmt.Sprintf("Please specify whether to solve for part one (-p 1) or part two (-p 2) and file path (-f /path)")
}

type ParseError struct {}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Failed parsing line for int")
}

func argumentsParse (args []string) (bool, error) {
	//Checks the arguments passed into day one 
	//If there no appropriate arguments throws error
	//Returns true if its part two of the daily challenge
	if len(args) != 4 {
		return false, &ArgError{}
	}
	if args[0] != "-p" {
		return false, &ArgError{}
	}
	if args[2] != "-f"{
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

func getInt(line string, partTwo bool) (int, error){
	digitArray := make([]int, 2, 2)
	modifiedDigits := 0
	
	for i:=0; i< len(line); i++{
		number, err := strconv.Atoi(string(line[i]))
		if err==nil{
			if modifiedDigits==0 {
					digitArray[0]=number
					modifiedDigits++
			} else {
					digitArray[1]=number
					modifiedDigits++
			}
		} else if partTwo{
		//imp later
		}
		}
	switch {
		case modifiedDigits < 1:
			return 0, &ParseError{}
		case modifiedDigits == 1:
			return digitArray[0] * 10 + digitArray[0], nil
		default:
			return digitArray[0] * 10 + digitArray[1], nil
		}

}

func main (){
	sum := 0

	args := os.Args[1:]
	isPartTwo, err := argumentsParse(args)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	readFile, err := os.Open(args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		number, err := getInt(fileScanner.Text(), isPartTwo)
		if err == nil {
			sum+=number
		}
	}

	fmt.Println(sum)
	

}
