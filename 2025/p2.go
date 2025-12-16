package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func P2(fi *os.File, debug bool) {
	curNum := 0
	n1 := 0
	n2 := 0
	sum := 0
	chunk := make([]byte, 128)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				n2 = curNum
				findAndSum(&sum, n1, n2)
				break
			} else {
				panic(err)
			}
		}

		for i := range n {
			if chunk[i] == '-' { // - hyphen
				n1 = curNum
				curNum = 0
				continue
			}
			if chunk[i] == ',' { // , comma
				n2 = curNum
				curNum = 0
				findAndSum(&sum, n1, n2)
				continue
			}
			curNum = curNum*10 + int(chunk[i]-48)
		}
	}
	fmt.Println(sum)
}

func findAndSum(sum *int, n1 int, n2 int) {
	// fmt.Println("finding ", n1, n2)
	digits := 0
	close10 := 1
	cur := n1
	for n1 >= close10 {
		digits++
		close10 = close10 * 10
	}
	if digits%2 == 1 {
		cur = close10
		close10 = close10 * 10
		digits++
	}
	for cur < n2 {
		half10 := int(math.Pow10(digits / 2))
		h1 := cur / half10
		searchedValue := h1*half10 + h1
		if searchedValue >= cur && searchedValue <= n2 {
			// fmt.Println("Found: ", searchedValue)
			*sum = *sum + searchedValue
		}
		cur = (h1 + 1) * half10
		if cur >= close10 {
			digits = digits + 2
			cur = close10 * 10
			close10 = close10 * 100
		}
	}
}

func P2_2(fi *os.File, debug bool) {
	curNum := 0
	n1 := 0
	n2 := 0
	sum := 0
	found := make(map[int]bool)
	chunk := make([]byte, 128)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				n2 = curNum
				findAndSum2(&sum, n1, n2, found)
				break
			} else {
				panic(err)
			}
		}

		for i := range n {
			if chunk[i] == '-' { // - hyphen
				n1 = curNum
				curNum = 0
				continue
			}
			if chunk[i] == ',' { // , comma
				n2 = curNum
				curNum = 0
				findAndSum2(&sum, n1, n2, found)
				continue
			}
			curNum = curNum*10 + int(chunk[i]-48)
		}
	}
	fmt.Println(sum)
}

func findAndSum2(sum *int, n1 int, n2 int, found map[int]bool) {
	// fmt.Println("finding ", n1, n2)
	digits := 0
	close10 := 1
	for n1 >= close10 {
		digits++
		close10 = close10 * 10
	}
	digitsTop := digits
	for n2 >= close10 {
		digitsTop++
		close10 = close10 * 10
	}
	startingNum := n1
	for digitCount := digits; digitCount <= digitsTop && startingNum <= n2; digitCount++ {
		for testingDigit := 1; testingDigit <= (digitCount / 2); testingDigit++ {
			if digitCount%testingDigit != 0 {
				continue
			}
			ocurForDigit := digitCount / testingDigit
			div10 := int(math.Pow10(testingDigit))
			cur := startingNum
			for cur <= n2 {
				h1 := cur
				for range ocurForDigit - 1 {
					h1 = h1 / div10
				}
				searchedValue := 0
				for range ocurForDigit {
					searchedValue = searchedValue*div10 + h1
				}
				if searchedValue >= cur && searchedValue <= n2 && !found[searchedValue] {
					// fmt.Println("Found: ", searchedValue)
					found[searchedValue] = true
					*sum = *sum + searchedValue
				}
				cur = h1 + 1
				if cur >= div10 {
					break
				}
				for range ocurForDigit - 1 {
					cur = cur * div10
				}
			}
		}
		startingNum = int(math.Pow10(digitCount))
	}

}
