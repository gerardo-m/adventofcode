package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	value string
	next  *ListNode
}

func P11() {

	fi, err := os.Open("input11.txt")
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
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		stones := strings.Split(line, " ")
		var head *ListNode
		var curStone *ListNode
		stoneCount := len(stones)
		for i := 0; i < len(stones); i++ {
			newNode := ListNode{stones[i], nil}
			if head == nil {
				head = &newNode
				curStone = &newNode
				continue
			}
			curStone.next = &newNode
			curStone = &newNode
		}
		for i := 0; i < 75; i++ {
			curStone = head
			for ; curStone != nil; curStone = curStone.next {
				if curStone.value == "0" {
					curStone.value = "1"
					continue
				}

				if len(curStone.value)%2 == 0 {
					val1 := curStone.value[:len(curStone.value)/2]
					val2 := strings.TrimPrefix(curStone.value, val1)
					val1 = strings.TrimLeft(val1, "0")
					val2 = strings.TrimLeft(val2, "0")
					if len(val1) == 0 {
						val1 = "0"
					}
					if len(val2) == 0 {
						val2 = "0"
					}
					addedNode := ListNode{val2, curStone.next}
					curStone.value = val1
					curStone.next = &addedNode
					curStone = &addedNode
					stoneCount++
					continue
				}
				intVal, _ := strconv.Atoi(curStone.value)
				intVal = intVal * 2024
				curStone.value = strconv.Itoa(intVal)
			}
			// fmt.Println(i)
			// fmt.Println(stoneCount)
		}
		fmt.Println(stoneCount)
	}
}

func P11_1() {
	fi, err := os.Open("input11.txt")
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
	nStones := make(map[int]int)
	gMap := make(map[int][]int)
	gMap[0] = []int{1}
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		stones := strings.Split(line, " ")
		for _, s := range stones {
			val, _ := strconv.Atoi(s)
			nStones[val]++
		}
		stoneCount := len(stones)
		for i := 0; i < 6; i++ {
			// fmt.Println(gMap)
			fmt.Println(nStones)
			fmt.Println("graph", gMap)
			nStonesClone := maps.Clone(nStones)
			for k := range maps.Keys(nStonesClone) {
				v := nStones[k]
				if v == 0 {
					continue
				}
				ar, ok := gMap[k]
				if !ok || len(ar) == 0 {
					if len(strconv.Itoa(k))%2 == 0 {
						toSplit := strconv.Itoa(k)
						val1 := toSplit[:len(toSplit)/2]
						val2 := strings.TrimPrefix(toSplit, val1)
						val1I, _ := strconv.Atoi(strings.TrimLeft(val1, "0"))
						val2I, _ := strconv.Atoi(strings.TrimLeft(val2, "0"))
						gMap[k] = []int{val1I, val2I}
					} else {
						gMap[k] = []int{k * 2024}
					}
					ar = gMap[k]
				}
				nStones[k] = nStones[k] - v
				stoneCount = stoneCount + (len(ar)-1)*v
				for _, nK := range ar {
					nStones[nK] = nStones[nK] + v
				}
			}
		}
		fmt.Println(stoneCount)
	}

}
