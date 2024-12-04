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

func P3_1() {
	doTokens := map[byte][]byte{
		0:   {'d'},
		'd': {'o'},
		'o': {'('},
		'(': {')'},
	}
	dontTokens := map[byte][]byte{
		0:    {'d'},
		'd':  {'o'},
		'o':  {'n'},
		'n':  {'\''},
		'\'': {'t'},
		't':  {'('},
		'(':  {')'},
	}
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
	enabled := true
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
			dontFound := false
			if enabled {
				found, num1, num2, currentToken = matchToken(expectedTokens, currentToken, chunk[i], num1, num2)
				if found && currentToken == ')' {
					sum = sum + num1*num2
					currentToken = byte(0)
					num1 = 0
					num2 = 0
				}
				if !found {
					dontFound, _, _, currentToken = matchToken(dontTokens, currentToken, chunk[i], 1, 1)
					if dontFound {
						if currentToken == ')' {
							enabled = false
							currentToken = byte(0)
						}
					} else {
						currentToken = byte(0)
					}
					num1 = 0
					num2 = 0
				}
			} else {
				found, _, _, currentToken = matchToken(doTokens, currentToken, chunk[i], 1, 1)
				if found && currentToken == ')' {
					enabled = true
					currentToken = byte(0)
				}
				if !found {
					currentToken = byte(0)
				}
			}
		}

	}
	fmt.Println(sum)
}

func matchToken(expectedTokens map[byte][]byte, currentToken, sourceByte byte, num1, num2 int) (f bool, rnum1, rnum2 int, currToken byte) {
	found := false
	for _, eToken := range expectedTokens[currentToken] {
		if isToken(sourceByte, eToken) {
			currentToken = eToken
			found = true
			if eToken == '0' {
				num1 = num1*10 + (int(sourceByte - 48))
			}
			if eToken == '1' {
				num2 = num2*10 + (int(sourceByte - 48))
			}
			break
		}
	}
	return found, num1, num2, currentToken
}
