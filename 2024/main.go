package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := getLineInt()[0]
	// P1()
	P2_1(n, getLineInt)
}

func getLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.Trim(line, "\r\n ")
	return line
}

func getLineInt() []int {
	line := getLine()
	splittedLine := strings.Split(line, " ")
	values := make([]int, len(splittedLine))
	for i := 0; i < len(splittedLine); i++ {
		values[i], _ = strconv.Atoi(splittedLine[i])
	}
	return values
}
