package main

import (
	"fmt"
	"io"
	"os"
)

func P7(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	firstLine := true
	beams := make(map[int]bool)
	col := 0
	divisions := 0
	for firstLine {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == '.' {
				col++
				continue
			}
			if chunk[i] == 'S' {
				beams[col] = true
			}
			if chunk[i] == 10 {
				firstLine = false
				fi.Seek(int64(i-127), 1)
				break
			}
		}
	}
	col = 0
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == '.' {
				col++
				continue
			}
			if chunk[i] == '^' {
				if !beams[col] {
					col++
					continue
				}
				delete(beams, col)
				beams[col-1] = true
				beams[col+1] = true
				divisions++
				col++
				continue
			}
			if chunk[i] == 10 {
				if debug {
					fmt.Println(len(beams), beams, col)
				}
				col = 0
			}
		}
	}
	fmt.Println(divisions)
}

func P7_2(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	firstLine := true
	beams := make(map[int]int)
	col := 0
	timelines := 1
	for firstLine {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == '.' {
				col++
				continue
			}
			if chunk[i] == 'S' {
				beams[col] = 1
			}
			if chunk[i] == 10 {
				firstLine = false
				fi.Seek(int64(i-127), 1)
				break
			}
		}
	}
	col = 0
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == '.' {
				col++
				continue
			}
			if chunk[i] == '^' {
				if beams[col] == 0 {
					col++
					continue
				}
				value := beams[col]
				delete(beams, col)
				beams[col-1] = beams[col-1] + value
				beams[col+1] = beams[col+1] + value
				timelines = timelines + value
				col++
				continue
			}
			if chunk[i] == 10 {
				if debug {
					fmt.Println(len(beams), beams, timelines)
				}
				col = 0
			}
		}
	}
	fmt.Println(timelines)
}
