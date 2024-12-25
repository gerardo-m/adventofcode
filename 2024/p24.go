package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
	a  string
	b  string
	op string
}

func P24() {
	fi, err := os.Open("input24.txt")
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
	wires := make(map[string]int)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		values := strings.Split(line, ": ")
		wires[values[0]], _ = strconv.Atoi(values[1])
	}
	operations := make(map[string]Operation)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		values := strings.Split(line, " ")
		// wires[values[4]] = execBoolOp(wires[values[0]], wires[values[2]], values[1])
		operations[values[4]] = Operation{values[0], values[2], values[1]}
		// fmt.Println(values)
		// fmt.Printf("%v %v = %v\n", a, b, wires[values[4]])
	}
	for result, _ := range operations {
		solveOperation(result, wires, operations, make(map[string]bool))
	}
	sum := 0
	fmt.Println(wires)
	for k, v := range wires {
		if k[0] != 'z' {
			continue
		}
		bit, _ := strconv.Atoi(k[1:])
		sum += v << bit
	}
	fmt.Println(sum)
}

func solveOperation(result string, wires map[string]int, operations map[string]Operation, vis map[string]bool) int {
	// fmt.Println(wires)
	// fmt.Println(wires, operations[result])
	if val, ex := wires[result]; ex {
		return val
	}
	if vis[result] {
		return -1
	}
	vis[result] = true
	var a, b int
	op := operations[result]
	if val, ex := wires[op.a]; ex {
		a = val
	} else {
		a = solveOperation(op.a, wires, operations, vis)
	}
	if val, ex := wires[op.b]; ex {
		b = val
	} else {
		b = solveOperation(op.b, wires, operations, vis)
	}
	wires[result] = execBoolOp(a, b, op.op)

	return wires[result]
}

func execBoolOp(a, b int, op string) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	}
	return 0
}

func P24_1() {
	fi, err := os.Open("input24.txt")
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
	wires := make(map[string]int)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		values := strings.Split(line, ": ")
		wires[values[0]], _ = strconv.Atoi(values[1])
	}
	operations := make(map[string]Operation)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		values := strings.Split(line, " ")
		operations[values[4]] = Operation{values[0], values[2], values[1]}
	}
	wiresA := maps.Clone(wires)
	for result := range operations {
		solveOperation(result, wires, operations, make(map[string]bool))
	}
	sum := 0
	// fmt.Println(wires)
	x := 0
	y := 0
	for k, v := range wires {
		if k[0] == 'x' {
			bit, _ := strconv.Atoi(k[1:])
			x += v << bit
		}
		if k[0] == 'y' {
			bit, _ := strconv.Atoi(k[1:])
			y += v << bit
		}
	}
	swappable := make(map[string]Operation)
	nonSwappable := make(map[string]Operation)
	expectedResult := x + y
	for k, v := range wires {
		if k[0] != 'z' {
			continue
		}
		bit, _ := strconv.Atoi(k[1:])
		mask := (expectedResult >> bit) & 1
		if mask != v {
			swappable[k] = operations[k]
		} else {
			nonSwappable[k] = operations[k]
		}
		sum += v << bit
	}
	markSwappable(maps.Clone(swappable), swappable, operations, wires)
	markSwappable(maps.Clone(nonSwappable), nonSwappable, operations, wires)
	for k := range nonSwappable {
		delete(swappable, k)
	}
	fmt.Printf("result:   %b\nexpected: %b\n", sum, expectedResult)
	swappable0 := make(map[string]Operation)
	swappable1 := make(map[string]Operation)
	for k, v := range swappable {
		if wires[k] == 0 {
			swappable0[k] = v
		} else {
			swappable1[k] = v
		}
	}
	fmt.Println(len(swappable0), len(swappable1))
	for i := range swappable0 {
		for a := range swappable1 {
			op1 := maps.Clone(operations)
			aux1 := op1[i]
			op1[i] = op1[a]
			op1[a] = aux1
			for j := range swappable0 {
				if j == i {
					continue
				}
				for b := range swappable1 {
					if a == b {
						continue
					}
					op2 := maps.Clone(op1)
					aux2 := op2[j]
					op2[j] = op2[b]
					op2[b] = aux2
					for k := range swappable0 {
						if k == j || k == i {
							continue
						}
						for c := range swappable1 {
							if c == b || c == a {
								continue
							}
							op3 := maps.Clone(op2)
							aux3 := op3[k]
							op3[k] = op3[c]
							op3[c] = aux3
							for l := range swappable0 {
								if l == k || l == j || j == i {
									continue
								}
								for d := range swappable1 {
									if d == c || d == b || d == a {
										continue
									}
									op4 := maps.Clone(op3)
									aux4 := op4[l]
									op4[l] = op4[d]
									op4[d] = aux4
									wiresB := maps.Clone(wiresA)
									// fmt.Println(i, j, k, l, a, b, c, d)
									// fmt.Println(op4)
									skip := false
									for result := range op4 {
										// fmt.Println(result, op4[result])
										invalid := solveOperation(result, wiresB, op4, make(map[string]bool))
										if invalid == -1 {
											skip = true
											break
										}
									}
									if skip {
										continue
									}
									sumA := 0
									for k1, v1 := range wires {
										if k1[0] != 'z' {
											continue
										}
										bit, _ := strconv.Atoi(k1[1:])
										sumA += v1 << bit
									}
									if sumA == expectedResult {
										fmt.Println("solved")
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func markSwappable(localSwap, swappable, operations map[string]Operation, wires map[string]int) {
	nSwap := make(map[string]Operation)
	for _, v := range localSwap {
		// if _, ex := swappable[k]; ex {
		// 	continue
		// }
		// ra, rb := willChange(k, operations, wires)
		if v.a[0] != 'x' && v.a[0] != 'y' {
			swappable[v.a] = operations[v.a]
			nSwap[v.a] = operations[v.a]
		}
		if v.b[0] != 'x' && v.b[0] != 'y' {
			swappable[v.b] = operations[v.b]
			nSwap[v.b] = operations[v.b]
		}
	}
	if len(nSwap) == 0 {
		return
	}
	markSwappable(nSwap, swappable, operations, wires)
}

func willChange(res string, operations map[string]Operation, wires map[string]int) (a, b bool) {
	op := operations[res]
	rab := execBoolOp(wires[op.a]^1, wires[op.b]^1, op.op)
	if rab != wires[res] {
		return true, true
	}
	ra := execBoolOp(wires[op.a]^1, wires[op.b], op.op)
	rb := execBoolOp(wires[op.a], wires[op.b]^1, op.op)
	return ra != wires[res], rb != wires[res]
}

func P24_1o() {
	fi, err := os.Open("input24.txt")
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
	wires := make(map[string]int)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		values := strings.Split(line, ": ")
		wires[values[0]], _ = strconv.Atoi(values[1])
	}
	wiresA := maps.Clone(wires)
	operations := make(map[string]Operation)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		values := strings.Split(line, " ")
		operations[values[4]] = Operation{values[0], values[2], values[1]}
	}
	for result, _ := range operations {
		solveOperation(result, wires, operations, make(map[string]bool))
	}
	x := 0
	y := 0
	for k, v := range wires {
		if k[0] == 'x' {
			bit, _ := strconv.Atoi(k[1:])
			x += v << bit
		}
		if k[0] == 'y' {
			bit, _ := strconv.Atoi(k[1:])
			y += v << bit
		}
	}
	expectedResult := x + y
	sum := 0
	notSwappable := make(map[string]bool)
	gates := make(map[string]bool)
	for {
		sum = 0
		swappable := make(map[string]Operation)
		opTreeLens := make(map[int]int)
		andCount := make(map[int]int)
		orCount := make(map[int]int)
		xorCount := make(map[int]int)
		wires = maps.Clone(wiresA)
		for result := range operations {
			solveOperation(result, wires, operations, make(map[string]bool))
		}
		for i := 0; i < 46; i++ {
			k := fmt.Sprintf("z%02d", i)
			v := wires[k]
			mask := (expectedResult >> i) & 1
			opTree := getOpTree(k, operations)
			opTreeLens[i] = len(opTree)
			andCount[i] = strings.Count(opTree, "AND")
			xorCount[i] = strings.Count(opTree, "XOR")
			orCount[i] = strings.Count(opTree, "OR") - xorCount[i]
			fmt.Println(k, andCount[i], orCount[i], xorCount[i])
			if mask != v && !notSwappable[k] {
				swappable[k] = operations[k]
			}
			sum += v << i
		}
		markSwappable(maps.Clone(swappable), swappable, operations, wires)
		finished := true
		for i := 0; i < 46; i++ {
			k := fmt.Sprintf("z%02d", i)
			foundI := false
			if operations[k].op == "AND" {
				for sw := range swappable {
					if operations[sw].op == "XOR" && expectedLength(i) == len(getOpTree(sw, operations)) && operations[sw].a[0] != 'x' && operations[sw].a[0] != 'y' { // && operations[sw].a[1:] == operations[k].a[1:] {
						fmt.Println("found z AND", operations[k], swappable[sw])
						gates[k] = true
						gates[sw] = true
						aux := operations[sw]
						operations[sw] = operations[k]
						operations[k] = aux
						notSwappable[sw] = true
						notSwappable[k] = true
						foundI = true
						break
					}
				}
			}
			if foundI {
				finished = false
				break
			}
		}
		if !finished {
			continue
		}
		for i := 0; i < 46; i++ {
			k := fmt.Sprintf("z%02d", i)
			foundI := false
			if operations[k].op == "OR" {
				for sw := range swappable {
					if operations[sw].op == "XOR" && expectedLength(i) == len(getOpTree(sw, operations)) && operations[sw].a[0] != 'x' && operations[sw].a[0] != 'y' { // && operations[sw].a[1:] == operations[k].a[1:] {
						fmt.Println("found z OR", operations[k], swappable[sw])
						gates[k] = true
						gates[sw] = true
						aux := operations[sw]
						operations[sw] = operations[k]
						operations[k] = aux
						notSwappable[sw] = true
						notSwappable[k] = true
						foundI = true
						break
					}
				}
			}
			if foundI {
				finished = false
				break
			}
		}
		if !finished {
			continue
		}
		for k, v := range swappable {
			foundI := false
			if k[0] != 'z' && v.a[0] != 'x' && v.a[0] != 'y' && v.op == "XOR" {
				foundI = true
				for sw := range swappable {
					if operations[sw].op != "XOR" && len(getOpTree(sw, operations)) == len(getOpTree(k, operations)) {
						fmt.Println("found not x,y XOR", operations[k], swappable[sw])
						gates[k] = true
						gates[sw] = true
						aux := operations[sw]
						operations[sw] = operations[k]
						operations[k] = aux
						break
					}
				}
			}
			if foundI {
				finished = false
				break
			}
		}
		if !finished {
			continue
		}

		for i := 1; i < 45; i++ {
			k := fmt.Sprintf("z%02d", i)
			foundI := false
			if andCount[i] > expectedAndCount(i) {
				k2 := fmt.Sprintf("z%02d", i+1)
				k3 := fmt.Sprintf("z%02d", i-1)
				depTree1, depTree2, depTree3 := make(map[string]bool), make(map[string]bool), make(map[string]bool)
				getDepTree(k, operations, depTree1)
				getDepTree(k2, operations, depTree2)
				getDepTree(k3, operations, depTree3)
				for dt1 := range depTree1 {
					if operations[dt1].op != "AND" {
						continue
					}

					if depTree3[dt1] {
						continue
					}
					fmt.Println("found AND extra1", operations[dt1])
					// if dt1[1:] < k2 {
					// 	continue
					// }
					if depTree2[dt1] {
						for dt2 := range depTree2 {
							fmt.Println("found AND extra2", operations[dt2])
							if operations[dt2].op != "XOR" {
								continue
							}

							if !depTree1[dt2] {
								if len(getOpTree(dt1, operations)) == len(getOpTree(dt2, operations)) {
									foundI = true
									fmt.Println("found AND extra", operations[dt1], swappable[dt2])
									gates[dt1] = true
									gates[dt2] = true
									aux := operations[dt1]
									operations[dt1] = operations[dt2]
									operations[dt2] = aux
									break
								}
							}
						}
						if foundI {
							break
						}
					}
				}
			}
			if foundI {
				finished = false
				break
			}
		}
		if !finished {
			continue
		}
		for i := 0; i < 46; i++ {
			// expectedAnd := i*2 - 1
			// if i == 0 {
			// 	expectedAnd = 0
			// }
			if opTreeLens[i] == expectedLength(i) {
				continue
			}
			fmt.Println("found inconsistency", i, opTreeLens[i], expectedLength(i))
			finished = false
			for sw := range swappable {
				// fmt.Println(len(getOpTree(sw, operations)))
				k := fmt.Sprintf("z%02d", i)
				if len(getOpTree(sw, operations)) == expectedLength(i) {
					fmt.Println("found", operations[k], swappable[sw])
					aux := operations[sw]
					operations[sw] = operations[k]
					operations[k] = aux
					break
				}
			}
			break
		}
		if finished {
			break
		}
	}
	fmt.Printf("result:   %b\nexpected: %b\n", sum, expectedResult)
	fmt.Println(sum, expectedResult)
	gatesA := slices.Collect(maps.Keys(gates))
	slices.Sort(gatesA)
	fmt.Println(strings.Join(gatesA, ","))
}

func getOpTree(result string, operations map[string]Operation) string {
	if result[0] == 'x' || result[0] == 'y' {
		return result
	}
	op := operations[result]
	return fmt.Sprintf("%v%v%v)", getOpTree(op.a, operations), op.op, getOpTree(op.b, operations))
}

func expectedLength(i int) int {
	if i == 0 {
		return 10
	}
	if i == 1 {
		return 24
	}
	if i == 45 {
		return 10 + (27 * (i - 1))
	}
	return 24 + (27 * (i - 1))
}

func expectedAndCount(i int) int {
	if i == 0 {
		return 0
	}
	return 2*i - 1
}

func getDepTree(res string, operations map[string]Operation, depTree map[string]bool) {
	if res[0] == 'x' || res[0] == 'y' {
		return
	}
	depTree[operations[res].a] = true
	depTree[operations[res].b] = true
	getDepTree(operations[res].a, operations, depTree)
	getDepTree(operations[res].b, operations, depTree)
}
