package arg

import "fmt"

type LengthError struct {}

func (e *LengthError) Error() string {
return fmt.Sprintf("No arguments provided")
}

//Makes a map of flags and the following string from a slice.
//If there is no preceding flag to a string it gives it a key
//of f. Only one none flag value can be used. 
func ParseMap (a []string ) (map[string]string, error){
	m := make(map[string]string)
	if len(a) < 1 {
		return m, &LengthError{}
	}
	for i:=0; i < len(a); i++ {
		switch {
			case a[i][0] == '-':
				m[a[i][1:]] = a[i+1]
				if i < len(a) {
				i++
			} else {
				break
			}
		default:
		m["f"] = a[i]
		}
	}
	return m, nil
}
