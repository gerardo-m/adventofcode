package main

import (
	"fmt"
	"io"
	"os"
)

func P3() {
	expectedTokens := map[byte][]byte{
		0:   {'m'},
		'm': {'u'},
		'u': {'l'},
		'l': {'('},
		'(': {'0'},
		'0': {'0', ','},
		',': {'1'},
		'1': {'1', ')'},
	}
	fi, err := os.Open("input3.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	sum := 0
	chunk := make([]byte, 128)
	num1 := 0
	num2 := 0
	currentToken := byte(0)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		for i := 0; i < n; i++ {
			found := false
			for _, eToken := range expectedTokens[currentToken] {
				if isToken(chunk[i], eToken) {
					currentToken = eToken
					found = true
					if eToken == '0' {
						num1 = num1*10 + (int(chunk[i] - 48))
					}
					if eToken == '1' {
						num2 = num2*10 + (int(chunk[i] - 48))
					}
					break
				}
			}
			if found && currentToken == ')' {
				sum = sum + num1*num2
				currentToken = byte(0)
				num1 = 0
				num2 = 0
			}
			if !found {
				currentToken = byte(0)
				num1 = 0
				num2 = 0
			}
		}

	}
	fmt.Println(sum)

}

func isToken(t byte, expected byte) bool {
	if expected == '0' || expected == '1' {
		return t > 47 && t < 58
	}
	return t == expected
}
