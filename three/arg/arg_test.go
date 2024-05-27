package arg

import (
	"testing"
)

type parseMapTest struct {
 message string
 in string
 out map[string]string
 expectError bool
}

var parseMapTests = [] parseMapTest{
	parseMapTest{"Expect to fail"}

}

func TestParseMap(t *testing.T) {

	t.Fatal("FAIL")
}
