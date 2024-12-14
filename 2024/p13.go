package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func P13() {
	fi, err := os.Open("input13.txt")
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
	sum := 0
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		minCost := 0

		bA := strings.Split(strings.Split(line1, ": ")[1], ", ")
		bB := strings.Split(strings.Split(line2, ": ")[1], ", ")
		pr := strings.Split(strings.Split(line3, ": ")[1], ", ")
		aX, _ := strconv.Atoi(strings.Split(bA[0], "+")[1])
		aY, _ := strconv.Atoi(strings.Split(bA[1], "+")[1])
		bX, _ := strconv.Atoi(strings.Split(bB[0], "+")[1])
		bY, _ := strconv.Atoi(strings.Split(bB[1], "+")[1])
		pX, _ := strconv.Atoi(strings.Split(pr[0], "=")[1])
		pY, _ := strconv.Atoi(strings.Split(pr[1], "=")[1])
		// fmt.Println(aX, aY, bX, bY, pX, pY)

		// dx1 := float64(pX - aX)
		// dy1 := float64(pY - aY)
		// disa := math.Sqrt(math.Pow(dx1, 2) + math.Pow(dy1, 2))
		// dx2 := float64(pX - bX*3)
		// dy2 := float64(pY - bY*3)
		// disb := math.Sqrt(math.Pow(dx2, 2) + math.Pow(dy2, 2))

		// if disa < disb {
		tx := pX / aX
		ty := pY / aY
		var mint int
		if tx < ty {
			mint = tx
		} else {
			mint = ty
		}
		if mint > 100 {
			mint = 100
		}
		dB := 1
		for ; mint > 0 && dB < 100; mint-- {
			totx := aX*mint + dB*bX
			toty := aY*mint + dB*bY
			bB := dB
			for ; totx < pX && toty < pY && dB < 100; dB++ {
				totx = aX*mint + dB*bX
				toty = aY*mint + dB*bY
				bB = dB
			}

			if totx == pX && toty == pY {
				minCost = mint*3 + bB
				break
			}
			dB--
		}
		// } else {
		tx = pX / bX
		ty = pY / bY
		// var mint int
		if tx < ty {
			mint = tx
		} else {
			mint = ty
		}
		if mint > 100 {
			mint = 100
		}
		dA := 1
		for ; mint > 0 && dA < 100; mint-- {
			totx := bX*mint + dA*aX
			toty := bY*mint + dA*aY
			bA := dA
			for ; totx < pX && toty < pY && dA < 100; dA++ {
				totx = bX*mint + dA*aX
				toty = bY*mint + dA*aY
				bA = dA
			}

			if totx == pX && toty == pY {
				minCost2 := mint + bA*3
				if minCost > 0 && minCost2 < minCost {
					minCost = minCost2
				}
				break
			}
			dA--
		}

		// }
		// fmt.Println(minCost)
		sum = sum + minCost

		if !scanner.Scan() {
			break
		}
	}
	fmt.Println(sum)
}

func P13_1() {
	fi, err := os.Open("input13.txt")
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
	sum := 0
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		minCost := 0

		bA := strings.Split(strings.Split(line1, ": ")[1], ", ")
		bB := strings.Split(strings.Split(line2, ": ")[1], ", ")
		pr := strings.Split(strings.Split(line3, ": ")[1], ", ")
		aX, _ := strconv.Atoi(strings.Split(bA[0], "+")[1])
		aY, _ := strconv.Atoi(strings.Split(bA[1], "+")[1])
		bX, _ := strconv.Atoi(strings.Split(bB[0], "+")[1])
		bY, _ := strconv.Atoi(strings.Split(bB[1], "+")[1])
		pX, _ := strconv.Atoi(strings.Split(pr[0], "=")[1])
		pY, _ := strconv.Atoi(strings.Split(pr[1], "=")[1])
		pX = pX + 10000000000000
		pY = pY + 10000000000000
		// tx := pX / aX
		// ty := pY / aY

		x1 := -(aX*bY*pX - aX*pY*bX) / (aY*bX - bY*aX)
		x2 := -(bX*aY*pX - bX*pY*aX) / (bY*aX - aY*bX)
		fmt.Println(x1)
		fmt.Println(x2)
		if x1%aX == 0 {
			rem := pX - x1
			if rem%bX == 0 {
				minCost = (x1/aX)*3 + rem/bX
			}
		}
		if x2%bX == 0 {
			rem := pX - x2
			if rem%aX == 0 {
				nMinCost := x2/bX + (rem/aX)*3
				if minCost == 0 || minCost > nMinCost {
					minCost = nMinCost
				}
			}
		}
		// var mint int
		// if tx < ty {
		// 	mint = tx
		// } else {
		// 	mint = ty
		// }
		// // if mint > 100 {
		// // 	mint = 100
		// // }
		// dB := 1
		// for ; mint > 0; mint-- {
		// 	totx := aX*mint + dB*bX
		// 	toty := aY*mint + dB*bY
		// 	bB := dB
		// 	for ; totx < pX && toty < pY; dB++ {
		// 		totx = aX*mint + dB*bX
		// 		toty = aY*mint + dB*bY
		// 		bB = dB
		// 	}

		// 	if totx == pX && toty == pY {
		// 		minCost = mint*3 + bB
		// 		break
		// 	}
		// 	dB--
		// }
		// // } else {
		// tx = pX / bX
		// ty = pY / bY
		// // var mint int
		// if tx < ty {
		// 	mint = tx
		// } else {
		// 	mint = ty
		// }
		// if mint > 100 {
		// 	mint = 100
		// }
		// dA := 1
		// for ; mint > 0; mint-- {
		// 	totx := bX*mint + dA*aX
		// 	toty := bY*mint + dA*aY
		// 	bA := dA
		// 	for ; totx < pX && toty < pY; dA++ {
		// 		totx = bX*mint + dA*aX
		// 		toty = bY*mint + dA*aY
		// 		bA = dA
		// 	}

		// 	if totx == pX && toty == pY {
		// 		minCost2 := mint + bA*3
		// 		if minCost > 0 && minCost2 < minCost {
		// 			minCost = minCost2
		// 		}
		// 		break
		// 	}
		// 	dA--
		// }
		fmt.Println(minCost)
		fmt.Println(math.MaxInt32)
		sum = sum + minCost

		if !scanner.Scan() {
			break
		}
	}
	fmt.Println(sum)
}
