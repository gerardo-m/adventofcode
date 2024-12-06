package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

// type exists struct{
// 	after map
// }

func P5() {
	fi, err := os.Open("input5.txt")
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
	afterRules := make(map[int]map[int]bool)
	beforeRules := make(map[int]map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if len(line) == 0 {
			break
		}
		var val1, val2 int
		fmt.Sscanf(line, "%d|%d", &val1, &val2)
		if afterRules[val1] == nil {
			afterRules[val1] = make(map[int]bool)
		}
		if beforeRules[val2] == nil {
			beforeRules[val2] = make(map[int]bool)
		}
		afterRules[val1][val2] = true
		beforeRules[val2][val1] = true
	}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		prevPages := make(map[int]bool)
		updatePages := strings.Split(line, ",")
		valid := true
		var middleVal int
		valid, middleVal = isValid(updatePages, beforeRules, afterRules, prevPages)
		if valid {
			sum = sum + middleVal
		}

	}
	fmt.Println(sum)
}

func isValid(updatePages []string, beforeRules map[int]map[int]bool, afterRules map[int]map[int]bool, prevPages map[int]bool) (valid bool, middleVal int) {
	valid = true
	middleIndex := (len(updatePages) - 1) / 2
	for index, sPage := range updatePages {
		page, _ := strconv.Atoi(sPage)
		for prev := range maps.Keys(prevPages) {
			if beforeRules[prev][page] || afterRules[page][prev] {
				valid = false
				break
			}
		}
		prevPages[page] = true
		if !valid {
			break
		}
		if index == middleIndex {
			middleVal = page
		}
	}
	return
}

func P5_2() {
	fi, err := os.Open("input5.txt")
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
	afterRules := make(map[int]map[int]bool)
	beforeRules := make(map[int]map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if len(line) == 0 {
			break
		}
		var val1, val2 int
		fmt.Sscanf(line, "%d|%d", &val1, &val2)
		if afterRules[val1] == nil {
			afterRules[val1] = make(map[int]bool)
		}
		if beforeRules[val2] == nil {
			beforeRules[val2] = make(map[int]bool)
		}
		afterRules[val1][val2] = true
		beforeRules[val2][val1] = true
	}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		prevPages := make(map[int]bool)
		updatePages := strings.Split(line, ",")
		valid := true
		middleIndex := (len(updatePages) - 1) / 2
		valid, _ = isValid(updatePages, beforeRules, afterRules, prevPages)
		if !valid {
			ordered := make([]int, 0, len(updatePages))
			for _, sPage := range updatePages {
				page, _ := strconv.Atoi(sPage)
				i := len(ordered) - 1
				for ; i >= 0; i-- {
					if afterRules[ordered[i]][page] {
						break
					}
				}
				// fmt.Printf("%v %v %v\n", len(ordered), i, page)
				ordered = slices.Insert(ordered, i+1, page)
			}
			sum = sum + ordered[middleIndex]
		}

	}
	fmt.Println(sum)
}
