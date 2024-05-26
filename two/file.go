package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
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
	dataStart := 0
	pointMultiplier := 1
	for i := 0; i < len(ln); i++ {
		if ln[i] == ' ' && points == 0 {
			i++
			for ln[i] != ':' {
				buf = append(buf, ln[i])
				i++
			}
			p, err := strconv.ParseInt(string(buf[dataStart:]), 10, 64)
			if err != nil {
				return 0, err
			}
			points = p
		} else if ln[i] == ' ' {
			i++
			dataStart = len(buf)
			for ln[i] != ' ' {
				buf = append(buf, ln[i])
				i++
			}
			count, err := strconv.ParseInt(string(buf[dataStart:]), 10, 64)
			if err != nil {
				return 0, err
			}
			switch ln[i+1] {
			case 'r':
				if pool[0] < count {
					pointMultiplier = 0
				}
			case 'g':
				if pool[1] < count {
					pointMultiplier = 0
				}
			case 'b':
				if pool[2] < count {
					pointMultiplier = 0
				}

			}

		}
	}
	return int(points) * pointMultiplier, nil
}

func Power(ln string) (int, error) {
	var (
		minR int64 = 0
		minG int64 = 0
		minB int64 = 0
	)
	buf := make([]byte, len(ln))
	start := false
	dataStart := 0
	for i := 0; i < len(ln); i++ {
		if ln[i] == ':' {
			start = true
		}
		if ln[i] == ' ' && start {
			i++
			dataStart = len(buf)
			for ln[i] != ' ' {
				buf = append(buf, ln[i])
				i++
			}
			v, err := strconv.ParseInt(string(buf[dataStart:]), 10, 64)
			if err != nil {
				return 0, err
			}
			switch ln[i+1] {
			case 'r':
				if v > minR {
					minR = v
				}
			case 'g':
				if v > minG {
					minG = v
				}
			case 'b':
				if v > minB {
					minB = v
				}

			}
		}

	}
	return int(minR * minG * minB), nil
}
