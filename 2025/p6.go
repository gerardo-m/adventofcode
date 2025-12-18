package main

import (
	"fmt"
	"io"
	"os"
)

type Node6 struct {
	value int
	next  *Node6
}

func P6(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	terms := make([]*Node6, 0, 10)
	var head, tail *Node6
	head = nil
	num := 0
	sum := 0
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 {
				if num > 0 {
					newNode := Node6{num, nil}
					if head == nil {
						head = &newNode
						tail = &newNode
					} else {
						tail.next = &newNode
						tail = &newNode
					}
					num = 0
				}
				terms = append(terms, head)
				if debug {
					fmt.Println("head", head.value, "tail", tail.value)
				}
				head = nil
				tail = nil
				continue
			}
			if chunk[i] == ' ' {
				if num == 0 {
					continue
				}
				newNode := Node6{num, nil}
				if head == nil {
					head = &newNode
					tail = &newNode
				} else {
					tail.next = &newNode
					tail = &newNode
				}
				num = 0
				continue
			}
			if chunk[i] == '+' {
				t := len(terms)
				pSum := 0
				for m := range t {
					pSum = pSum + terms[m].value
					terms[m] = terms[m].next
				}
				if debug {
					fmt.Println(pSum)
				}
				sum = sum + pSum
				continue
			}
			if chunk[i] == '*' {
				t := len(terms)
				pMul := 1
				for m := range t {
					pMul = pMul * terms[m].value
					terms[m] = terms[m].next
				}
				if debug {
					fmt.Println(pMul)
				}
				sum = sum + pMul
				continue
			}
			curDig := int(chunk[i] - 48)
			num = num*10 + curDig
		}
	}
	fmt.Println(sum)
}

func P6_2(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	nums := make(map[int]int)
	col := 0
	op := 0
	sum := 0
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 {
				col = 0
				continue
			}
			if chunk[i] == ' ' {
				col++
				continue
			}
			if chunk[i] == '+' {
				v, ok := nums[op]
				pSum := 0
				for ok {
					pSum = pSum + v
					op++
					v, ok = nums[op]
				}
				if debug {
					fmt.Println(pSum)
				}
				sum = sum + pSum
				op++
				continue
			}
			if chunk[i] == '*' {
				v, ok := nums[op]
				pMul := 1
				for ok {
					pMul = pMul * v
					op++
					v, ok = nums[op]
				}
				if debug {
					fmt.Println(pMul)
				}
				sum = sum + pMul
				op++
				continue
			}
			curDig := int(chunk[i] - 48)
			nums[col] = nums[col]*10 + curDig
			col++
		}
	}
	fmt.Println(sum)
}
