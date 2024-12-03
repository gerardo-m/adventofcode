package main

import (
	"fmt"
	"math"
	"slices"
)

func P2(size int, readLine func() []int) {
	sum := 0
	for i := 0; i < size; i++ {
		values := readLine()
		var prevDiff int
		safe := 1
		for i := 0; i < len(values)-1; i++ {
			diff := values[i] - values[i+1]
			if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
				safe = 0
				break
			}
			distance := math.Abs(float64(diff))
			if distance < 1 || distance > 3 {
				safe = 0
				break
			}
			prevDiff = diff
		}
		sum = sum + safe
	}
	fmt.Println(sum)
}

func P2_1(size int, readLine func() []int) {
	sum := 0
	for i := 0; i < size; i++ {
		values := readLine()
		var prevDiff int
		safe := 1
		for i := 0; i < len(values)-1; i++ {
			diff := values[i] - values[i+1]
			if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
				safe = 0
				break
			}
			distance := math.Abs(float64(diff))
			if distance < 1 || distance > 3 {
				safe = 0
				break
			}
			prevDiff = diff
		}
		if safe == 0 {
			deletedElement := values[0]
			nValues := slices.Delete(values, 0, 1)
			for i := 0; i < len(values); i++ {
				safe = 1
				prevDiff = 0
				for j := 0; j < len(nValues)-1; j++ {
					diff := nValues[j] - nValues[j+1]
					if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
						safe = 0
						break
					}
					distance := math.Abs(float64(diff))
					if distance < 1 || distance > 3 {
						safe = 0
						break
					}
					prevDiff = diff
				}
				if safe == 1 {
					break
				}
				if i >= len(nValues) {
					break
				}
				aux := nValues[i]
				nValues[i] = deletedElement
				deletedElement = aux
			}
		}
		sum = sum + safe
	}
	fmt.Println(sum)
}
