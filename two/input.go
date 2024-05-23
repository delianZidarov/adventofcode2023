package main

import (
	"fmt"
	"strconv"
)

type (
	ArgLenError struct{}
	InputError  struct {
		argument string
		value    string
	}
)

func (e *ArgLenError) Error() string {
	return fmt.Sprintf("Not enough arguments were provided. -p (part) -f (file) -r (red) -g (green) -b (blue) required")
}

func (e *InputError) Error() string {
	return fmt.Sprintf("There is a problem with argument %v and value %v", e.argument, e.value)
}

func parseInput(arg []string) (part int64, path string, pool [3]int64, e error) {
	if len(arg) != 10 {
		e = &ArgLenError{}
		return
	}
	for i := 0; i < 9; i++ {
		switch arg[i] {
		case "-p":
			p, err := strconv.ParseInt(arg[i+1], 10, 64)
			if err != nil {
				e = &InputError{argument: arg[i], value: arg[i+1]}
				i++
				break
			}
			part = p
			i++
		case "-f":
			path = arg[i+1]
			i++
		case "-r":
			r, err := strconv.ParseInt(arg[i+1], 10, 64)
			if err != nil {
				e = &InputError{argument: arg[i], value: arg[i+1]}
				i++
			}
			pool[0] = r
			i++
		case "-g":
			g, err := strconv.ParseInt(arg[i+1], 10, 64)
			if err != nil {
				e = &InputError{argument: arg[i], value: arg[i+1]}
				i++
				break
			}
			pool[1] = g
			i++
		case "-b":
			b, err := strconv.ParseInt(arg[i+1], 10, 64)
			if err != nil {
				e = &InputError{argument: arg[i], value: arg[i+1]}
				i++
				break
			}
			pool[2] = b
			i++
		default:
			e = &InputError{argument: arg[i], value: arg[i+1]}
			i++
		}
	}
	return
}
