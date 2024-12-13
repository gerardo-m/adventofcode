package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
)

type Region struct {
	pos  int
	val  byte
	area int
	per  int
}

func P12() {
	fi, err := os.Open("input12.txt")
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
	lines := make([]string, 0, 1000)
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lines = append(lines, line)
	}
	lineC := len(lines)
	lineL := len(lines[0])
	rMap := make([]int, lineC*lineL)
	for i := 0; i < len(rMap); i++ {
		rMap[i] = -1
	}
	var regions []Region
	for i := 0; i < lineC; i++ {
		for j := 0; j < lineL; j++ {
			pos := i*lineL + j
			// fmt.Printf("%v, %v and it value is %v\n", i, j, rMap[pos])
			if rMap[pos] >= 0 {
				continue
			}
			val := lines[i][j]
			st := Stack{}
			reg := Region{pos, val, 0, 0}
			regions = append(regions, reg)
			curRegion := len(regions) - 1
			st.Push(Node{pos, 0})
			for len(st.items) > 0 {
				wPos := st.Pop().pos
				if rMap[wPos] >= 0 {
					continue
				}
				rMap[wPos] = curRegion
				regions[curRegion].area++
				exAdj := 0
				allAdj := Adj(lineL, lineC, wPos/lineL, wPos%lineL)
				for _, adj := range allAdj {
					// fmt.Printf("with curRegion %v, val %v, reg %v, value %v\n", curRegion, val, lines[adj/lineL][adj%lineL], rMap[adj])
					if rMap[adj] == curRegion {
						exAdj++
						continue
					}
					if rMap[adj] >= 0 {
						continue
					}
					if lines[adj/lineL][adj%lineL] == val {
						st.Push(Node{adj, 0})
					}
				}
				regions[curRegion].per = regions[curRegion].per + 4 - (exAdj * 2)
				// fmt.Println(regions)
			}
		}
	}
	price := 0
	for i := 0; i < len(regions); i++ {
		price = price + regions[i].per*regions[i].area
	}
	// fmt.Println(regions)
	fmt.Println(price)
}

func Adj(lineL, lineC, r, c int) []int {
	adj := make([]int, 0, 4)
	if r > 0 {
		adj = append(adj, (r-1)*lineL+c)
	}
	if r < lineC-1 {
		adj = append(adj, (r+1)*lineL+c)
	}
	if c > 0 {
		adj = append(adj, r*lineL+c-1)
	}
	if c < lineL-1 {
		adj = append(adj, r*lineL+c+1)
	}
	return adj
}

type Range struct {
	start int
	end   int
}

func P12_1() {
	fi, err := os.Open("input12.txt")
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
	lines := make([]string, 0, 1000)
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		lines = append(lines, line)
	}
	lineC := len(lines)
	lineL := len(lines[0])
	rMap := make([]int, lineC*lineL)
	for i := 0; i < len(rMap); i++ {
		rMap[i] = -1
	}
	var regions []Region
	for i := 0; i < lineC; i++ {
		for j := 0; j < lineL; j++ {
			pos := i*lineL + j
			if rMap[pos] >= 0 {
				continue
			}
			val := lines[i][j]
			st := Stack{}
			reg := Region{pos, val, 0, 0}
			regions = append(regions, reg)
			curRegion := len(regions) - 1
			st.Push(Node{pos, 0})
			for len(st.items) > 0 {
				wPos := st.Pop().pos
				if rMap[wPos] >= 0 {
					continue
				}
				rMap[wPos] = curRegion
				regions[curRegion].area++
				// exAdj := 0
				allAdj := Adj(lineL, lineC, wPos/lineL, wPos%lineL)
				for _, adj := range allAdj {
					if rMap[adj] == curRegion {
						// exAdj++
						continue
					}
					if rMap[adj] >= 0 {
						continue
					}
					if lines[adj/lineL][adj%lineL] == val {
						st.Push(Node{adj, 0})
					}
				}
				// regions[curRegion].per = regions[curRegion].per + 4 - (exAdj * 2)
			}
		}
	}
	price := 0
	h := make(map[int][][]Range)
	v := make(map[int][][]Range)
	for i := 0; i < lineC; i++ {
		l := make(map[int][]Range)
		var curReg int
		for j := 0; j < lineL; j++ {
			pos := i*lineL + j
			curReg = rMap[pos]

			s := j
			for j < lineL && rMap[i*lineL+j] == curReg {
				j++
			}
			j--
			r := Range{s, j}
			l[curReg] = append(l[curReg], r)
		}
		for key := range maps.Keys(l) {
			h[key] = append(h[key], l[key])
		}
	}
	for j := 0; j < lineL; j++ {
		l := make(map[int][]Range)
		var curReg int
		for i := 0; i < lineC; i++ {
			pos := i*lineL + j
			curReg = rMap[pos]

			s := i
			for i < lineC && rMap[i*lineL+j] == curReg {
				i++
			}
			i--
			r := Range{s, i}
			l[curReg] = append(l[curReg], r)
		}
		for key := range maps.Keys(l) {
			v[key] = append(v[key], l[key])
		}
	}
	for i := 0; i < len(regions); i++ {
		region := regions[i]
		ha := h[i]
		va := v[i]
		hv := 2 * len(ha[0])
		vv := 2 * len(va[0])
		for hi := 1; hi < len(ha); hi++ {
			prev := len(ha[hi-1])
			sm := make(map[int]bool)
			sn := make(map[int]bool)
			for k := 0; k < len(ha[hi-1]); k++ {
				sm[ha[hi-1][k].start] = true
				sn[ha[hi-1][k].end] = true
			}
			for k := 0; k < len(ha[hi]); k++ {
				sm[ha[hi][k].start] = true
				sn[ha[hi][k].end] = true
			}
			dif := len(sm) - prev + len(sn) - prev
			hv = hv + dif

		}
		for vi := 1; vi < len(va); vi++ {
			prev := len(va[vi-1])
			sm := make(map[int]bool)
			sn := make(map[int]bool)
			for k := 0; k < len(va[vi-1]); k++ {
				sm[va[vi-1][k].start] = true
				sn[va[vi-1][k].end] = true
			}
			for k := 0; k < len(va[vi]); k++ {
				sm[va[vi][k].start] = true
				sn[va[vi][k].end] = true
			}
			dif := len(sm) - prev + len(sn) - prev
			vv = vv + dif
		}

		sides := hv + vv
		price = price + region.area*sides
	}
	// mov := []int{
	// 	1,
	// 	lineL,
	// 	-1,
	// 	-lineL,
	// }
	// for i := 0; i < len(regions); i++ {
	// 	curMov := 0
	// 	curPos := regions[i].pos
	// 	sides := 1
	// 	for {
	// 		// fmt.Println(curPos, curMov)
	// 		if curPos == regions[i].pos && curMov == 3 {
	// 			break
	// 		}
	// 		for dp := -1; dp < 3; dp++ {
	// 			if curPos == regions[i].pos && dp == 2 {
	// 				curMov = 3
	// 				if sides == 2 {
	// 					sides = 4
	// 				}
	// 				// sides++
	// 				break
	// 			}
	// 			can, newPos := canMove(lineC, lineL, curPos, mov[(4+curMov+dp)%4], regions[i].val, lines)
	// 			// fmt.Println(dp, can, newPos, (4+curMov+dp)%4, regions[i].val)
	// 			if dp > 0 {
	// 				sides++
	// 			}
	// 			if can {
	// 				curPos = newPos
	// 				curMov = (4 + curMov + dp) % 4
	// 				if dp == -1 {
	// 					sides++
	// 				}
	// 				break
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(price, sides)
	// 	price = price + sides*regions[i].area
	// }
	// fmt.Println(regions)
	fmt.Println(price)
}

func canMove(lineC, lineL, pos, move int, val byte, content []string) (can bool, newPos int) {
	// r := pos / lineL
	c := pos % lineL
	newPos = pos + move
	if newPos < 0 || newPos >= (lineC*lineL) {
		return false, 0
	}
	if (move == 1 && c == lineL-1) || (move == -1 && c == 0) {
		return false, 0
	}

	if content[newPos/lineL][newPos%lineL] == val {
		return true, newPos
	}
	return false, 0
}
