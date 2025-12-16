package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func P1(fi *os.File, debug bool) {
	var line string
	pos := 50
	zeroCounter := 0
	for {
		_, err := fmt.Fscanln(fi, &line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		dir := line[:1]
		rot, _ := strconv.Atoi(line[1:])
		if dir == "L" {
			rot = 100 - (rot % 100)
		}
		pos = pos + rot
		pos = pos % 100
		if pos == 0 {
			zeroCounter++
		}
	}
	fmt.Println(zeroCounter)
}

func P1_2(fi *os.File, debug bool) {
	var line string
	pos := 50
	zeroCounter := 0
	atZero := false
	for {
		_, err := fmt.Fscanln(fi, &line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		dir := line[:1]
		rot, _ := strconv.Atoi(line[1:])
		atZero = pos == 0
		if dir == "L" {
			pos = pos - rot
			for pos < 0 {
				pos = pos + 100
				if atZero {
					atZero = false
				} else {
					zeroCounter++
				}
			}
			if pos == 0 {
				zeroCounter++
			}
		} else {
			pos = pos + rot
			for pos > 99 {
				pos = pos - 100
				zeroCounter++
			}
		}
	}
	fmt.Println(zeroCounter)
}
