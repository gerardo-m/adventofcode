package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
	"strings"
)

func P15() {
	fi, err := os.Open("input15.txt")
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
	b := make(map[int]bool)
	var r, lineL, lineC, sum int
	for lineC = 0; scanner.Scan(); lineC++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		if len(strings.Trim(line, " ")) == 0 {
			break
		}
		lineL = len(line)
		for j := 0; j < lineL; j++ {
			pos := lineC*lineL + j
			w[pos] = line[j] == '#'
			b[pos] = line[j] == 'O'
			if line[j] == '@' {
				r = pos
			}
		}
	}
	moves := map[byte]int{
		'^': -lineL,
		'v': lineL,
		'>': 1,
		'<': -1,
	}
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			if can, newPos := tryMove(w, b, lineC, lineL, r, moves[line[j]]); can {
				r = newPos
			}
			// for x := 0; x < lineC; x++ {
			// 	for y := 0; y < lineL; y++ {
			// 		pos := x*lineL + y
			// 		if b[pos] {
			// 			fmt.Print("0")
			// 			continue
			// 		}
			// 		if w[pos] {
			// 			fmt.Print("#")
			// 			continue
			// 		}
			// 		if r == pos {
			// 			fmt.Print("@")
			// 			continue
			// 		}
			// 		fmt.Print(".")
			// 	}
			// 	fmt.Println()
			// }
		}
	}

	for i := range maps.Keys(b) {
		if b[i] {
			x := i % lineL
			y := i / lineL
			sum = sum + y*100 + x
		}
	}
	fmt.Println(sum)
}

func tryMove(w, b map[int]bool, lineC, lineL, curr, move int) (canMove bool, newPos int) {
	newPos = curr + move
	if newPos < 0 || newPos >= lineC*lineL {
		return false, 0
	}
	if (move == 1 && newPos%lineL == 0) || (move == -1 && newPos%lineL == lineL-1) {
		return false, 0
	}
	if w[newPos] {
		return false, 0
	}
	canMove = true
	if b[newPos] {
		canMove, _ = tryMove(w, b, lineC, lineL, newPos, move)
	}
	if canMove {
		makeMove(b, curr, newPos)
	}
	return
}

func makeMove(b map[int]bool, from, to int) {
	b[to] = b[from]
	delete(b, from)
}

func P15_1() {
	fi, err := os.Open("input15.txt")
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
	b := make(map[int]bool)
	var r, lineL, lineC, sum int
	for lineC = 0; scanner.Scan(); lineC++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		if len(strings.Trim(line, " ")) == 0 {
			break
		}
		lineL = len(line) * 2
		for j := 0; j < len(line); j++ {
			pos := lineC*lineL + j*2
			w[pos] = line[j] == '#'
			b[pos] = line[j] == 'O'
			if line[j] == '@' {
				r = pos
			}
		}
	}
	moves := map[byte]int{
		'^': -lineL,
		'v': lineL,
		'>': 1,
		'<': -1,
	}
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			if can, newPos := tryMove1(w, b, lineC, lineL, r, moves[line[j]], false); can {
				r = newPos
			}
			// fmt.Printf("%v, %v\n", i*len(line)+j+1, moves[line[j]])
			// for x := 0; x < lineC; x++ {
			// 	for y := 0; y < lineL; y++ {
			// 		pos := x*lineL + y
			// 		if b[pos] {
			// 			fmt.Print("[")
			// 			continue
			// 		}
			// 		if b[pos-1] {
			// 			fmt.Print("]")
			// 			continue
			// 		}
			// 		if w[pos] {
			// 			fmt.Print("#")
			// 			continue
			// 		}
			// 		if w[pos-1] {
			// 			fmt.Print("#")
			// 			continue
			// 		}
			// 		if r == pos {
			// 			fmt.Print("@")
			// 			continue
			// 		}
			// 		fmt.Print(".")
			// 	}
			// 	fmt.Println()
			// }
		}
	}

	for i := range maps.Keys(b) {
		if b[i] {
			x := i % lineL
			y := i / lineL
			sum = sum + y*100 + x
		}
	}
	for x := 0; x < lineC; x++ {
		for y := 0; y < lineL; y++ {
			pos := x*lineL + y
			if b[pos] {
				fmt.Print("[")
				continue
			}
			if b[pos-1] {
				fmt.Print("]")
				continue
			}
			if w[pos] {
				fmt.Print("#")
				continue
			}
			if w[pos-1] {
				fmt.Print("#")
				continue
			}
			if r == pos {
				fmt.Print("@")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println(sum)
}

func tryMove1(w, b map[int]bool, lineC, lineL, curr, move int, isBox bool) (canMove bool, newPos int) {
	newPos = curr + move
	if w[newPos] || w[newPos-1] {
		return false, 0
	}
	canMove = true
	if move == 1 {
		if b[newPos] {
			canMove, _ = tryMove1(w, b, lineC, lineL, newPos+1, move, true)
		}
		if canMove {
			makeMove(b, curr-1, curr)
		}
	} else if move == -1 {
		if b[newPos-1] {
			canMove, _ = tryMove1(w, b, lineC, lineL, newPos-1, move, true)
		}
		if canMove {
			makeMove(b, curr, curr-1)
		}
	} else if move == lineL || move == -lineL {
		if isBox && w[newPos+1] {
			return false, 0
		}
		// cBox := newPos
		if b[newPos] {
			canMove, _ = tryMove1(w, b, lineC, lineL, newPos, move, true)
		} else if b[newPos-1] {
			canMove, _ = tryMove1(w, b, lineC, lineL, newPos-1, move, true)
			// cBox = newPos - 1
			// if canMove {
			// 	makeMove1(b, curr, newPos-1)
			// }
		}

		if canMove && isBox {
			if b[newPos+1] {
				canMove, _ = tryMove1(w, b, lineC, lineL, newPos+1, move, true)
				// cBox = newPos + 1
			}
			// if canMove {
			// 	makeMove1(b, curr, newPos+1)
			// }
		}
		if canMove && !isBox {
			makeMove1(b, curr, move, false)
		}
	}

	return
}

func makeMove1(b map[int]bool, from, move int, isBox bool) {
	to := from + move
	if b[to-1] {
		makeMove1(b, to-1, move, true)
	} else if b[to] {
		makeMove1(b, to, move, true)
	}
	if isBox {
		if b[to+1] {
			makeMove1(b, to+1, move, true)
		}
	}

	b[to] = b[from]
	delete(b, from)
}
