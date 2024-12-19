package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P18() {
	fi, err := os.Open("input18.txt")
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
	size := 71
	cor := make(map[int]bool)
	for i := 0; i < 1024; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		cor[y*size+x] = true
	}
	vis := make(map[int]bool)
	next := make([]int, 1)
	next[0] = 0
	end := size*size - 1
	for i := 0; ; i++ {
		cur := next
		next = make([]int, 0)
		// fmt.Println(cur)
		// fmt.Println(vis)
		for _, p := range cur {
			for _, adj := range adjacent(p, size) {
				if cor[adj] {
					continue
				}
				if vis[adj] {
					continue
				}
				vis[adj] = true
				next = append(next, adj)
			}
		}
		if vis[end] {
			fmt.Println(i + 1)
			break
		}
	}
}

func adjacent(pos, size int) []int {
	x := pos % size
	y := pos / size
	var adj []int
	if x > 0 {
		adj = append(adj, y*size+x-1)
	}
	if x < size-1 {
		adj = append(adj, y*size+x+1)
	}
	if y > 0 {
		adj = append(adj, y*size+x-size)
	}
	if y < size-1 {
		adj = append(adj, y*size+x+size)
	}
	return adj
}

func P18_1() {
	fi, err := os.Open("input18.txt")
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
	size := 71
	cor := make(map[int]bool)
	for i := 0; i < 1024; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		cor[y*size+x] = true
	}
	for i := 1024; ; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		cor[y*size+x] = true
		l, reachEnd := findPath(cor, size)
		if !reachEnd {
			fmt.Println(l)
			fmt.Printf("%d,%d\n", x, y)
			break
		}
	}

}

func findPath(cor map[int]bool, size int) (length int, reachEnd bool) {
	vis := make(map[int]bool)
	next := make([]int, 1)
	next[0] = 0
	end := size*size - 1
	for i := 0; ; i++ {
		cur := next
		next = make([]int, 0)
		if len(cur) == 0 {
			return i + 1, false
		}
		for _, p := range cur {
			for _, adj := range adjacent(p, size) {
				if cor[adj] {
					continue
				}
				if vis[adj] {
					continue
				}
				vis[adj] = true
				next = append(next, adj)
			}
		}
		if vis[end] {
			return i + 1, true
		}
	}
}
