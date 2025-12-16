package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node5 struct {
	lim1 int
	lim2 int
	next *Node5
}

func P5(fi *os.File, debug bool) {
	var ranges *Node5
	ranges = nil
	reader := bufio.NewReader(fi)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("%T\n", err)
				panic(err)
			}
		}
		if line == "" {
			break
		}
		lineSp := strings.Split(line, "-")
		lim1, _ := strconv.Atoi(lineSp[0])
		lim2, _ := strconv.Atoi(lineSp[1])
		nRange := Node5{lim1, lim2, ranges}
		ranges = &nRange
	}
	count := 0
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			if err == io.EOF {
				valueToSearch, _ := strconv.Atoi(line)
				if checkFresh(valueToSearch, ranges, debug) {
					count++
				}
				break
			} else {
				panic(err)
			}
		}
		valueToSearch, _ := strconv.Atoi(line)
		if checkFresh(valueToSearch, ranges, debug) {
			count++
		}
	}
	fmt.Println(count)
}

func checkFresh(valueToSearch int, ranges *Node5, debug bool) bool {
	for point := ranges; point != nil; point = point.next {
		if valueToSearch >= point.lim1 && valueToSearch <= point.lim2 {
			if debug {
				fmt.Println("Found", valueToSearch, "in", point.lim1, "-", point.lim2)
			}
			return true
		}
	}
	return false
}

func P5_2(fi *os.File, debug bool) {
	ranges := make(map[int]int)
	reader := bufio.NewReader(fi)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		if line == "" {
			break
		}
		lineSp := strings.Split(line, "-")
		lim1, _ := strconv.Atoi(lineSp[0])
		lim2, _ := strconv.Atoi(lineSp[1])
		for k, v := range ranges {
			if doIntersect(lim1, lim2, k, v) {
				if k < lim1 {
					lim1 = k
				}
				if v > lim2 {
					lim2 = v
				}
			}
		}
		ranges[lim1] = lim2
		for k := range ranges {
			if k > lim1 && k <= lim2 {
				delete(ranges, k)
			}
		}
	}
	count := 0
	for k, v := range ranges {
		if debug {
			fmt.Println(k, "-", v)
		}
		count = count + v - k + 1
	}
	fmt.Println(count)
}

func doIntersect(a1, a2, b1, b2 int) bool {
	return (a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2) || (b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2)
}
