package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func P22() {
	fi, err := os.Open("input22.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		fmt.Println("aa", val)
		for i := 0; i < 10; i++ {
			val = ((val * 64) ^ val) % 16777216
			val = ((val / 32) ^ val) % 16777216
			val = ((val * 2048) ^ val) % 16777216
			fmt.Println(val)
		}
		sum += val
	}
	fmt.Println(sum)
}

func P22_1() {
	fi, err := os.Open("input22.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	sum := 0
	nSeq := make([][]int, 0)
	for i := -9; i < 10; i++ {
		for j := -9; j < 10; j++ {
			if i+j >= -9 && i+j <= +9 {
				nSeq = append(nSeq, []int{i, j})
			}
		}
	}

	fmt.Println(len(nSeq))
	seq := nSeq
	nSeq = make([][]int, 0)
	for i := 0; i < len(seq); i++ {
		cur := seq[i]
		for j := -9; j < 10; j++ {
			tot := cur[0] + cur[1] + j
			if tot >= -9 && tot <= 9 {
				nopt := append(cur, j)
				nSeq = append(nSeq, nopt)
			}
		}
	}
	// fmt.Println(nSeq, len(nSeq))
	seq = nSeq
	nSeq = make([][]int, 0)
	for i := 0; i < len(seq); i++ {
		cur := seq[i]
		for j := -9; j < 10; j++ {
			tot := cur[0] + cur[1] + cur[2] + j
			if tot >= 0 && tot <= 9 {
				nopt := append(slices.Clone(cur), j)
				// fmt.Println(cur)
				// fmt.Println(nopt)
				nSeq = append(nSeq, nopt)
				// fmt.Println(nSeq[len(nSeq)-1])
			}
		}
	}
	// fmt.Println(nSeq, len(nSeq))
	fmt.Println(len(nSeq))
	sortedOpt := make(map[int]map[int][]int)
	for i := 0; i < len(nSeq); i++ {
		if _, ok := sortedOpt[nSeq[i][0]]; !ok {
			sortedOpt[nSeq[i][0]] = make(map[int][]int)
		}
		sortedOpt[nSeq[i][0]][i] = nSeq[i]
	}
	// fmt.Println(sortedOpt[-1])
	// fmt.Println(nSeq, len(nSeq))
	totalSold := make(map[int]int)
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		// fmt.Println("aa", val)
		pPrice := val % 10
		pMap := make(map[int]int)
		sold := make(map[int]int)
		for i := 0; i < 2000; i++ {
			val = ((val * 64) ^ val) % 16777216
			val = ((val / 32) ^ val) % 16777216
			val = ((val * 2048) ^ val) % 16777216
			price := val % 10
			dif := price - pPrice
			for k, v := range pMap {
				if nSeq[k][v] == dif {

					pMap[k]++
					if pMap[k] == 3 {
						// fmt.Println("almost sold", nSeq[k], price)
					}
					if pMap[k] == 4 {
						//SOLD
						sold[k] = price
						// fmt.Println("sold", nSeq[k], price)
						delete(pMap, k)
					}
				} else {
					delete(pMap, k)
				}
			}
			possibleOpt := sortedOpt[dif]
			for k := range possibleOpt {
				if _, ok := sold[k]; ok {
					continue
				}
				if _, ok := pMap[k]; ok {
					continue
				}
				pMap[k] = 1
			}
			pPrice = price
			// fmt.Println(val, dif)
			// fmt.Println(pMap[12473])
			// fmt.Println(len(possibleOpt))
		}
		sum += val
		for k, v := range sold {
			totalSold[k] += v
		}
		fmt.Println(sum)
	}
	max := 0
	maxK := 0
	for k, v := range totalSold {
		// fmt.Println(nSeq[k], v)
		if v > max {
			max = v
			maxK = k
		}
	}
	fmt.Println(nSeq[maxK], max)
	// fmt.Println(totalSold)
}

type CombKey struct {
	v1 int
	v2 int
	v3 int
	v4 int
}

func P22_1o() {
	fi, err := os.Open("input22.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	totalSold := make(map[CombKey]int)
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		pPrice := 0
		var v1, v2, v3, v4 int
		sold := make(map[CombKey]int)
		for i := 0; i < 2000; i++ {
			val = ((val * 64) ^ val) % 16777216
			val = ((val / 32) ^ val) % 16777216
			val = ((val * 2048) ^ val) % 16777216
			price := val % 10
			dif := price - pPrice
			v1 = v2
			v2 = v3
			v3 = v4
			v4 = dif
			pPrice = price
			if i < 3 {
				continue
			}
			key := CombKey{v1, v2, v3, v4}
			if _, ok := sold[key]; ok {
				continue
			}
			sold[key] = price
			pPrice = price
		}
		for k, v := range sold {
			totalSold[k] += v
		}
	}
	max := 0
	var maxK CombKey
	for k, v := range totalSold {
		if v > max {
			max = v
			maxK = k
		}
	}
	fmt.Println(maxK, max)
}
