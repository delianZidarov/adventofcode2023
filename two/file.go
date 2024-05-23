package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strings"
)

type fileReader struct {
	P string
	o int64
	f *os.File
}

const bufSize = 200

func (fr *fileReader) Open() {
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
	fr.o += int64(lnEnd) + 1
	return string(buf[:lnEnd]), nil
}

func Score(ln string, pool *[3]int64) (int, error) {
	var points int64
	buf := make([]byte, 0, len(ln))
	for i, ch := range ln {
		if ch == ' ' && points == 0 {
			for ln[i] != ':' {
				buf = append(buf, ln[i])
				i++
			}
			p, err := strconv.ParseInt(strings.TrimSpace(string(buf)), 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			points = p
		}
	}
	//	fmt.Println(games, points, *pool)
	fmt.Println(points)
	//just checking github
	return 0, nil
}
