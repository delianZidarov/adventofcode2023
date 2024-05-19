package main

import (
	"bytes"
	"fmt"
	"os"
)

const buffSize = 200

func fileReader (p string){
	r, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buff := make([]byte, buffSize)
	for i:=0;i<5;i++ {
		n, err := r.Read(buff)
		if err != nil {
			break
		}
		lnEnd := bytes.IndexByte(buff[:n], byte('\n'))
		fmt.Println("here", string(buff[:lnEnd]), bytes.IndexByte(buff[:n],byte('\n')), buff[lnEnd+2])
		r.Seek(int64(lnEnd - buffSize + 1), 1)
	}
}
