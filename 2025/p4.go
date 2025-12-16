package main

import (
	"fmt"
	"io"
	"os"
)

func P4(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	adjCount := make(map[int]int)
	ind := 0
	lineL := 0
	row := 0
	col := 0
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
			if chunk[i] == 10 {
				lineL = col
				col = 0
				row++
				continue
			}
			if chunk[i] == '.' {
				ind++
				col++
				continue
			}
			adjCount[ind] = 0
			if row > 0 {
				if col > 0 {
					updateRolls(adjCount, ind-lineL-1, ind)
				}
				updateRolls(adjCount, ind-lineL, ind)
				if col < lineL-1 {
					updateRolls(adjCount, ind-lineL+1, ind)
				}
			}
			if col > 0 {
				updateRolls(adjCount, ind-1, ind)
			}
			col++
			ind++
		}
	}
	count := 0
	if debug {
		printRolls(ind+1, lineL, adjCount)
	}
	for _, val := range adjCount {
		if val < 4 {
			count++
		}
	}
	fmt.Println(count)
}

func updateRolls(adjCount map[int]int, i1 int, i2 int) {
	_, ok := adjCount[i1]
	if ok {
		adjCount[i2] = adjCount[i2] + 1
		adjCount[i1] = adjCount[i1] + 1
	}
}

func P4_2(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	adjCount := make(map[int]int)
	ind := 0
	lineL := 0
	row := 0
	col := 0
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
			if chunk[i] == 10 {
				lineL = col
				col = 0
				row++
				continue
			}
			if chunk[i] == '.' {
				ind++
				col++
				continue
			}
			adjCount[ind] = 0
			if row > 0 {
				if col > 0 {
					updateRolls(adjCount, ind-lineL-1, ind)
				}
				updateRolls(adjCount, ind-lineL, ind)
				if col < lineL-1 {
					updateRolls(adjCount, ind-lineL+1, ind)
				}
			}
			if col > 0 {
				updateRolls(adjCount, ind-1, ind)
			}
			col++
			ind++
		}
	}

	count := 1
	totalCount := 0
	for count > 0 {
		count = 0
		if debug {
			printRolls(ind+1, lineL, adjCount)
		}
		for k, val := range adjCount {
			if val < 4 {
				count++
				totalCount++
				delete(adjCount, k)
			} else {
				adjCount[k] = 0
			}
		}
		for k := range adjCount {
			row = k / lineL
			col = k % lineL
			if row > 0 {
				if col > 0 {
					updateRolls(adjCount, k-lineL-1, k)
				}
				updateRolls(adjCount, k-lineL, k)
				if col < lineL-1 {
					updateRolls(adjCount, k-lineL+1, k)
				}
			}
			if col > 0 {
				updateRolls(adjCount, k-1, k)
			}
		}
	}

	fmt.Println(totalCount)
}

func printRolls(total int, lineL int, adjCount map[int]int) {
	line := ""
	for t := range total {
		if t%lineL == 0 {
			fmt.Println(line)
			line = ""
		}
		if adjCount[t] == 0 {
			line = line + "."
			continue
		}
		if adjCount[t] < 4 {
			line = line + "x"
			continue
		}
		line = line + "@"
	}
}
