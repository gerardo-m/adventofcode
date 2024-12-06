package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func P6() {
	fi, err := os.Open("input6.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var pos, lineL int
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lineL = len(line)
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		pos = strings.Index(line, "^")
		if pos != -1 {
			pos = lineL*i + pos
			break
		}
	}
	movement := []int{
		-lineL,
		1,
		lineL,
		-1,
	}
	currMov := 0
	currPos := pos
	visitedPos := make(map[int]bool)
	adjustedLineL := lineL + 2
	for {
		visitedPos[currPos] = true
		if currPos%lineL == 0 && movement[currMov] == -1 {
			break
		}
		if currPos%lineL == lineL-1 && movement[currMov] == 1 {
			break
		}
		futurePos := currPos + movement[currMov]
		offset := (futurePos/lineL)*adjustedLineL + (futurePos % lineL)
		val := make([]byte, 1)
		_, err := fi.ReadAt(val, int64(offset))
		if err == io.EOF {
			break
		}
		if val[0] == '#' {
			currMov = (currMov + 1) % len(movement)
			continue
		}
		currPos = futurePos
		if currPos < 0 {
			break
		}
	}
	fmt.Println(len(visitedPos))
}
