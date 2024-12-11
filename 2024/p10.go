package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Node struct {
	pos   int
	level int
}

type Stack struct {
	items []Node
}

func (s *Stack) Push(data Node) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (data Node) {
	if len(s.items) == 0 {
		return
	}
	data = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

func P10() {
	fi, err := os.Open("input10.txt")
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
	uphill := make(map[int][]int)
	lines := make([]string, 0, 20)
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lines = append(lines, line)
	}
	lineCount := len(lines)
	trailheads := make([]int, 0, 100)
	for i := 0; i < lineCount; i++ {
		lineL := len(lines[i])
		for j := 0; j < lineL; j++ {
			pos := i*lineL + j
			uphill[pos] = make([]int, 0, 4)
			val := lines[i][j] - 48
			if val == 0 {
				trailheads = append(trailheads, pos)
			}
			// fmt.Printf("pos %v, i %v, j %v\n", pos, i, j)
			if j < lineL-1 {
				next := lines[i][j+1] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos+1)
				}
			}
			if i < lineCount-1 {
				next := lines[i+1][j] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos+lineL)
				}
			}
			if j > 0 {
				next := lines[i][j-1] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos-1)
				}
			}
			if i > 0 {
				next := lines[i-1][j] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos-lineL)
				}
			}
		}
	}
	// fmt.Println(uphill)
	score := 0
	for i := 0; i < len(trailheads); i++ {
		visited := make(map[int]bool)
		completed := 0
		stack := Stack{}
		stack.Push(Node{trailheads[i], 0})
		level := 0
		for len(stack.items) > 0 {
			cur := stack.Pop()
			if visited[cur.pos] {
				continue
			}
			visited[cur.pos] = true
			if cur.level == 9 {
				completed++
				continue
			}
			level = cur.level + 1
			for j := 0; j < len(uphill[cur.pos]); j++ {
				stack.Push(Node{uphill[cur.pos][j], level})
			}
		}
		score = score + completed
	}
	fmt.Println(score)
}

func P10_1() {
	fi, err := os.Open("input10.txt")
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
	uphill := make(map[int][]int)
	lines := make([]string, 0, 20)
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lines = append(lines, line)
	}
	lineCount := len(lines)
	trailheads := make([]int, 0, 100)
	for i := 0; i < lineCount; i++ {
		lineL := len(lines[i])
		for j := 0; j < lineL; j++ {
			pos := i*lineL + j
			uphill[pos] = make([]int, 0, 4)
			val := lines[i][j] - 48
			if val == 0 {
				trailheads = append(trailheads, pos)
			}
			if j < lineL-1 {
				next := lines[i][j+1] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos+1)
				}
			}
			if i < lineCount-1 {
				next := lines[i+1][j] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos+lineL)
				}
			}
			if j > 0 {
				next := lines[i][j-1] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos-1)
				}
			}
			if i > 0 {
				next := lines[i-1][j] - 48
				if next-val == 1 {
					uphill[pos] = append(uphill[pos], pos-lineL)
				}
			}
		}
	}
	rating := 0
	for i := 0; i < len(trailheads); i++ {
		ways := 0
		stack := Stack{}
		stack.Push(Node{trailheads[i], 0})
		level := 0
		for len(stack.items) > 0 {
			cur := stack.Pop()
			if cur.level == 9 {
				ways++
				continue
			}
			level = cur.level + 1
			for j := 0; j < len(uphill[cur.pos]); j++ {
				stack.Push(Node{uphill[cur.pos][j], level})
			}
		}
		rating = rating + ways
	}
	fmt.Println(rating)
}
