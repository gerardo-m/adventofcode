package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Point9 struct {
	x int
	y int
}

func P9(fi *os.File, debug bool) {
	chunk := make([]byte, 128)
	curNum, x, y := 0, 0, 0
	maxArea := 0
	points := make([]Point9, 0, 1000)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				y = curNum
				curNum = 0
				maxArea = calculateMaxArea(points, x, y, maxArea)
				if debug {
					fmt.Println(x, y, maxArea)
				}
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 {
				y = curNum
				curNum = 0
				maxArea = calculateMaxArea(points, x, y, maxArea)
				if debug {
					fmt.Println(x, y, maxArea)
				}
				points = append(points, Point9{x, y})
				continue
			}
			if chunk[i] == ',' {
				x = curNum
				curNum = 0
				continue
			}
			curDig := int(chunk[i] - 48)
			curNum = curNum*10 + curDig
		}
	}
	fmt.Println(maxArea)
}

func calculateMaxArea(points []Point9, x, y, curMax int) int {
	for _, v := range points {
		w := math.Abs(float64(x-v.x)) + 1
		h := math.Abs(float64(y-v.y)) + 1
		area := int(w * h)
		if area > curMax {
			curMax = area
		}
	}
	return curMax
}

func P9_2(fi *os.File, debug bool) {
	count = 0
	chunk := make([]byte, 128)
	curNum := 0
	var x, y int
	points := make([]Point9, 0, 1000)
	lines := make(map[int]map[int]bool)
	insideArea := make(map[int]map[int]bool)
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				point := Point9{x, y}
				addLine(points[len(points)-1], point, lines)
				//closing the figure
				addLine(point, points[0], lines)
				points = append(points, point)
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 {
				y = curNum
				curNum = 0
				point := Point9{x, y}
				if len(points) > 0 {
					addLine(points[len(points)-1], point, lines)
				}
				points = append(points, point)
				continue
			}
			if chunk[i] == ',' {
				x = curNum
				curNum = 0
				continue
			}
			curDig := int(chunk[i] - 48)
			curNum = curNum*10 + curDig
		}
	}
	n := len(points)
	max := 0
	for i := range n {
		max = checkPoint(points, i, lines, insideArea, max)
		if debug {
			fmt.Println(max, "checked:", count, "at", i)
		}
	}
	fmt.Println(max)
}

func addLine(point1 Point9, point2 Point9, lines map[int]map[int]bool) {
	minX := min(point1.x, point2.x)
	maxX := max(point1.x, point2.x)
	minY := min(point1.y, point2.y)
	maxY := max(point1.y, point2.y)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			_, ok := lines[j]
			if !ok {
				lines[j] = make(map[int]bool)
			}
			lines[j][i] = true
		}
	}
}

func checkPoint(points []Point9, i int, lines, insideArea map[int]map[int]bool, curMax int) int {
	for j := range i {
		minX := min(points[i].x, points[j].x)
		maxX := max(points[i].x, points[j].x)
		minY := min(points[i].y, points[j].y)
		maxY := max(points[i].y, points[j].y)
		w := maxX - minX + 1
		h := maxY - minY + 1
		area := w * h
		if area > curMax {
			// fmt.Println("checking", area, minX, minY, maxX, maxY)
			if checkRectangle(minX, maxX, minY, maxY, points, lines, insideArea) {
				curMax = area
			}
		}
	}
	return curMax
}

func checkRectangle(minX, maxX, minY, maxY int, points []Point9, lines, insideArea map[int]map[int]bool) bool {
	insideLine := true
	for a := minX; a <= maxX; a++ {
		pInside := lines[minY][a]
		if pInside {
			insideLine = true
			continue
		}
		if !insideLine {
			continue
		}
		if isInsideArea(a, minY, points, insideArea) {
			_, ok := insideArea[minY]
			if !ok {
				insideArea[minY] = make(map[int]bool)
			}
			insideArea[minY][a] = true
		} else {
			return false
		}
	}
	for a := minY + 1; a <= maxY; a++ {
		pInside := lines[a][maxX]
		if pInside {
			insideLine = true
			continue
		}
		if !insideLine {
			continue
		}
		if isInsideArea(maxX, a, points, insideArea) {
			_, ok := insideArea[a]
			if !ok {
				insideArea[a] = make(map[int]bool)
			}
			insideArea[a][maxX] = true
		} else {
			return false
		}
	}
	for a := maxX - 1; a >= minX; a-- {
		pInside := lines[maxY][a]
		if pInside {
			insideLine = true
			continue
		}
		if !insideLine {
			continue
		}
		if isInsideArea(a, maxY, points, insideArea) {
			_, ok := insideArea[maxY]
			if !ok {
				insideArea[maxY] = make(map[int]bool)
			}
			insideArea[maxY][a] = true
		} else {
			return false
		}
	}
	for a := maxY - 1; a > minY; a-- {
		pInside := lines[a][minX]
		if pInside {
			insideLine = true
			continue
		}
		if !insideLine {
			continue
		}
		if isInsideArea(minX, a, points, insideArea) {
			_, ok := insideArea[a]
			if !ok {
				insideArea[a] = make(map[int]bool)
			}
			insideArea[a][minX] = true
		} else {
			return false
		}
	}
	return true
}

var count int

func isInsideArea(x, y int, points []Point9, insideArea map[int]map[int]bool) bool {
	opt, ok := insideArea[y]
	if ok {
		if opt[x] {
			return true
		}
	}
	count++
	n := len(points)
	p1 := points[0]
	inside := false
	var p2 Point9
	for i := 1; i <= n; i++ {
		p2 = points[i%n]
		if x == p2.x && y == p2.y {
			return true
		}
		maxX := max(p1.x, p2.x)
		minX := min(p1.x, p2.x)
		if p1.y == p2.y {
			if p1.y == y {
				if x <= maxX && x >= minX {
					return true
				}
			}
			p1 = p2
			continue
		}
		if p1.x < x {
			p1 = p2
			continue
		}
		maxY := max(p1.y, p2.y)
		minY := min(p1.y, p2.y)
		if y <= maxY && y > minY {
			// fmt.Println("flipped", p1, p2, pointToCheck)
			if x == p1.x {
				return true
			}
			inside = !inside
		}
		p1 = p2
	}
	return inside
}

// func P9_2(fi *os.File, debug bool) {
// 	chunk := make([]byte, 128)
// 	curNum, x, y := 0, 0, 0
// 	maxArea := 0
// 	points := make([]Point9, 0, 1000)
// 	// line -> column -> direction(up: true, down: false)
// 	vertLines := make(map[int]map[int]bool)
// 	// line -> column -> isVertice:true
// 	horLines := make(map[int]map[int]bool)
// 	var lastPoint *Point9

// 	for {
// 		n, err := fi.Read(chunk)
// 		if err != nil {
// 			if err == io.EOF {
// 				y = curNum
// 				curNum = 0
// 				if lastPoint != nil {
// 					if lastPoint.x == x {
// 						dir := lastPoint.y > y
// 						minY := min(lastPoint.y, y)
// 						maxY := max(lastPoint.y, y)
// 						for p := minY; p <= maxY; p++ {
// 							_, ok := vertLines[p]
// 							if !ok {
// 								vertLines[p] = make(map[int]bool)
// 							}
// 							vertLines[p][x] = dir
// 						}
// 					} else {
// 						_, ok := horLines[y]
// 						if !ok {
// 							horLines[y] = make(map[int]bool)
// 						}
// 						horLines[y][x] = true
// 						horLines[y][lastPoint.x] = true
// 					}
// 				}
// 				lastPoint = &Point9{x, y}
// 				points = append(points, *lastPoint)
// 				break
// 			} else {
// 				panic(err)
// 			}
// 		}
// 		for i := range n {
// 			if chunk[i] == 10 {
// 				y = curNum
// 				curNum = 0
// 				if lastPoint != nil {
// 					if lastPoint.x == x {
// 						dir := lastPoint.y > y
// 						minY := min(lastPoint.y, y)
// 						maxY := max(lastPoint.y, y)
// 						for p := minY; p <= maxY; p++ {
// 							_, ok := vertLines[p]
// 							if !ok {
// 								vertLines[p] = make(map[int]bool)
// 							}
// 							vertLines[p][x] = dir
// 						}
// 					} else {
// 						_, ok := horLines[y]
// 						if !ok {
// 							horLines[y] = make(map[int]bool)
// 						}
// 						horLines[y][x] = true
// 						horLines[y][lastPoint.x] = true
// 					}
// 				}
// 				lastPoint = &Point9{x, y}
// 				points = append(points, *lastPoint)
// 				continue
// 			}
// 			if chunk[i] == ',' {
// 				x = curNum
// 				curNum = 0
// 				continue
// 			}
// 			curDig := int(chunk[i] - 48)
// 			curNum = curNum*10 + curDig
// 		}
// 	}
// 	x = points[0].x
// 	y = points[0].y
// 	if lastPoint.x == x {
// 		dir := lastPoint.y > y
// 		minY := min(lastPoint.y, y)
// 		maxY := max(lastPoint.y, y)
// 		for p := minY; p <= maxY; p++ {
// 			_, ok := vertLines[p]
// 			if !ok {
// 				vertLines[p] = make(map[int]bool)
// 			}
// 			vertLines[p][x] = dir
// 		}
// 	}
// 	n := len(points)
// 	if debug {
// 		fmt.Println(n, points)
// 	}
// 	for i := range n {
// 		maxArea = calculateMaxValidArea(points, i, points[i].x, points[i].y, maxArea, vertLines, horLines)
// 		if debug {
// 			fmt.Println(maxArea)
// 		}
// 	}
// 	fmt.Println(maxArea)
// }

// func calculateMaxValidArea(points []Point9, i, x, y, curMax int, vertLines, horLines map[int]map[int]bool) int {
// 	for j := range i {
// 		v := points[j]
// 		if v.x != x && v.y != y {
// 			if !verifyHorLine(x, v.x, y, vertLines, horLines) {
// 				continue
// 			}
// 			if !verifyHorLine(x, v.x, v.y, vertLines, horLines) {
// 				continue
// 			}
// 		}
// 		w := math.Abs(float64(x-v.x)) + 1
// 		h := math.Abs(float64(y-v.y)) + 1
// 		area := int(w * h)
// 		fmt.Println("accepted", x, y, v.x, v.y)
// 		if area > curMax {
// 			curMax = area
// 			// fmt.Println(curMax)
// 		}
// 	}
// 	return curMax
// }

// func verifyHorLine(x1, x2, y int, vertLines, horLines map[int]map[int]bool) bool {
// 	minX := min(x1, x2)
// 	vertc1 := 0
// 	vertUp1 := 0
// 	vertDown1 := 0
// 	pureLines1 := 0
// 	vertc2 := 0
// 	vertUp2 := 0
// 	vertDown2 := 0
// 	pureLines2 := 0
// 	isLine1 := false
// 	isLine2 := false
// 	countBetween := 0
// 	row := vertLines[y]
// 	for col, dir := range row {
// 		if col < minX {
// 			continue
// 		}
// 		hl, ok := horLines[y]
// 		if ok {
// 			if hl[col] {
// 				if col >= x1 {
// 					if col == x1 {
// 						isLine1 = true
// 					}
// 					vertc1++
// 					if dir {
// 						vertUp1++
// 					} else {
// 						vertDown1++
// 					}
// 					if col <= x2 {
// 						if dir {
// 							countBetween = countBetween + 1
// 						} else {
// 							countBetween = countBetween - 1
// 						}
// 					}
// 				}
// 				if col >= x2 {
// 					if col == x2 {
// 						isLine2 = true
// 					}
// 					vertc2++
// 					if dir {
// 						vertUp2++
// 					} else {
// 						vertDown2++
// 					}
// 					if col <= x1 {
// 						if dir {
// 							countBetween = countBetween + 1
// 						} else {
// 							countBetween = countBetween - 1
// 						}
// 					}
// 				}
// 			} else {
// 				if col >= x1 {
// 					pureLines1++
// 				}
// 				if col >= x2 {
// 					pureLines2++
// 				}
// 			}
// 		} else {
// 			if col >= x1 {
// 				pureLines1++
// 			}
// 			if col >= x2 {
// 				pureLines2++
// 			}
// 		}
// 	}
// 	if pureLines1 != pureLines2 {
// 		// fmt.Println(x1, x2, y, "dif", pureLines1, pureLines2)
// 		// if x1 == 2 && x2 == 9 {
// 		// 	fmt.Println(vertLines[3], horLines[3])
// 		// }
// 		return false
// 	}
// 	vertDif1 := (vertUp1 - vertDown1) / 2
// 	if vertDif1 < 0 {
// 		vertDif1 = vertDif1 * -1
// 	}
// 	if !isLine1 && ((vertc1%2 == 0) && (vertDif1+pureLines1)%2 == 0) {
// 		// fmt.Println(x1, x2, y, "dif1", vertc1, vertDif1, pureLines1)
// 		return false
// 	}
// 	vertDif2 := (vertUp2 - vertDown2) / 2
// 	if vertDif2 < 0 {
// 		vertDif2 = vertDif2 * -1
// 	}
// 	if !isLine2 && ((vertc2%2 == 0) && (vertDif2+pureLines2)%2 == 0) {
// 		// fmt.Println(x1, x2, y, "dif2", vertc2, vertDif2, pureLines2)
// 		return false
// 	}

// 	if vertDif1+pureLines1 != vertDif2+pureLines2 {
// 		fmt.Println(x1, x2, y, "dif3", vertDif1, pureLines1, vertDif2, pureLines2)
// 		return false
// 	}
// 	return true
// }

// func isInsideArea(points []Point9, pointToCheck Point9) bool {
// 	n := len(points)
// 	// x := float64(pointToCheck.x)
// 	// y := float64(pointToCheck.y)
// 	p1 := points[0]
// 	inside := false
// 	var p2 Point9
// 	// for i := 1; i <= n; i++ {
// 	// 	p2 = points[i%n]
// 	// 	if y >= math.Min(float64(p1.y), float64(p2.y)) {
// 	// 		if y <= math.Max(float64(p1.y), float64(p2.y)) {
// 	// 			// fmt.Println("y passed")
// 	// 			if x <= math.Max(float64(p1.x), float64(p2.x)) {
// 	// 				xIntersection := (y-float64(p1.y))*float64(p2.x-p1.x)/float64(p2.y-p1.y) + float64(p1.x)
// 	// 				if p1.x == p2.x || x <= xIntersection {
// 	// 					fmt.Println("flipped", p1, p2, xIntersection)
// 	// 					inside = !inside
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 	}
// 	// 	fmt.Println(pointToCheck, "inside:", inside)
// 	// 	p1 = p2
// 	// }
// 	for i := 1; i <= n; i++ {
// 		p2 = points[i%n]
// 		if pointToCheck.x == p2.x && pointToCheck.y == p2.y {
// 			return true
// 		}
// 		maxX := max(p1.x, p2.x)
// 		minX := min(p1.x, p2.x)
// 		if p1.y == p2.y {
// 			if p1.y == pointToCheck.y {
// 				if pointToCheck.x <= maxX && pointToCheck.x >= minX {
// 					return true
// 				}
// 			}
// 			p1 = p2
// 			continue
// 		}
// 		if p1.x < pointToCheck.x {
// 			p1 = p2
// 			continue
// 		}
// 		maxY := max(p1.y, p2.y)
// 		minY := min(p1.y, p2.y)
// 		if pointToCheck.y <= maxY && pointToCheck.y > minY {
// 			// fmt.Println("flipped", p1, p2, pointToCheck)
// 			if pointToCheck.x == p1.x {
// 				return true
// 			}
// 			inside = !inside
// 		}
// 		p1 = p2
// 	}
// 	return inside
// }

//4712892210 too high
//4623743592 too high
//4534328125 incorrect
//3564423390 incorrect
//1465767840
