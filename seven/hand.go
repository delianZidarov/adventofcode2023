package main

import (
	"math"
)

func cardValue(c byte) int {
	switch c {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return 0
	}
}

func handValue(h string) int {
	hand := make(map[int]int)
	for _, c := range h {
		val := cardValue(byte(c))
		_, ok := hand[val]
		if ok {
			hand[val] += 1
		} else {
			hand[val] = 1
		}
	}

	switch {
	case len(hand) == 1:
		return 7
	case len(hand) == 2:
		values := make([]int, 0)
		for _, value := range hand {
			values = append(values, value)
		}
		if values[0] == 4 || values[1] == 4 {
			return 6
		} else {
			return 5
		}
	case len(hand) == 3:
		values := make([]int, 0)
		for _, value := range hand {
			values = append(values, value)
		}
		if values[0] == 3 || values[1] == 3 || values[2] == 3 {
			return 4
		} else {
			return 3
		}
	case len(hand) == 4:
		return 2
	default:
		return 1
	}
}

func handScore(h string) int {
	hV := handValue(h)
	one := cardValue(h[0])
	two := cardValue(h[1])
	three := cardValue(h[2])
	four := cardValue(h[3])
	five := cardValue(h[4])
	return hV*int(math.Pow10(10)) + one*int(math.Pow10(8)) +
		two*int(math.Pow10(6)) + three*int(math.Pow10(4)) + four*int(math.Pow10(2)) + five
}

func jackHandValue(h string) int {
	m := make(map[int]int)
	for _, c := range h {
		card := jackCardValue(byte(c))
		_, ok := m[card]
		if ok {
			m[card] += 1
		} else {
			m[card] = 1
		}
	}

	_, ok := m[1]

	if !ok {
		switch {
		case len(m) == 1:
			return 7
		case len(m) == 2:
			values := make([]int, 0)
			for _, value := range m {
				values = append(values, value)
			}
			if values[0] == 4 || values[1] == 4 {
				return 6
			} else {
				return 5
			}
		case len(m) == 3:
			values := make([]int, 0)
			for _, value := range m {
				values = append(values, value)
			}
			if values[0] == 3 || values[1] == 3 || values[2] == 3 {
				return 4
			} else {
				return 3
			}
		case len(m) == 4:
			return 2
		default:
			return 1
		}
	} else {
		values := make([]int, 0)
		for key, value := range m {
			if key != 1 {
				values = append(values, value)
			}
		}

		switch {
		case len(values) == 0:
		  return 7
		case len(values) == 1:
			return 7
		case len(values) == 2:
			if values[0] == 2 && values[1] ==2 {
       return 5
			}
			return 6
		case len(values) == 3:
			return 4
		case len(values) == 4:
			return 2
		}
	}
	return 0
}

func jackCardValue(c byte) int {
	switch c {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 1
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return 0
	}
}

func jackHandScore(h string) int {
	hV := jackHandValue(h)

	one := jackCardValue(h[0])
	two := jackCardValue(h[1])
	three := jackCardValue(h[2])
	four := jackCardValue(h[3])
	five := jackCardValue(h[4])
	
	return hV*int(math.Pow10(10)) + one*int(math.Pow10(8)) +
		two*int(math.Pow10(6)) + three*int(math.Pow10(4)) + four*int(math.Pow10(2)) + five
}
