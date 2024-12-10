package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
)

func P8() {
	fi, err := os.Open("input8.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	var lineL, adjustedLineL int
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineL = len(line) // Adjusted for end of line
		adjustedLineL = lineL + 2
		break
	}
	fi.Seek(0, 0)
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanBytes)
	anten := make(map[byte]map[int]bool)
	antinodes := make(map[int]bool)
	i := 0
	for ; scanner.Scan(); i++ {
		val := scanner.Bytes()
		err := scanner.Err()
		if err == io.EOF {
			break
		}
		if val[0] == '\r' || val[0] == '\n' || val[0] == '.' {
			continue
		}
		prevAnt, exists := anten[val[0]]
		if !exists {
			anten[val[0]] = make(map[int]bool)
		} else {
			for a := range maps.Keys(prevAnt) {
				ix, iy := i%adjustedLineL, i/adjustedLineL
				ax, ay := a%adjustedLineL, a/adjustedLineL
				difx := ix - ax
				dify := iy - ay
				ant1x, ant1y := ax-difx, ay-dify
				if ant1x >= 0 && ant1x < (lineL) && ant1y >= 0 {
					antinodes[ant1y*lineL+ant1x] = true
				}
				ant2x, ant2y := ix+difx, iy+dify
				if ant2x >= 0 && ant2x < (lineL) && ant2y >= 0 {
					antinodes[ant2y*lineL+ant2x] = true
				}
				// fmt.Printf("prev %v %v, curr %v %v\n", ay, ax, iy, ix)
				// fmt.Printf("ant 1: %v %v, ant 2: %v %v\n", ant1y, ant1x, ant2y, ant2x)
			}
		}
		anten[val[0]][i] = true
	}
	lineCount := (i + 2) / adjustedLineL
	totalB := lineCount * lineL

	for a := range maps.Keys(antinodes) {
		if a >= totalB {
			delete(antinodes, a)
		}
	}
	fmt.Println(len(antinodes))
}

func P8_1() {
	fi, err := os.Open("input8.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	var lineL, adjustedLineL, lineCount int
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for lineCount = 0; scanner.Scan(); lineCount++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lineL = len(line) // Adjusted for end of line
		adjustedLineL = lineL + 2

	}
	fi.Seek(0, 0)
	scanner = bufio.NewScanner(fi)
	scanner.Split(bufio.ScanBytes)
	anten := make(map[byte]map[int]bool)
	antinodes := make(map[int]bool)
	i := 0
	for ; scanner.Scan(); i++ {
		val := scanner.Bytes()
		err := scanner.Err()
		if err == io.EOF {
			break
		}
		if val[0] == '\r' || val[0] == '\n' || val[0] == '.' {
			continue
		}
		prevAnt, exists := anten[val[0]]
		if !exists {
			anten[val[0]] = make(map[int]bool)
		} else {
			for a := range maps.Keys(prevAnt) {
				ix, iy := i%adjustedLineL, i/adjustedLineL
				ax, ay := a%adjustedLineL, a/adjustedLineL
				difx := ix - ax
				dify := iy - ay
				for ant1x, ant1y := ax-difx, ay-dify; ant1x >= 0 && ant1x < (lineL) && ant1y >= 0 && ant1y < lineCount; ant1x, ant1y = ant1x-difx, ant1y-dify {
					antinodes[ant1y*lineL+ant1x] = true
				}
				for ant2x, ant2y := ix+difx, iy+dify; ant2x >= 0 && ant2x < (lineL) && ant2y >= 0 && ant2y < lineCount; ant2x, ant2y = ant2x+difx, ant2y+dify {
					antinodes[ant2y*lineL+ant2x] = true
				}
				antinodes[ay*lineL+ax] = true
				antinodes[iy*lineL+ix] = true
				// fmt.Printf("prev %v %v, curr %v %v\n", ay, ax, iy, ix)
				// fmt.Printf("ant 1: %v %v, ant 2: %v %v\n", ant1y, ant1x, ant2y, ant2x)
			}
		}
		anten[val[0]][i] = true
	}
	totalB := lineCount * lineL

	for a := range maps.Keys(antinodes) {
		if a >= totalB {
			delete(antinodes, a)
		}
	}
	fmt.Println(len(antinodes))
}
