package main

import (
	"bytes"
	"fmt"
	"os"
)

type fileReader struct {
	P string
	o int64
	f *os.File
}

const bufSize = 200

func (fr *fileReader) Open(){
 r, err := os.Open(fr.P) 
 if err != nil {
	fmt.Println(err)
	os.Exit(1)
	}
	fr.f = r
}

func (fr *fileReader) NextLn() (string, error) {
	buf := make([]byte, bufSize)
	fr.f.Seek(fr.o, 0)
	n, err := fr.f.Read(buf)
	if err != nil {
		return string(buf[:n]), err
	}
	lnEnd := bytes.IndexByte(buf, byte('\n'))
	fr.o += int64(lnEnd)+1
	return string(buf[:lnEnd]), nil
}

