package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
)

type Vert struct {
	cost int
	node *GNode
}

type GNode struct {
	w    int
	pos  int
	vert map[int]map[int]*Vert
}

func P16_ob() {
	fi, err := os.Open("input16.txt")
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
	w := make(map[int]bool)
	var s, e, lineL int
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lineL = len(line)
		for j := 0; j < len(line); j++ {
			pos := i*len(line) + j
			if line[j] == '.' {
				continue
			}
			if line[j] == '#' {
				w[pos] = true
				continue
			}
			if line[j] == 'S' {
				s = pos
			}
			if line[j] == 'E' {
				e = pos
			}
		}
	}
	nodes := make(map[int]map[int]*GNode)
	moves := []int{
		1,
		lineL,
		-1,
		-lineL,
	}
	fmt.Println(s, e)
	start := addNode(nodes, w, s, e, nil, 0, 0, moves)
	// fmt.Println(len(nodes))
	navigate(start, 0)
	for n := range maps.Keys(nodes[e]) {
		fmt.Println(nodes[e][n])
	}

}

func navigate(node *GNode, cost int) {
	if node.pos == 28 {
		fmt.Println("END", cost)
	}
	if node.w == 0 {
		node.w = cost
		fmt.Print(node.pos, " ")
	} else {
		if cost <= node.w {
			node.w = cost
			fmt.Print(node.pos, " ")
			// fmt.Println(node.pos, node.w, " ")
		} else {
			fmt.Print(node.w)
			fmt.Println()
			return
		}
	}
	for k := range maps.Keys(node.vert) {
		v := node.vert[k]
		for l := range maps.Keys(v) {
			navigate(v[l].node, node.w+v[l].cost)
		}

	}
}

func addNode(nodes map[int]map[int]*GNode, w map[int]bool, curr, e int, prev *GNode, cost, dir int, moves []int) *GNode {
	// fmt.Println("add", curr)
	if w[curr] {
		return nil
	}
	nodePos, ok := nodes[curr]
	var cNode *GNode
	if ok && nodePos[dir] != nil {
		cNode = nodePos[dir]
	} else {
		cNode = &GNode{0, curr, make(map[int]map[int]*Vert)}
	}
	if prev != nil {
		v := Vert{cost, cNode}
		if _, ok := prev.vert[cNode.pos]; !ok {
			prev.vert[cNode.pos] = make(map[int]*Vert)
		}
		prev.vert[cNode.pos][dir] = &v
	}
	if _, ok := nodes[curr]; !ok {
		nodes[curr] = make(map[int]*GNode)
	}
	nodes[curr][dir] = cNode
	if curr == e {
		return cNode
	}
	// v2:= Vert{cost, cNode}
	// prev.vert = append(prev.vert, &v2)
	cNode.pos = curr
	if _, ok := cNode.vert[curr+moves[dir]][dir]; !ok {
		addNode(nodes, w, curr+moves[dir], e, cNode, 1, dir, moves)
	}
	newDir := (dir + 1) % len(moves)
	if v := cNode.vert[curr]; v[newDir] == nil {
		addNode(nodes, w, curr, e, cNode, 1000, newDir, moves)
	}
	newDir = (dir + len(moves) - 1) % len(moves)
	if v := cNode.vert[curr]; v[newDir] == nil {
		addNode(nodes, w, curr, e, cNode, 1000, newDir, moves)
	}
	return cNode
}

func P16() {
	fi, err := os.Open("input16.txt")
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
	w := make(map[int]bool)
	var s, e, lineL int
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lineL = len(line)
		for j := 0; j < len(line); j++ {
			pos := i*len(line) + j
			if line[j] == '.' {
				continue
			}
			if line[j] == '#' {
				w[pos] = true
				continue
			}
			if line[j] == 'S' {
				s = pos
			}
			if line[j] == 'E' {
				e = pos
			}
		}
	}
	moves := []int{
		1,
		lineL,
		-1,
		-lineL,
	}
	costOf := make(map[int]int)
	curPos := s
	curDir := 0
	calculate(costOf, w, curPos, curDir, 0, e, moves)
	fmt.Println(costOf[e])
}

func calculate(costOf map[int]int, w map[int]bool, curPos, curDir, bCost, e int, moves []int) {
	if w[curPos] {
		return
	}
	cost := costOf[curPos]
	if cost == 0 || bCost < cost {
		cost = bCost
		costOf[curPos] = cost
	} else {
		return
	}
	if curPos == e {
		return
	}
	calculate(costOf, w, curPos+moves[curDir], curDir, cost+1, e, moves)
	newDir := (curDir + 1) % len(moves)
	calculate(costOf, w, curPos+moves[newDir], newDir, cost+1001, e, moves)
	newDir = (curDir + len(moves) - 1) % len(moves)
	calculate(costOf, w, curPos+moves[newDir], newDir, cost+1001, e, moves)
}

func P16_1() {
	fi, err := os.Open("input16.txt")
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
	w := make(map[int]bool)
	var s, e, lineL int
	var i int
	for i = 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lineL = len(line)
		for j := 0; j < len(line); j++ {
			pos := i*len(line) + j
			if line[j] == '.' {
				continue
			}
			if line[j] == '#' {
				w[pos] = true
				continue
			}
			if line[j] == 'S' {
				s = pos
			}
			if line[j] == 'E' {
				e = pos
			}
		}
	}
	moves := []int{
		1,
		lineL,
		-1,
		-lineL,
	}
	costOf := make(map[int]int)
	road := make(map[int]int)
	curPos := s
	curDir := 0
	reachedEnd := calculate1(costOf, road, w, curPos, curDir, 0, e, moves)
	count := 1
	for k := range maps.Keys(road) {
		if road[k] == reachedEnd {
			count++
		}
	}
	// fmt.Println(costOf)
	// fmt.Println(road)
	for index := 0; index < i; index++ {
		for j := 0; j < lineL; j++ {
			pos := index*lineL + j
			if w[pos] {
				fmt.Print("#")
				continue
			}
			if road[pos] == reachedEnd {
				fmt.Print("O")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	// fmt.Println(road[153])
	// fmt.Println(costOf[153])
	// fmt.Println(costOf[168])
	// fmt.Println(costOf[138])
	// fmt.Println(lineL)
	fmt.Println(count)
}

func calculate1(costOf, road map[int]int, w map[int]bool, curPos, curDir, bCost, e int, moves []int) (reachEnd int) {
	if w[curPos] {
		return
	}
	cost := costOf[curPos]
	if cost == 0 || bCost < cost {
		cost = bCost
		costOf[curPos] = cost
	} else if bCost == cost {
		return road[curPos]
	} else if bCost+1 == costOf[curPos+moves[curDir]] {
		return road[curPos]
	} else {
		return
	}
	if curPos == e {
		return cost
	}
	c1 := calculate1(costOf, road, w, curPos+moves[curDir], curDir, cost+1, e, moves)
	newDir := (curDir + 1) % len(moves)
	c2 := calculate1(costOf, road, w, curPos+moves[newDir], newDir, cost+1001, e, moves)
	newDir = (curDir + len(moves) - 1) % len(moves)
	c3 := calculate1(costOf, road, w, curPos+moves[newDir], newDir, cost+1001, e, moves)
	ct := c1
	if ct == 0 || (c2 > 0 && c2 < ct) {
		ct = c2
	}
	if ct == 0 || (c3 > 0 && c3 < ct) {
		ct = c3
	}
	road[curPos] = ct
	return ct
}
