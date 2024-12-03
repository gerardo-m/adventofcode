package main

import (
	"fmt"
	"io"
	"os"
)

func P1() {
	fi, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	var list1 []int
	var list2 []int
	for {
		var val1 int
		_, err = fmt.Fscan(fi, &val1)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		var val2 int
		_, err = fmt.Fscan(fi, &val2)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}
	// slices.Sort(list1)
	// slices.Sort(list2)
	sum := 0
	// var list3 []int
	similarity := make(map[int]int)
	for i := 0; i < len(list2); i++ {
		// distance := math.Abs(float64(list1[i] - list2[i]))
		// sum = sum + distance
		curr := similarity[list2[i]]
		similarity[list2[i]] = curr + 1
		// c := slices.Contains(maps.Keys(similarity), list1[i])
		// if {

		// }
	}
	for i := 0; i < len(list1); i++ {
		sum = sum + list1[i]*similarity[list1[i]]
	}
	// intsum := int(sum)
	fmt.Println(sum)
}
