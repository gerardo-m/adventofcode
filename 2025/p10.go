package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var bits = [...]int{
	1, 2, 4, 8, 16, 32, 64, 128, 256, 512,
}

type WiringPair struct {
	value   int
	pathLen int
}

type Queue10 struct {
	data []WiringPair
}

func newQueue() Queue10 {
	return Queue10{make([]WiringPair, 0, 10000)}
}

func (q *Queue10) push(newValue, pathLen int) {
	q.data = append(q.data, WiringPair{newValue, pathLen})
	// fmt.Println("pushed", q.data)
}

func (q *Queue10) pull() WiringPair {
	v := q.data[0]
	q.data = q.data[1:]
	// fmt.Println("pulled", q.data)
	return v
}

func P10(fi *os.File, debug bool) {

	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanWords)
	sum := 0
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lights := scanner.Text()
		lights = lights[1 : len(lights)-1]
		n := len(lights)
		lightValue := 0
		for i := range n {
			if lights[i] == '#' {
				lightValue = lightValue ^ bits[i]
			}
		}
		if debug {
			fmt.Println(lightValue)
		}
		buttonWirings := make([]int, 0, 10)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			nums := scanner.Text()
			if nums[0] == '{' {
				break
			}
			nums = nums[1 : len(nums)-1]
			buttons := strings.Split(nums, ",")
			value := 0
			for _, j := range buttons {
				value = value ^ bits[j[0]-48]
			}
			buttonWirings = append(buttonWirings, value)
		}
		if debug {
			fmt.Println(buttonWirings)
		}
		b := findShortestArrangement(lightValue, buttonWirings)
		if debug {
			fmt.Println(b)
		}
		sum = sum + b
	}
	fmt.Println(sum)
}

func findShortestArrangement(value int, wirings []int) int {
	q := newQueue()
	visited := make(map[int]bool)
	q.push(value, 1)
	visited[value] = true
	for len(q.data) > 0 {
		testingPair := q.pull()
		for _, w := range wirings {
			r := testingPair.value ^ w
			if r == 0 {
				return testingPair.pathLen
			}
			if _, ok := visited[r]; !ok {
				visited[r] = true
				q.push(r, testingPair.pathLen+1)
			}
		}

	}
	return 0
}

func P10_2(fi *os.File, debug bool) {

}
