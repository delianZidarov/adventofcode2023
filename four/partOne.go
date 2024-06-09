package main

import (
	"bufio"
	"fmt"
	"strings"
)

func partOne(s *bufio.Scanner) {
	sum := 0
	for s.Scan() {
		card := s.Text()
		n := value(card)
sum += n
	}
	fmt.Println("TOTAL: ", sum)
}

func value(s string) (v int) {
	if len(s) == 0 {
		return
	}

	cDelim := false
	wDelim := false
	n := 0
	m := make(map[string]int)
	stS := strings.Split(s, " ")
	for i := 0; i < len(stS); i++ {
		switch {
		case stS[i] == "":
			continue
		case stS[i][len(stS[i])-1] == ':':
			cDelim = true

		case stS[i] == "|":
			wDelim = true

		case cDelim && !wDelim:
			m[stS[i]] = 1

		case cDelim && wDelim:
			_, ok := m[stS[i]]
			if ok {
				n += 1
			}
		}
	}
	switch n {
	case 0:
		v = 0
	case 1:
		v = 1
	default:
		v = 2 << (n - 2)

	}
	return
}
