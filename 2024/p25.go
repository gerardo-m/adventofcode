package main

import (
	"bufio"
	"fmt"
	"os"
)

func P25() {
	fi, err := os.Open("input25.txt")
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
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		// fmt.Println(line)
		if line[0] == '#' {
			locks = append(locks, getLockKey(scanner))
		} else {
			keys = append(keys, getLockKey(scanner))
		}
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		scanner.Text()
	}
	sum := 0
	for _, lock := range locks {
		for _, key := range keys {
			sum++
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					sum--
					break
				}
			}
		}
	}
	fmt.Println(sum)
}

func getLockKey(scanner *bufio.Scanner) (pins [5]int) {
	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < 5; j++ {
			if line[j] == '#' {
				pins[j]++
			}
		}
	}
	return
}
