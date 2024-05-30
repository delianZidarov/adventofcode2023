package arg

import (
	"testing"
	"fmt"
)

type parseMapTest struct {
 in []string
 out map[string]string
 expectError bool
}


var parseMapTests = [] parseMapTest{
	{[]string{}, make(map[string]string) , true},
	{[]string{"/path"},map[string]string{"f":"/path"},false},
	{[]string{"/path", "-p", "1"}, map[string]string{"f":"/path", "p":"1"},false},
	{[]string{"-p", "1", "/path"}, map[string]string{"f":"/path", "p":"1"},false},
	{[]string{"/path", "p", "1"}, map[string]string{"f":"1"},false},

}

func TestParseMap(t *testing.T) {
	for _, test := range parseMapTests {
		v, err := ParseMap(test.in)
		if test.expectError && err == nil{
			t.Errorf("Expected error to be: %v", test.expectError )
		}
		if fmt.Sprint(v) != fmt.Sprint(test.out) {
			t.Errorf("Expected: %q    Recieved: %q",test.out, v )
    }
	}
}
