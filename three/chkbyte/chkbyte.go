package chkbyte

import (
	"os"
	"bytes"
)

func BMatrix (f string) (matrix [][]byte, err error) {
	fl, err := os.Open(f)
	defer fl.Close()
	if err != nil {
		return matrix, err
	}
	//19740 is the size of day3s input
	data := make([]byte, 19740)
  _, err = fl.Read(data)
	if err != nil {
		return matrix, err
	}
	return bytes.Split(data,[]byte("\n")), nil
}
