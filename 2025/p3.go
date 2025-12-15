package main

import (
	"fmt"
	"io"
	"os"
)

func P3(fi *os.File) {
	chunk := make([]byte, 128)
	sum := 0
	n1 := 0
	n2 := 0
	n1Cand := 0
	for {
		n, err := fi.Read(chunk)
		if err != nil {
			if err == io.EOF {
				maxJoltage := n1*10 + n2
				// fmt.Println(maxJoltage)
				sum = sum + maxJoltage
				break
			} else {
				panic(err)
			}
		}
		for i := range n {
			if chunk[i] == 10 { // EOL
				maxJoltage := n1*10 + n2
				// fmt.Println(maxJoltage)
				sum = sum + maxJoltage
				n1 = 0
				n2 = 0
				n1Cand = 0
				continue
			}
			curDig := int(chunk[i] - 48)
			if n1Cand > n1 {
				n1 = n1Cand
				n2 = 0
			}
			if curDig > n2 {
				n2 = curDig
			}
			if curDig > n1 {
				n1Cand = curDig
			}
		}
	}
	fmt.Println(sum)
}

func P3_2(fi *os.File) {

}
