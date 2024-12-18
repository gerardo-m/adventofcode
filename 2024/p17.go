package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func P17() {
	fi, err := os.Open("input17.txt")
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
	var a, b, c int
	// out := make([]int, 0)
	var p string
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		a, _ = strconv.Atoi(strings.Split(line, ": ")[1])
		scanner.Scan()
		line = scanner.Text()
		b, _ = strconv.Atoi(strings.Split(line, ": ")[1])
		scanner.Scan()
		line = scanner.Text()
		c, _ = strconv.Atoi(strings.Split(line, ": ")[1])
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		line = scanner.Text()
		p = strings.Split(line, ": ")[1]
	}
	fmt.Println(a, b, c, p)
	for i := 0; i < len(p); i = i + 4 {
		// fmt.Println(p[i], p[i+2])
		switch p[i] {
		case '0':
			a = a / int(math.Pow(2.0, float64(getComboOp(a, b, c, p[i+2]))))
		case '1':
			b = b ^ int(p[i+2]-48)
		case '2':
			b = getComboOp(a, b, c, p[i+2]) % 8
		case '3':
			if a != 0 {
				i = int(p[i+2]-48) - 4
			}
		case '4':
			b = b ^ c
		case '5':
			v := getComboOp(a, b, c, p[i+2]) % 8
			fmt.Print(v, ",")
		case '6':
			b = a / int(math.Pow(2.0, float64(getComboOp(a, b, c, p[i+2]))))
		case '7':
			c = a / int(math.Pow(2.0, float64(getComboOp(a, b, c, p[i+2]))))
		}
	}
}

func getComboOp(a, b, c int, op byte) int {
	switch op {
	case '4':
		return a
	case '5':
		return b
	case '6':
		return c
	default:
		return int(op - 48)
	}
}

func P17_1() {
	fi, err := os.Open("input17.txt")
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
	var a, b, c int
	var p string
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		a, _ = strconv.Atoi(strings.Split(line, ": ")[1])
		scanner.Scan()
		line = scanner.Text()
		b, _ = strconv.Atoi(strings.Split(line, ": ")[1])
		scanner.Scan()
		line = scanner.Text()
		c, _ = strconv.Atoi(strings.Split(line, ": ")[1])
		scanner.Scan()
		scanner.Text()
		scanner.Scan()
		line = scanner.Text()
		p = strings.Split(line, ": ")[1]
	}
	fmt.Println(p, a)

	for i := 136904917000000; i < 136905019000000; i = i + 1 {
		opt := runProgram(p, i, b, c)
		// fmt.Println(opt, i)
		if opt[0] == '2' && opt[2] == '4' {
			fmt.Println(i)
		}
		if opt == p {
			fmt.Println(i)
			break
		}
	}
}

func runProgram(p string, a, b, c int) string {
	out := make([]string, 0)
	for i := 0; i < len(p); i = i + 4 {
		switch p[i] {
		case '0':
			a = a / int(math.Pow(2.0, float64(getComboOp(a, b, c, p[i+2]))))
		case '1':
			b = b ^ int(p[i+2]-48)
		case '2':
			b = getComboOp(a, b, c, p[i+2]) % 8
		case '3':
			if a != 0 {
				i = int(p[i+2]-48) - 4
			}
		case '4':
			b = b ^ c
		case '5':
			v := getComboOp(a, b, c, p[i+2]) % 8
			out = append(out, strconv.Itoa(v))
		case '6':
			b = a / int(math.Pow(2.0, float64(getComboOp(a, b, c, p[i+2]))))
		case '7':
			c = a / int(math.Pow(2.0, float64(getComboOp(a, b, c, p[i+2]))))
		}
	}
	return strings.Join(out, ",")
}
