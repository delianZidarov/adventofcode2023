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
	numbers := []byte{'0','1','2','3','4','5','6','7','8','9'}
  for _,n := range numbers{
	 if IsNumber(n) != true {
     t.Fatalf("Expected %v to return true",n)
		}
	}
}
