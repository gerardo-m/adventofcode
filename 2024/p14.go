package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func P14() {
	fi, err := os.Open("input14.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	for k := 1698; k < 100000000; k = k + 103 {
		time.Sleep(time.Millisecond * 200)
		// var v string
		// fmt.Scan(&v)
		// if v == "r" {
		// 	k = k - 2
		// 	continue
		// }
		fi.Seek(0, 0)
		scanner := bufio.NewScanner(fi)
		scanner.Split(bufio.ScanLines)
		// sum := 0
		w := 101
		h := 103
		bath := make([]int, w*h)
		for scanner.Scan() {
			if scanner.Err() == io.EOF {
				break
			}
			line := scanner.Text()
			input := strings.Split(line, " ")
			posi := strings.Split(strings.Split(input[0], "=")[1], ",")
			velo := strings.Split(strings.Split(input[1], "=")[1], ",")
			px, _ := strconv.Atoi(posi[0])
			py, _ := strconv.Atoi(posi[1])
			vx, _ := strconv.Atoi(velo[0])
			vy, _ := strconv.Atoi(velo[1])
			nx := (px + vx*k)
			ny := (py + vy*k)
			if nx < 0 {
				nx = nx * -1
				nx--
				nx = nx % w
				nx = w - nx - 1
			} else {
				nx = nx % w
			}
			if ny < 0 {
				ny = ny * -1
				ny--
				ny = ny % h
				ny = h - ny - 1
			} else {
				ny = ny % h
			}
			bath[ny*w+nx]++
		}
		for i := 0; i < h; i++ {
			// fmt.Println(bath[i*w : (i+1)*w-1])
			for j := i * w; j < (i+1)*w; j++ {
				if bath[j] == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(bath[j])
				}
			}
			fmt.Println()
		}
		fmt.Println(k)
	}

	// q1, q2, q3, q4 := 0, 0, 0, 0
	// for i := 0; i < w/2; i++ {
	// 	for j := 0; j < h/2; j++ {
	// 		q1 = q1 + bath[j*w+i]
	// 	}
	// 	for j := h/2 + 1; j < h; j++ {
	// 		q2 = q2 + bath[j*w+i]
	// 	}
	// }
	// for i := w/2 + 1; i < w; i++ {
	// 	for j := 0; j < h/2; j++ {
	// 		q3 = q3 + bath[j*w+i]
	// 	}
	// 	for j := h/2 + 1; j < h; j++ {
	// 		q4 = q4 + bath[j*w+i]
	// 	}
	// }
	// fmt.Println(q1 * q2 * q3 * q4)
}
