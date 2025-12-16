package main

import (
	"fmt"
	"io"
	"os"
)

func P3(fi *os.File) {
	chunk := make([]byte, 128)
	sum := 0
	n1 := 0
	n2 := 0
	n1Cand := 0
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				maxJoltage := n1*10 + n2
				// fmt.Println(maxJoltage)
				sum = sum + maxJoltage
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 { // EOL
				maxJoltage := n1*10 + n2
				// fmt.Println(maxJoltage)
				sum = sum + maxJoltage
				n1 = 0
				n2 = 0
				n1Cand = 0
				continue
			}
			curDig := int(chunk[i] - 48)
			if n1Cand > n1 {
				n1 = n1Cand
				n2 = 0
			}
			if curDig > n2 {
				n2 = curDig
			}
			if curDig > n1 {
				n1Cand = curDig
			}
		}
	}
	fmt.Println(sum)
}

func P3_2(fi *os.File) {
	chunk := make([]byte, 128)
	sum := 0
	totalDigits := 12
	num := make([]int, totalDigits)
	curPos := 0
	replacedAt := 1
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				maxJoltage := joinJoltage(num)
				// fmt.Println(maxJoltage)
				sum = sum + maxJoltage
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 { // EOL
				maxJoltage := joinJoltage(num)
				// fmt.Println(maxJoltage)
				sum = sum + maxJoltage
				zeroArray(num)
				replacedAt = 1
				curPos = 0
				continue
			}
			curDig := int(chunk[i] - 48)
			if curPos < totalDigits {
				num[curPos] = curDig
				curPos++
				continue
			}
			if replacedAt == 0 {
				replacedAt = 1
			}
			if replacedAt >= totalDigits {
				replacedAt = totalDigits - 1
			}
			j := replacedAt
			for ; j < totalDigits; j++ {
				if num[j] > num[j-1] {
					break
				}
			}
			replacedAt = j - 1
			for ; j < totalDigits; j++ {
				num[j-1] = num[j]
			}
			if replacedAt < totalDigits-1 || curDig > num[totalDigits-1] {
				num[totalDigits-1] = curDig
			}
		}
	}
	fmt.Println(sum)
}

func joinJoltage(nums []int) int {
	r := 0
	for _, num := range nums {
		r = r*10 + num
	}
	return r
}

func zeroArray(nums []int) {
	n := len(nums)
	for i := range n {
		nums[i] = 0
	}
}
