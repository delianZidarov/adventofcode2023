package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ArgError struct{}

func (e *ArgError) Error() string {
	return fmt.Sprintf("Please specify whether to solve for part one (-p 1) or part two (-p 2) and file path (-f /path)")
}

type ParseError struct{}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Failed parsing line for int")
}

// Checks the arguments passed in
// If there no appropriate arguments throws error
// Returns true if its part two of the daily challenge
func argumentsParse(args []string) (bool, error) {
	if len(args) != 4 {
		return false, &ArgError{}
	}
	if args[0] != "-p" {
		return false, &ArgError{}
	}
	if args[2] != "-f" {
		return false, &ArgError{}
	}
	number, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
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

// Tries to find the first and last digit in a string
// it looks for the string of single digit numbers as well
// if partTwo is set to true
func getInt(line string, partTwo bool) (int, error) {
	digitArray := make([]int, 2, 2)
	modifiedDigits := 0
	writeDigit := func(value int) {
	if modifiedDigits == 0 {
				digitArray[0] = value
			modifiedDigits++
			} else {
				digitArray[1] = value 
				modifiedDigits++
			}
	}


	for i := 0; i < len(line); i++ {
		number, err := strconv.Atoi(string(line[i]))
		if err == nil {
			writeDigit(number)
		} else if partTwo {
			// imp later
			letter := string(line[i]); 
			switch letter {
			case "o":
				if i+3 <= len(line) && line[i:i+3] == "one" {
					writeDigit(1)
				}
			case "s":
				if i+3 <= len(line) && line[i:i+3] == "six"{
					writeDigit(6)
				} else if i+5 <= len(line) && line[i:i+5] == "seven"{
					writeDigit(7)
				}
			case "t":
				if i+3 <= len(line) && line[i:i+3] == "two" {
					writeDigit(2)
				} else if i+5 <= len(line) && line[i:i+5] == "three" {
					writeDigit(3)
				}
			case "e":
				if i+5 <= len(line) && line[i:i+5] == "eight" {
					writeDigit(8)
				}
			case "n":
				if i+4 <= len(line) && line[i:i+4] == "nine"{
					writeDigit(9)
				}
			case "f":
				if i+4 <= len(line) { 
				word := line[i:i+4]
				if word == "four"{
					writeDigit(4)
				}
				if word == "five"{
					writeDigit(5)
				}
			}
			}
		}
	}
	switch {
	case modifiedDigits < 1:
		return 0, &ParseError{}
	case modifiedDigits == 1:
		return digitArray[0]*10 + digitArray[0], nil
	default:
		return digitArray[0]*10 + digitArray[1], nil
	}
}

func main() {
	sum := 0

	args := os.Args[1:]
	isPartTwo, err := argumentsParse(args)
	if err != nil {
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
			sum += number
		}
	}
	fmt.Println(sum)
}
