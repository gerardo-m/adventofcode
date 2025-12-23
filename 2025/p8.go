package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

type Distance8 struct {
	value int
	p1    int
	p2    int
}

func P8(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	howManyDistances := 1000
	topDistances := make([]Distance8, 0, howManyDistances)
	points := make([][3]int, 0, 1000)
	curNum := 0
	var curPoint [3]int
	c := 0
	pointLen := 0
	circuits := make(map[int][]int)
	belongC := make(map[int]int)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				curPoint[c] = curNum
				topDistances = searchDistance(topDistances, points, curPoint, false)
				points = append(points, curPoint)
				pointLen++
				curPoint = [3]int{0, 0, 0}
				c = 0
				curNum = 0
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == ',' {
				curPoint[c] = curNum
				c++
				curNum = 0
				continue
			}
			if chunk[i] == 10 {
				curPoint[c] = curNum
				topDistances = searchDistance(topDistances, points, curPoint, false)
				points = append(points, curPoint)
				pointLen++
				curPoint = [3]int{0, 0, 0}
				c = 0
				curNum = 0
				continue
			}
			curDigit := int(chunk[i] - 48)
			curNum = curNum*10 + curDigit
		}
	}
	currentCircuit := 1
	for _, distance := range topDistances {
		if debug {
			fmt.Println(circuits)
			fmt.Println(distance)
			// fmt.Println(belongC)
		}
		a := belongC[distance.p1]
		b := belongC[distance.p2]
		if a == 0 && b == 0 {
			belongC[distance.p1] = currentCircuit
			belongC[distance.p2] = currentCircuit
			pile := make([]int, 2, 1000)
			pile[0] = distance.p1
			pile[1] = distance.p2
			circuits[currentCircuit] = pile
			currentCircuit++
			continue
		}
		if a == 0 {
			belongC[distance.p1] = b
			circuits[b] = append(circuits[b], distance.p1)
			continue
		}
		if b == 0 {
			belongC[distance.p2] = a
			circuits[a] = append(circuits[a], distance.p2)
			continue
		}
		if a == b {
			continue
		}
		circuits[a] = slices.Concat(circuits[a], circuits[b])
		for _, e := range circuits[b] {
			belongC[e] = a
		}
		delete(circuits, b)
	}
	top3 := [3]int{0, 0, 0}
	for _, v := range circuits {
		size := len(v)
		for index := range 3 {
			if size < top3[index] {
				break
			}
			if index > 0 {
				top3[index-1] = top3[index]
			}
			top3[index] = size
		}
	}
	if debug {
		fmt.Println(topDistances)
		fmt.Println(circuits)
		fmt.Println(top3)
	}
	fmt.Println(top3[0] * top3[1] * top3[2])
}

func searchDistance(topDistances []Distance8, points [][3]int, curPoint [3]int, infinite bool) []Distance8 {
	n := len(points)
	for i, point := range points {
		p1 := (point[0] - curPoint[0]) * (point[0] - curPoint[0])
		p2 := (point[1] - curPoint[1]) * (point[1] - curPoint[1])
		p3 := (point[2] - curPoint[2]) * (point[2] - curPoint[2])
		distance := Distance8{p1 + p2 + p3, i, n}
		pos, _ := slices.BinarySearchFunc(topDistances, distance, compareDistances)
		if infinite || len(topDistances) < cap(topDistances) {
			topDistances = append(topDistances, distance)
		}
		n2 := len(topDistances)
		for j := n2 - 1; j > pos; j-- {
			topDistances[j] = topDistances[j-1]
		}
		if pos == n2 {
			continue
		}
		topDistances[pos] = distance
	}
	return topDistances
}

func compareDistances(d1 Distance8, d2 Distance8) int {
	return d1.value - d2.value
}

func P8_2(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	topDistances := make([]Distance8, 0, 10000)
	points := make([][3]int, 0, 1000)
	curNum := 0
	var curPoint [3]int
	c := 0
	circuits := make(map[int][]int)
	belongC := make(map[int]int)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				curPoint[c] = curNum
				topDistances = searchDistance(topDistances, points, curPoint, true)
				points = append(points, curPoint)
				curPoint = [3]int{0, 0, 0}
				c = 0
				curNum = 0
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == ',' {
				curPoint[c] = curNum
				c++
				curNum = 0
				continue
			}
			if chunk[i] == 10 {
				curPoint[c] = curNum
				topDistances = searchDistance(topDistances, points, curPoint, false)
				points = append(points, curPoint)
				curPoint = [3]int{0, 0, 0}
				c = 0
				curNum = 0
				continue
			}
			curDigit := int(chunk[i] - 48)
			curNum = curNum*10 + curDigit
		}
	}
	currentCircuit := 1
	var lastDistance Distance8
	lastDistance = Distance8{0, 0, 0}
	for _, distance := range topDistances {
		if distance.value < lastDistance.value {
			fmt.Println("FAILED")
			break
		}
		lastDistance = distance
		a := belongC[distance.p1]
		b := belongC[distance.p2]
		if a == 0 && b == 0 {
			belongC[distance.p1] = currentCircuit
			belongC[distance.p2] = currentCircuit
			pile := make([]int, 2, 1000)
			pile[0] = distance.p1
			pile[1] = distance.p2
			circuits[currentCircuit] = pile
			currentCircuit++
			continue
		}
		if a == 0 {
			belongC[distance.p1] = b
			circuits[b] = append(circuits[b], distance.p1)
			if len(circuits) == 1 {
				if len(points) == len(circuits[b]) {
					break
				}
			}
			continue
		}
		if b == 0 {
			belongC[distance.p2] = a
			circuits[a] = append(circuits[a], distance.p2)
			if len(circuits) == 1 {
				if len(points) == len(circuits[a]) {
					break
				}
			}
			continue
		}
		if a == b {
			continue
		}
		circuits[a] = slices.Concat(circuits[a], circuits[b])
		for _, e := range circuits[b] {
			belongC[e] = a
		}
		delete(circuits, b)
		if len(circuits) == 1 {
			if len(points) == len(circuits[a]) {
				break
			}
		}
	}
	if debug {
		fmt.Println(circuits)
	}

	fmt.Println(points[lastDistance.p1][0] * points[lastDistance.p2][0])
}
