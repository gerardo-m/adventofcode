package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func P19() {
	fi, err := os.Open("input19.txt")
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
	towels := make(map[string]bool)
	scanner.Scan()
	fLine := strings.Split(scanner.Text(), ", ")
	maxL := 0
	for i := 0; i < len(fLine); i++ {
		towels[fLine[i]] = true
		if len(fLine[i]) > maxL {
			maxL = len(fLine[i])
		}
	}
	scanner.Scan()
	scanner.Text()
	count := 0
	confirmed := make(map[string]int)
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		if possible(towels, confirmed, line, maxL) {
			count++
		}
	}
	fmt.Println(count)
}

func possible(towels map[string]bool, confirmed map[string]int, s string, maxL int) bool {
	// fmt.Println(s)
	l := len(s)
	if l == 0 {
		return true
	}
	if confirmed[s] > 0 {
		return true
	}
	if confirmed[s] == -1 {
		return false
	}
	var end int
	if l > maxL {
		end = maxL
	} else {
		end = l
	}
	for i := end; i >= 0; i-- {
		cur := s[0:i]
		rem := s[i:]
		var isPossible bool
		if towels[cur] {
			isPossible = possible(towels, confirmed, rem, maxL)
			if isPossible {
				confirmed[rem] = 1
				return true
			} else {
				confirmed[rem] = -1
			}
		}
	}
	return false
}

func P19_1() {
	fi, err := os.Open("input19.txt")
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
	towels := make(map[string]bool)
	scanner.Scan()
	fLine := strings.Split(scanner.Text(), ", ")
	maxL := 0
	for i := 0; i < len(fLine); i++ {
		towels[fLine[i]] = true
		if len(fLine[i]) > maxL {
			maxL = len(fLine[i])
		}
	}
	scanner.Scan()
	scanner.Text()
	count := 0
	confirmed := make(map[string]int)
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		count += possible2(towels, confirmed, line, maxL)
	}
	fmt.Println(count)
}

func possible2(towels map[string]bool, confirmed map[string]int, s string, maxL int) int {
	// fmt.Println(s)
	l := len(s)
	if l == 0 {
		return 1
	}
	if confirmed[s] > 0 {
		return confirmed[s]
	}
	if confirmed[s] == -1 {
		return -1
	}
	var end int
	if l > maxL {
		end = maxL
	} else {
		end = l
	}
	totalPossibilities := 0
	for i := end; i >= 0; i-- {
		cur := s[0:i]
		rem := s[i:]
		var possibleCom int
		if towels[cur] {
			possibleCom = possible2(towels, confirmed, rem, maxL)
			if possibleCom > 0 {
				totalPossibilities += possibleCom
			}
			if possibleCom <= 0 {
				confirmed[s] = -1
			}
		}
	}
	if totalPossibilities == 0 {
		confirmed[s] = -1
	} else {
		confirmed[s] = totalPossibilities
	}
	return totalPossibilities
}
