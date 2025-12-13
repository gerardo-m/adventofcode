package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func P1(fi *os.File) {
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
		// fmt.Println(line + " " + dir + " " + strconv.Itoa(rot))
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

func P1_2(fi *os.File) {
	var a string
	fmt.Fscanln(fi, &a)
	fmt.Println("bye" + a)
}
