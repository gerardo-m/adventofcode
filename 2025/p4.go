package main

import (
	"fmt"
	"io"
	"os"
)

func P4(fi *os.File) {
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
	// line := ""
	// for t := range ind + 1 {
	// 	if t%lineL == 0 {
	// 		fmt.Println(line)
	// 		line = ""
	// 	}
	// 	if adjCount[t] == 0 {
	// 		line = line + "."
	// 		continue
	// 	}
	// 	if adjCount[t] < 4 {
	// 		line = line + "x"
	// 		continue
	// 	}
	// 	line = line + "@"
	// }
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

func P4_2(fi *os.File) {

}
