package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func P6(obstacle int) {
	fi, err := os.Open("input6.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var pos, lineL int
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lineL = len(line)
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		pos = strings.Index(line, "^")
		if pos != -1 {
			pos = lineL*i + pos
			break
		}
	}
	movement := []int{
		-lineL,
		1,
		lineL,
		-1,
	}
	currMov := 0
	currPos := pos
	visitedPos := make(map[int]bool)
	adjustedLineL := lineL + 2
	i := 0
	for ; i < 100000; i++ {
		visitedPos[currPos] = true
		if currPos%lineL == 0 && movement[currMov] == -1 {
			break
		}
		if currPos%lineL == lineL-1 && movement[currMov] == 1 {
			break
		}
		futurePos := currPos + movement[currMov]
		offset := (futurePos/lineL)*adjustedLineL + (futurePos % lineL)
		val := make([]byte, 1)
		_, err := fi.ReadAt(val, int64(offset))
		if err == io.EOF {
			break
		}
		if val[0] == '#' || futurePos == obstacle {
			currMov = (currMov + 1) % len(movement)
			continue
		}
		currPos = futurePos
		if currPos < 0 {
			break
		}
	}
	if i < 100000 {
		fmt.Println("no loop at", obstacle, obstacle/lineL, obstacle%lineL, i)

	}
	// fmt.Println(len(visitedPos))
}

/*

	escape_path < will include direction

	from the end
		put an obstacle and go back 1 position

		travel until
			- land in a visited pos with same direction (loop)
			- left the space ( no loop)
	go back 1 position and repeat
*/

type PathCell struct {
	pos int
	dir int
}

func P6_1() {
	fi, err := os.Open("input6.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	var pos, lineL int
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lineL = len(line)
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		pos = strings.Index(line, "^")
		if pos != -1 {
			pos = lineL*i + pos
			break
		}
	}
	movement := []int{
		-lineL,
		1,
		lineL,
		-1,
	}
	currMov := 0
	currPos := pos
	escapePath := make(map[int]PathCell)
	visitedPos := make(map[int]map[int]bool)
	adjustedLineL := lineL + 2
	i := 0
	for ; ; i++ {
		// visitedPos[currPos] = true
		if _, ok := visitedPos[currPos]; !ok {
			visitedPos[currPos] = make(map[int]bool)
		}
		visitedPos[currPos][currMov] = true
		escapePath[i] = PathCell{currPos, currMov}
		if currPos%lineL == 0 && movement[currMov] == -1 {
			break
		}
		if currPos%lineL == lineL-1 && movement[currMov] == 1 {
			break
		}
		futurePos := currPos + movement[currMov]
		offset := (futurePos/lineL)*adjustedLineL + (futurePos % lineL)
		val := make([]byte, 1)
		_, err := fi.ReadAt(val, int64(offset))
		if err == io.EOF {
			break
		}
		if val[0] == '#' {
			currMov = (currMov + 1) % len(movement)
			continue
		}
		currPos = futurePos
		if currPos < 0 {
			break
		}
	}
	fmt.Println(len(visitedPos))
	// newVisitedPos := maps.Clone(visitedPos)
	// delete(newVisitedPos[escapePath[i].pos], escapePath[i].dir)
	// if len(newVisitedPos[escapePath[i].pos]) == 0 {
	// 	delete(newVisitedPos, escapePath[i].pos)
	// }
	loops := 0
	// printState(escapePath)
	foundObstacles := make(map[int]bool)
	visitedObstacles := make(map[int]bool)
	lastObstacle := -1
	// p := i - 2
	for k := 1; k <= i; k++ {
		// fmt.Println(i)
		obstaclePos := escapePath[k].pos
		// delete(newVisitedPos[escapePath[i-1].pos], escapePath[i-1].dir)
		alternatePath := make(map[int]map[int]bool)
		// fmt.Println(newVisitedPos[escapePath[i].pos])
		// if len(newVisitedPos[escapePath[i-1].pos]) == 0 {
		// 	delete(newVisitedPos, escapePath[i-1].pos)
		// }
		if foundObstacles[obstaclePos] || obstaclePos == pos || lastObstacle == obstaclePos || visitedObstacles[obstaclePos] {
			continue
		}
		lastObstacle = obstaclePos
		visitedObstacles[obstaclePos] = true
		currPos = pos //escapePath[i-1].pos
		currMov = 0   //escapePath[i-1].dir

		altEscPath := make(map[int]PathCell)
		// fmt.Println(altEscPath[0])
		// fmt.Println(newVisitedPos)
		// fmt.Println(escapePath[i].pos, escapePath[i].dir)

		for j := 0; ; j++ {
			if _, ok := alternatePath[currPos]; !ok {
				alternatePath[currPos] = make(map[int]bool)
			}
			// fmt.Println(newVisitedPos[currPos], alternatePath[currPos])
			if currPos%lineL == 0 && movement[currMov] == -1 {
				break
			}
			if currPos%lineL == lineL-1 && movement[currMov] == 1 {
				break
			}
			if alternatePath[currPos][currMov] {
				// fmt.Println(newVisitedPos)
				// fmt.Println("found at", currPos, currMov, obstaclePos)
				if !foundObstacles[obstaclePos] {
					foundObstacles[obstaclePos] = true
					loops++
				}
				break
			}
			alternatePath[currPos][currMov] = true
			altEscPath[j] = PathCell{currPos, currMov}
			// fmt.Println(currPos, currMov)

			futurePos := currPos + movement[currMov]
			// fmt.Printf("move %d\n", futurePos)
			offset := (futurePos/lineL)*adjustedLineL + (futurePos % lineL)
			val := make([]byte, 1)
			_, err := fi.ReadAt(val, int64(offset))
			if err == io.EOF {
				// fmt.Println("breaking")
				break
			}

			// fmt.Println(currPos, currMov)
			// fmt.Println(byte('#'), val[0], obstaclePos, futurePos)
			if val[0] == '#' || futurePos == obstaclePos {
				currMov = (currMov + 1) % len(movement)
				continue
			}

			currPos = futurePos

			if currPos < 0 {
				break
			}
		}
		// println("--------")
		// printState(altEscPath)
	}

	// for obs := range maps.Keys(foundObstacles) {
	// 	P6(obs)
	// }
	fmt.Println(loops, len(foundObstacles))
}

func printState(escapePath map[int]PathCell) {
	const lineL = 10
	var lines [lineL]string
	for i := 0; i < len(lines); i++ {
		lines[i] = ".........."
	}
	for i := 0; i < len(escapePath); i++ {
		y := escapePath[i].pos / lineL
		x := escapePath[i].pos % lineL
		out := []byte(lines[y])
		if escapePath[i].dir == 0 {
			out[x] = '^'
		}
		if escapePath[i].dir == 1 {
			out[x] = '>'
		}
		if escapePath[i].dir == 2 {
			out[x] = 'v'
		}
		if escapePath[i].dir == 3 {
			out[x] = '<'
		}

		lines[y] = string(out)
	}
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

}
