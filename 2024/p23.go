package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func P23() {
	fi, err := os.Open("input23.txt")
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
	pcc := make(map[string]map[string]bool)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		pcs := strings.Split(line, "-")
		pc1, pc2 := pcs[0], pcs[1]
		if len(pcc[pc1]) == 0 {
			pcc[pc1] = make(map[string]bool)
		}
		pcc[pc1][pc2] = true
		if len(pcc[pc2]) == 0 {
			pcc[pc2] = make(map[string]bool)
		}
		pcc[pc2][pc1] = true
	}
	existing := make(map[string][]string)
	for k, v := range pcc {
		if k[0] != 't' {
			continue
		}
		for k1 := range v {
			for k2 := range pcc[k1] {
				if pcc[k][k2] {
					if k1[0] == 't' && slices.Contains(existing[k1], k) && slices.Contains(existing[k1], k2) {
						continue
					}
					if k2[0] == 't' && slices.Contains(existing[k2], k) && slices.Contains(existing[k2], k1) {
						continue
					}
					sum++
					fmt.Println(k, k1, k2)
					if len(existing[k]) == 0 {
						existing[k] = make([]string, 0)
					}
					existing[k] = append(existing[k], k1)
					existing[k] = append(existing[k], k2)
				}
			}
		}
	}
	fmt.Println(sum / 2)
}

func P23_1() {
	fi, err := os.Open("input23.txt")
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
	// sum := 0
	pcc := make(map[string]map[string]bool)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		pcs := strings.Split(line, "-")
		pc1, pc2 := pcs[0], pcs[1]
		if len(pcc[pc1]) == 0 {
			pcc[pc1] = make(map[string]bool)
		}
		pcc[pc1][pc2] = true
		if len(pcc[pc2]) == 0 {
			pcc[pc2] = make(map[string]bool)
		}
		pcc[pc2][pc1] = true
	}
	max := 0
	result := []string{}
	for k, v := range pcc {
		if max > len(v) {
			continue
		}
		others := slices.Collect(maps.Keys(v))
		seq := getLongestSeq(pcc, k, slices.Clone(others), false)
		// fmt.Println(seq)
		if len(seq) > max {
			max = len(seq)
			result = seq
		}
	}
	fmt.Println(strings.Join(result, ","))
	slices.Sort(result)
	fmt.Println(strings.Join(result, ","))
}

func getLongestSeq(pcc map[string]map[string]bool, root string, others []string, debug bool) []string {
	if len(others) == 0 {
		return []string{root}
	}
	maxL := 0
	maxA := []string{root}
	if debug {
		fmt.Printf("starting %v, %v\n", root, others)
	}
	validOthers := make([]string, 0)
	for _, other := range others {
		if pcc[root][other] {
			validOthers = append(validOthers, other)
		}
	}
	for i, other := range validOthers {

		if pcc[root][other] {
			start := i + 1
			var subSlice []string
			if start < len(validOthers) {
				subSlice = slices.Clone(validOthers[start:])
			}
			if debug {
				fmt.Printf("verifying %v, %v, %v, %v\n", root, others, other, subSlice)
			}
			cand := append([]string{root}, getLongestSeq(pcc, other, slices.Clone(subSlice), debug)...)
			if debug {
				fmt.Printf("got %v, %v, %v, %v, %v\n", root, others, other, subSlice, cand)
			}
			if len(cand) > maxL {
				maxL = len(cand)
				maxA = cand
			}

		}
	}
	return maxA
}
