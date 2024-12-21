package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func P20() {
	fi, err := os.Open("input20.txt")
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
	lines := make([]string, 0)
	var s, e, lineL int
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lines = append(lines, line)
		lineL = len(line)
		ps := strings.Index(line, "S")
		if ps != -1 {
			s = i*lineL + ps
		}
		pe := strings.Index(line, "E")
		if pe != -1 {
			e = i*lineL + pe
		}
	}
	lineC := len(lines)
	path := make([]int, 0)
	pMap := make(map[int]int)
	path = append(path, s)
	pMap[s] = 0
	curr := s
	fmt.Println(curr, s, e)
	for curr != e {
		for _, adj := range adjacentXY(curr, lineL, lineC) {
			if _, ok := pMap[adj]; ok {
				continue
			}
			v := lines[adj/lineL][adj%lineL]
			if v == '.' || v == 'E' {
				path = append(path, adj)
				pMap[adj] = len(path) - 1
				curr = adj
			}
		}
	}
	// fmt.Println(path, len(path))
	vis := make(map[int]bool)
	cheat := make(map[int]int)
	count := 0
	for i := 0; i < len(path); i++ {
		curr := path[i]
		for _, adj := range adjacentXY(curr, lineL, lineC) {
			if lines[adj/lineL][adj%lineL] != '#' || vis[adj] {
				continue
			}
			// fmt.Println(lines[adj/lineL][adj%lineL], '#')
			// fmt.Printf("looking at %v, %v", adj/lineL, adj%lineL)
			vis[adj] = true
			for _, adj2 := range adjacentXY(adj, lineL, lineC) {
				if adj2 == curr {
					continue
				}
				v, ok := pMap[adj2]
				if !ok {
					continue
				}
				saved := v - i - 2
				if saved >= 100 {
					count++
				}
				cheat[saved]++
				// fmt.Println(saved)
			}
		}
	}

	fmt.Println(cheat)
	fmt.Println(count)
}

func adjacentXY(pos, lineL, lineC int) (adj []int) {
	x := pos % lineL
	y := pos / lineL
	if x > 0 {
		adj = append(adj, y*lineL+x-1)
	}
	if x < lineL-1 {
		adj = append(adj, y*lineL+x+1)
	}
	if y > 0 {
		adj = append(adj, y*lineL+x-lineL)
	}
	if y < lineC-1 {
		adj = append(adj, y*lineL+x+lineL)
	}
	return
}

func P20_1() {
	fi, err := os.Open("input20.txt")
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
	lines := make([]string, 0)
	var s, e, lineL int
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lines = append(lines, line)
		lineL = len(line)
		ps := strings.Index(line, "S")
		if ps != -1 {
			s = i*lineL + ps
		}
		pe := strings.Index(line, "E")
		if pe != -1 {
			e = i*lineL + pe
		}
	}
	lineC := len(lines)
	path := make([]int, 0)
	pMap := make(map[int]int)
	path = append(path, s)
	pMap[s] = 0
	curr := s
	fmt.Println(curr, s, e)
	for curr != e {
		for _, adj := range adjacentXY(curr, lineL, lineC) {
			if _, ok := pMap[adj]; ok {
				continue
			}
			v := lines[adj/lineL][adj%lineL]
			if v == '.' || v == 'E' {
				path = append(path, adj)
				pMap[adj] = len(path) - 1
				curr = adj
			}
		}
	}
	cheat := make(map[int]int)
	count := 0
	for i := 0; i < len(path); i++ {
		pos := path[i]
		x := pos % lineL
		y := pos / lineL
		for dy := -20; dy <= 20; dy++ {
			dif := int(20 - math.Abs(float64(dy)))
			for dx := -dif; dx <= dif; dx++ {
				nx, ny := x+dx, y+dy
				if nx < 0 || nx >= lineL || ny < 0 || ny >= lineC {
					continue
				}
				if lines[ny][nx] == '#' {
					continue
				}
				npos := ny*lineL + nx
				distance := math.Abs(float64(dy)) + math.Abs(float64(dx))
				saved := pMap[npos] - i - int(distance)
				if saved >= 100 {
					cheat[saved]++
					count++
				}
			}
		}
	}
	fmt.Println(cheat)
	fmt.Println(count)
}
