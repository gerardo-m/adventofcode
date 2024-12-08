package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func P7() {
	fi, err := os.Open("input7.txt")
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
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Err() == io.EOF {
			break
		}
		splittedLine := strings.Split(line, ": ")
		y, _ := strconv.Atoi(splittedLine[0])
		ecuation := strings.Split(splittedLine[1], " ")
		val := make([]int, len(ecuation))
		for i, ec := range ecuation {
			val[i], _ = strconv.Atoi(ec)
		}
		solved, _ := resolve(y, val, val[0], 1, "sum")
		if solved {
			sum = sum + y
			continue
		}
		solved, _ = resolve(y, val, val[0], 1, "mul")
		if solved {
			sum = sum + y
		}
	}
	fmt.Println(sum)
}

func resolve(result int, val []int, accu int, curr int, op string) (equal bool, opResult int) {
	if curr == len(val) {
		return equal, accu
	}
	if curr == 0 {
		return false, val[curr]
	}
	var newR int
	if op == "sum" {
		newR = accu + val[curr]
	} else if op == "mul" {
		newR = accu * val[curr]
	} else {
		currVal := val[curr]
		i := 1
		for ; currVal > 0; i = i * 10 {
			currVal = currVal / 10
		}
		newR = accu*i + val[curr]
	}
	if newR > result {
		return false, newR
	}
	if curr == len(val)-1 {
		return newR == result, newR
	}
	solved, finalR := resolve(result, val, newR, curr+1, "sum")
	if solved {
		return true, finalR
	}
	solved, finalR = resolve(result, val, newR, curr+1, "mul")
	if solved {
		return true, finalR
	}
	return resolve(result, val, newR, curr+1, "app")
}

func P7_1() {
	fi, err := os.Open("input7.txt")
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
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Err() == io.EOF {
			break
		}
		splittedLine := strings.Split(line, ": ")
		y, _ := strconv.Atoi(splittedLine[0])
		ecuation := strings.Split(splittedLine[1], " ")
		val := make([]int, len(ecuation))
		for i, ec := range ecuation {
			val[i], _ = strconv.Atoi(ec)
		}
		solved, _ := resolve(y, val, val[0], 1, "sum")
		if solved {
			sum = sum + y
			continue
		}
		solved, _ = resolve(y, val, val[0], 1, "mul")
		if solved {
			sum = sum + y
			continue
		}
		solved, _ = resolve(y, val, val[0], 1, "app")
		if solved {
			sum = sum + y
		}
	}
	fmt.Println(sum)
}
