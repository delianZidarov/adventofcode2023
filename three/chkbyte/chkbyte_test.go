package chkbyte

import (
	"testing"
)

func TestIsDot(t *testing.T) {
	for i := 0; i < 128; i++ {
		if i == 46 {
			want := true
			got := IsDot(byte(i))
			if want != got {
				t.Fatalf("Expected %v to return true", byte(i))
			}
		} else {
			want := false
			got := IsDot(byte(i))
			if want != got {
				t.Fatalf("Expected %v to return false", byte(i))
			}
		}
	}
	want := true
	got := IsDot('.')
	if want != got {
		t.Fatalf("Expected '.' to return true")
	}
}

func TestIsNumber(t *testing.T) {
	numbers := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, n := range numbers {
		if IsNumber(n) != true {
			t.Fatalf("Expected %v to return true", n)
		}
	}
	for i := 0; i < 128; i++ {
		if i < 48 || i > 57 {
			if IsNumber(byte(i)) == true {
				t.Fatalf("Expected %v to return false", i)
			}
		}
	}
}

func TestIsAsterix(t *testing.T) {
	for i := 0; i < 128; i++ {
		if i == 42 && !IsAsterix(byte(i)) {
			t.Fatalf("Expected %v to return true", string(byte(i)))
		} else if i != 42 && IsAsterix(byte(i)) {
			t.Fatalf("%v Expected %v to return false", i, string(byte(i)))
		}
	}
}

type hasSymbolNeighborTest struct {
	row      int
	column   int
	expected bool
}

var hasSymbolNeighborTests = []hasSymbolNeighborTest{
	{0, 1, false},
	{0, 3, false},
	{0, 4, true},
	{2, 2, true},
	{2, 5, false},
	{4, 0, true},
}

func TestHasSymbolNeighbor(t *testing.T) {
	testMatrix := [][]byte{
		// 0   1   2   3   4   5
		{'.', '4', '.', '2', '3', '-'}, // 0
		{'.', '.', '.', '.', '.', '.'}, // 1
		{'.', '+', '5', '.', '.', '1'}, // 2
		{'.', '.', '.', '.', '.', '.'}, // 3
		{'2', '.', '.', '.', '.', '.'}, // 4
		{'.', '/', '.', '.', '.', '.'}, // 5
	}
	for _, test := range hasSymbolNeighborTests {
		got := HasSymbolNeighbor(test.row, test.column, &testMatrix)
		if got != test.expected {
			t.Fatalf("Expected hasNeighbor to return % v got %v", test.expected, got)
		}
	}
}

type numberTest struct {
	row    int
	column int
	expect int
}

var numberTests = []numberTest{
	{0, 1, 4},
	{0, 5, 231},
	{5, 4, 246},
	{4, 0, 2},
}

func TestNumber(t *testing.T) {
	testMatrix := [][]byte{
		// 0    1    2    3    4    5
		{'.', '4', '.', '2', '3', '1'}, // 0
		{'.', '.', '.', '.', '.', '.'}, // 1
		{'.', '+', '5', '.', '.', '.'}, // 2
		{'.', '.', '.', '.', '.', '.'}, // 3
		{'2', '.', '.', '.', '.', '.'}, // 4
		{'.', '/', '.', '2', '4', '6'}, // 5
	}
	for _, test := range numberTests {
		want := test.expect
		got, err := Number(test.row, test.column, &testMatrix)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if want != got {
			t.Fatalf("Expected to get number %v got %v", want, got)
		}
	}
}

type numberNeighborTest struct {
	row    int
	column int
	expext [][]int
}

var numberNeighborTests = []numberNeighborTest {
	{},
  {},
  {},
  {},
}

func TestCheckNumberNeighbor (t *testing.T) {
	testMatrix := [][]byte{
		// 0    1    2    3    4    5
		{'.', '4', '.', '2', '3', '1'}, // 0
		{'.', '.', '.', '.', '.', '.'}, // 1
		{'.', '+', '5', '.', '.', '.'}, // 2
		{'.', '.', '.', '.', '.', '.'}, // 3
		{'2', '.', '.', '.', '.', '.'}, // 4
		{'.', '/', '.', '2', '4', '6'}, // 5
	}
}
