package main

import (
	"bufio"
	"strconv"
	"strings"
)

func main() {
	// n := getLineInt()[0]
	// P1()
	// P2_1(n, getLineInt)
	// P3_1()
	// P4(getLine)
	P18_1()
}

func getLine(reader bufio.Reader) (string, error) {
	// reader := bufio.NewReader(source)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.Trim(line, "\r\n ")
	return line, nil
}

func getLineInt(reader bufio.Reader) []int {
	line, _ := getLine(reader)
	splittedLine := strings.Split(line, " ")
	values := make([]int, len(splittedLine))
	for i := 0; i < len(splittedLine); i++ {
		values[i], _ = strconv.Atoi(splittedLine[i])
	}
	return values
}
