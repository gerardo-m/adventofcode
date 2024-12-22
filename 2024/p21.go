package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func P21() {
	fi, err := os.Open("input21.txt")
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
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		seqs := getSeq2(line)
		min := 0
		for _, seq := range seqs {
			fmt.Println(seq)
			seqs2 := getRSequence2(seq)

			for _, seq2 := range seqs2 {
				fmt.Println(seq2, len(seq2))
				seqs3 := getRSequence2(seq2)
				for _, seq3 := range seqs3 {
					if min == 0 || len(seq3) < min {
						min = len(seq3)
					}
					fmt.Println(seq3, len(seq3))
				}

			}

			// fmt.Println(seq3, len(seq3))
		}
		numVal := numPart(line)
		sum += min * numVal
		fmt.Println(numVal, min)

	}
	fmt.Println(sum)
}

func numPart(code string) int {
	res := ""
	for i := 0; i < len(code); i++ {
		if code[i] >= 48 && code[i] < 58 {
			res = res + string(code[i])
		}
	}
	rint, _ := strconv.Atoi(res)
	return rint
}

var bPos map[byte]int = map[byte]int{
	'A': 11,
	'0': 10,
	'1': 6,
	'2': 7,
	'3': 8,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 0,
	'8': 1,
	'9': 2,
}

var rPos map[byte]int = map[byte]int{
	'A': 2,
	'^': 1,
	'<': 3,
	'v': 4,
	'>': 5,
}

func getNDis(a, b int) (dx, dy int) {
	dy = a/3 - b/3
	dx = a%3 - b%3
	return
}

func getSeq2(line string) []string {
	var cc byte = 'A'
	result := make([]string, 1, 8)
	result[0] = ""
	for i := 0; i < len(line); i++ {
		nc := line[i]
		dx, dy := getNDis(bPos[cc], bPos[nc])
		ex := make(map[string]bool)
		nResult := make([]string, 0)
		presult := ""
		if dy > 0 {
			presult += strings.Repeat("^", dy)
		}
		if dx < 0 {
			presult += strings.Repeat(">", -dx)
		}
		if dy < 0 {
			presult += strings.Repeat("v", -dy)
		}
		if dx > 0 {
			presult += strings.Repeat("<", dx)
		}
		presult += "A"
		ex[presult] = true
		for _, pr := range result {
			nResult = append(nResult, pr+presult)
		}
		presult = ""
		if dx < 0 {
			presult += strings.Repeat(">", -dx)
		}
		if dy > 0 {
			presult += strings.Repeat("^", dy)
		}
		if dx > 0 {
			presult += strings.Repeat("<", dx)
		}
		if dy < 0 {
			presult += strings.Repeat("v", -dy)
		}
		presult += "A"
		if !ex[presult] {
			ex[presult] = true
			for _, pr := range result {
				nResult = append(nResult, pr+presult)
			}
		}
		if bPos[cc] < 9 || bPos[nc]%3 > 0 {
			presult = ""
			if dx > 0 {
				presult += strings.Repeat("<", dx)
			}
			if dy > 0 {
				presult += strings.Repeat("^", dy)
			}
			if dx < 0 {
				presult += strings.Repeat(">", -dx)
			}
			if dy < 0 {
				presult += strings.Repeat("v", -dy)
			}
			presult += "A"
			if !ex[presult] {
				ex[presult] = true
				for _, pr := range result {
					nResult = append(nResult, pr+presult)
				}
			}
		}
		if bPos[nc] < 9 || bPos[cc]%3 > 0 {
			presult = ""
			if dy < 0 {
				presult += strings.Repeat("v", -dy)
			}
			if dx < 0 {
				presult += strings.Repeat(">", -dx)
			}
			if dy > 0 {
				presult += strings.Repeat("^", dy)
			}
			if dx > 0 {
				presult += strings.Repeat("<", dx)
			}
			presult += "A"
			if !ex[presult] {
				ex[presult] = true
				for _, pr := range result {
					nResult = append(nResult, pr+presult)
				}
			}
		}
		// fmt.Println(nResult)
		result = nResult
		cc = nc
	}
	max := 0
	betterS := ""

	for _, r := range result {
		max1 := strings.Count(r, "A<")
		if r[0] == '<' {
			max1++
		}
		max2 := strings.Count(r, "Av")
		if r[0] == 'v' {
			max2++
		}
		if max == 0 || max1+max2 > max {
			betterS = r
			max = max1 + max2
		}
	}
	actualResult := make([]string, 1)
	actualResult[0] = betterS
	return actualResult
}

func getRSequence2(rLine string) []string {
	var cc byte = 'A'
	result := make([]string, 1, 8)
	result[0] = ""
	for i := 0; i < len(rLine); i++ {
		// fmt.Println(i)
		// fmt.Println(len(result))
		nc := rLine[i]
		dx, dy := getNDis(rPos[cc], rPos[nc])
		ex := make(map[string]bool)
		nResult := make([]string, 0)
		presult := ""
		if rPos[nc]%3 > 0 {
			presult = ""
			if dx > 0 {
				presult += strings.Repeat("<", dx)
			}
			if dy < 0 {
				presult += strings.Repeat("v", -dy)
			}

			if dy > 0 {
				presult += strings.Repeat("^", dy)
			}
			if dx < 0 {
				presult += strings.Repeat(">", -dx)
			}
			presult += "A"
			if !ex[presult] {
				ex[presult] = true
				for _, pr := range result {
					nResult = append(nResult, pr+presult)
				}
			}
			result = nResult
			cc = nc
			// fmt.Println(cc, nc, result)
			continue
			// if presult[0] == '<' {

			// }
		}
		if rPos[cc]%3 > 0 {
			presult = ""
			if dx > 0 {
				presult += strings.Repeat("<", dx)
			}
			if dy < 0 {
				presult += strings.Repeat("v", -dy)
			}
			if dy > 0 {
				presult += strings.Repeat("^", dy)
			}
			if dx < 0 {
				presult += strings.Repeat(">", -dx)
			}

			presult += "A"
			if !ex[presult] {
				ex[presult] = true
				for _, pr := range result {
					nResult = append(nResult, pr+presult)
				}
			}
			result = nResult
			cc = nc
			continue
			// if presult[0] == '^' {

			// }
		}
		presult = ""
		if dy < 0 {
			presult += strings.Repeat("v", -dy)
		}
		if dx > 0 {
			presult += strings.Repeat("<", dx)
		}
		if dy > 0 {
			presult += strings.Repeat("^", dy)
		}
		if dx < 0 {
			presult += strings.Repeat(">", -dx)
		}

		presult += "A"
		if !ex[presult] {
			ex[presult] = true
			for _, pr := range result {
				nResult = append(nResult, pr+presult)
			}
		}
		// presult = ""
		// if dx < 0 {
		// 	presult += strings.Repeat(">", -dx)
		// }
		// if dy < 0 {
		// 	presult += strings.Repeat("v", -dy)
		// }
		// if dx > 0 {
		// 	presult += strings.Repeat("<", dx)
		// }
		// if dy > 0 {
		// 	presult += strings.Repeat("^", dy)
		// }
		// presult += "A"
		// if !ex[presult] {
		// 	ex[presult] = true
		// 	for _, pr := range result {
		// 		nResult = append(nResult, pr+presult)
		// 	}
		// }

		// fmt.Println(nResult)
		result = nResult
		cc = nc
	}
	max := 0
	betterS := ""
	fmt.Println(len(result))
	for _, r := range result {
		max1 := strings.Count(r, "A<")
		if r[0] == '<' {
			max1++
		}
		max2 := strings.Count(r, "Av")
		if r[0] == 'v' {
			max2++
		}
		if max == 0 || max1+max2 > max {
			betterS = r
			max = max1 + max2
		}
	}
	actualResult := make([]string, 1)
	actualResult[0] = betterS
	fmt.Println(actualResult)
	return actualResult
}

func getSeq(line string) []string {
	var cc byte = 'A'
	result1 := ""
	result2 := ""
	result3 := ""
	result4 := ""
	result5 := ""
	result6 := ""
	for i := 0; i < len(line); i++ {
		nc := line[i]
		dx, dy := getNDis(bPos[cc], bPos[nc])
		result11 := ""
		result44 := ""
		if dy > 0 {
			result11 = result11 + strings.Repeat("^", dy)
		}
		if dx < 0 {
			result11 = result11 + strings.Repeat(">", -dx)
		}
		if dy < 0 {
			result11 = result11 + strings.Repeat("v", -dy)
		}
		if dx > 0 {
			result11 = result11 + strings.Repeat("<", dx)
		}
		result11 = result11 + "A"
		result1 = result1 + result11
		if dx < 0 {
			result44 += strings.Repeat(">", -dx)
		}
		if dy > 0 {
			result44 += strings.Repeat("^", dy)
		}
		if dx > 0 {
			result44 += strings.Repeat("<", dx)
		}
		if dy < 0 {
			result44 += strings.Repeat("v", -dy)
		}
		result44 += "A"
		result4 += result44
		if bPos[cc] < 9 || bPos[nc]%3 > 0 {
			if dx > 0 {
				result2 = result2 + strings.Repeat("<", dx)
				result5 += strings.Repeat("<", dx)
			}
			if dy > 0 {
				result2 = result2 + strings.Repeat("^", dy)
				result5 += strings.Repeat("^", dy)
			}
			if dx < 0 {
				result2 = result2 + strings.Repeat(">", -dx)
				result5 += strings.Repeat(">", -dx)
			}
			if dy < 0 {
				result2 = result2 + strings.Repeat("v", -dy)
				result5 += strings.Repeat("v", -dy)
			}
			result2 += "A"
			result5 += "A"
		} else {
			result2 = result2 + result11
			result5 += result44
		}
		if bPos[nc] < 9 || bPos[cc]%3 > 0 {
			if dy < 0 {
				result3 += strings.Repeat("v", -dy)
				result6 += strings.Repeat("v", -dy)
			}
			if dx < 0 {
				result3 += strings.Repeat(">", -dx)
				result6 += strings.Repeat(">", -dx)
			}
			if dy > 0 {
				result3 += strings.Repeat("^", dy)
				result6 += strings.Repeat("^", dy)
			}
			if dx > 0 {
				result3 += strings.Repeat("<", dx)
				result6 += strings.Repeat("<", dx)
			}
			result3 += "A"
			result6 += "A"
		} else {
			result3 += result11
			result6 += result44
		}
		cc = nc
	}
	result := make([]string, 6)
	result[0] = result1
	result[1] = result2
	result[2] = result3
	result[3] = result4
	result[4] = result5
	result[5] = result6
	return result
}

func getRSequence(rline string) []string {
	var cc byte = 'A'
	result1 := ""
	result2 := ""
	result3 := ""
	result4 := ""
	result5 := ""
	result6 := ""
	for i := 0; i < len(rline); i++ {
		nc := rline[i]
		dx, dy := getNDis(rPos[cc], rPos[nc])
		result11 := ""
		result44 := ""
		if dy < 0 {
			result11 += strings.Repeat("v", -dy)
		}
		if dx < 0 {
			result11 += strings.Repeat(">", -dx)
		}
		if dy > 0 {
			result11 += strings.Repeat("^", dy)
		}
		if dx > 0 {
			result11 += strings.Repeat("<", dx)
		}
		result11 += "A"
		result1 += result11
		if dx < 0 {
			result44 += strings.Repeat(">", -dx)
		}
		if dy < 0 {
			result44 += strings.Repeat("v", -dy)
		}
		if dx > 0 {
			result44 += strings.Repeat("<", dx)
		}
		if dy > 0 {
			result44 += strings.Repeat("^", dy)
		}

		result44 += "A"
		result4 += result44
		if rPos[nc] > 3 {
			if dx > 0 {
				result2 += strings.Repeat("<", dx)
				result5 += strings.Repeat("<", dx)
			}
			if dy < 0 {
				result2 += strings.Repeat("v", -dy)
				result5 += strings.Repeat("v", -dy)
			}
			if dx < 0 {
				result2 += strings.Repeat(">", -dx)
				result5 += strings.Repeat(">", -dx)
			}
			if dy > 0 {
				result2 += strings.Repeat("^", dy)
				result5 += strings.Repeat("^", dy)
			}
			result2 += "A"
			result5 += "A"
		} else {
			result2 += result11
			result5 += result44
		}
		if rPos[cc] > 3 {
			if dy > 0 {
				result3 += strings.Repeat("^", dy)
				result6 += strings.Repeat("^", dy)
			}
			if dx < 0 {
				result3 += strings.Repeat(">", -dx)
				result6 += strings.Repeat(">", -dx)
			}
			if dy < 0 {
				result3 += strings.Repeat("v", -dy)
				result6 += strings.Repeat("v", -dy)
			}
			if dx > 0 {
				result3 += strings.Repeat("<", dx)
				result6 += strings.Repeat("<", dx)
			}
			result3 += "A"
			result6 += "A"
		} else {
			result3 += result11
			result6 += result44
		}

		cc = nc
	}
	result := make([]string, 6)
	result[0] = result1
	result[1] = result2
	result[2] = result3
	result[3] = result4
	result[4] = result5
	result[5] = result6
	return result
}

func P21_1() {
	fi, err := os.Open("input21.txt")
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
	for i := 0; scanner.Scan(); i++ {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		seqs := getSeq2(line)
		min := 0
		moveMap := make(map[MovKey]string)
		lenMap := make(map[LenKey]int)
		for _, seq := range seqs {
			fmt.Println(seq)
			nm := getSeqLength(seq, 25, moveMap, lenMap)
			if min == 0 || nm < min {
				min = nm
			}
			// seqsf := getSeq2Lvl(seq, 12)

			// for _, seq2 := range seqsf {
			// 	if min == 0 || len(seq2) < min {
			// 		min = len(seq2)
			// 	}
			// 	// fmt.Println(seq2, len(seq2))
			// }

			// fmt.Println(seq3, len(seq3))
		}
		numVal := numPart(line)
		sum += min * numVal
		fmt.Println(numVal, min)

	}
	fmt.Println(sum)
}

func getSeq2Lvl(line string, lvl int) []string {
	if lvl == 1 {
		return getRSequence2(line)
	}
	res := make([]string, 0)
	aux := getRSequence2(line)
	for _, val := range aux {
		res = append(res, getSeq2Lvl(val, lvl-1)...)
	}
	return res
}

type MovKey struct {
	from byte
	to   byte
}

type LenKey struct {
	seq string
	lvl int
}

func getSeqLength(line string, lvl int, moveMap map[MovKey]string, lenMap map[LenKey]int) int {
	if lvl == 0 {
		return len(line)
	}
	sum := 0
	var cc byte = 'A'
	for i := 0; i < len(line); i++ {
		var ns string
		if a, ok := moveMap[MovKey{cc, line[i]}]; ok {
			ns = a
		} else {
			ns = getRSequence3(cc, line[i])
			moveMap[MovKey{cc, line[i]}] = ns
		}
		var nl int
		if a, ok := lenMap[LenKey{ns, lvl - 1}]; ok {
			nl = a
		} else {
			nl = getSeqLength(ns, lvl-1, moveMap, lenMap)
		}
		sum += nl
		cc = line[i]
	}
	lenMap[LenKey{line, lvl}] = sum
	return sum
}

func getRSequence3(cc, nc byte) (ns string) {
	dx, dy := getNDis(rPos[cc], rPos[nc])
	presult := ""
	if rPos[nc]%3 > 0 {
		if rPos[cc]%3 > 0 {
			if dx > 0 {
				presult += strings.Repeat("<", dx)
			}
			if dy < 0 {
				presult += strings.Repeat("v", -dy)
			}
			if dy > 0 {
				presult += strings.Repeat("^", dy)
			}
			if dx < 0 {
				presult += strings.Repeat(">", -dx)
			}
			presult += "A"
			return presult
		}
		presult = ""
		if dx > 0 {
			presult += strings.Repeat("<", dx)
		}
		if dy < 0 {
			presult += strings.Repeat("v", -dy)
		}
		if dx < 0 {
			presult += strings.Repeat(">", -dx)
		}
		if dy > 0 {
			presult += strings.Repeat("^", dy)
		}

		presult += "A"
		return presult
	}
	presult = ""
	if rPos[cc]%3 > 0 {
		if dy < 0 {
			presult += strings.Repeat("v", -dy)
		}
		if dx > 0 {
			presult += strings.Repeat("<", dx)
		}
		if dy > 0 {
			presult += strings.Repeat("^", dy)
		}
		if dx < 0 {
			presult += strings.Repeat(">", -dx)
		}
		presult += "A"
		return presult
	}
	if dy < 0 {
		presult += strings.Repeat("v", -dy)
	}
	if dx > 0 {
		presult += strings.Repeat("<", dx)
	}
	if dx < 0 {
		presult += strings.Repeat(">", -dx)
	}
	if dy > 0 {
		presult += strings.Repeat("^", dy)
	}
	presult += "A"
	return presult
}
