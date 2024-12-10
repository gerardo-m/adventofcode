package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func P9() {
	fi, err := os.Open("input9.txt")
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
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		totalL := len(line)
		i, j := 0, totalL-1
		if j%2 != 0 {
			j = j - 1
		}
		result := make([]int, 0, 1024)
		remi := line[i] - 48
		remj := line[j] - 48
		reme := byte(0)
		fmt.Println(totalL)
		for i < j {
			val := i / 2
			if remi > 0 {
				i++
				if i >= j {
					break
				}
			}
			for ; remi > 0; remi-- {
				if val > 5300 && val < 5310 {
					fmt.Printf("adding %v, with %v, %v\n", val, remi, remj)
				}

				result = append(result, val)
			}
			if reme == 0 {
				reme = line[i] - 48
			}

			val = j / 2
			for remj > 0 && reme > 0 {
				if val > 5300 && val < 5310 {
					fmt.Printf("adding %v, with %v, %v, %v\n", val, remi, remj, reme)
				}
				result = append(result, val)
				remj--
				reme--
			}
			if reme == 0 {
				i++
				remi = line[i] - 48
			}
			if remj == 0 {
				j = j - 2
				if i >= j {
					break
				}
				remj = line[j] - 48
			}
		}
		for ; remj > 0; remj-- {
			result = append(result, (j / 2))
		}
		checksum := 0
		fmt.Println(len(result))
		for i, val := range result {
			checksum = checksum + (i * int(val))
			if checksum < 0 {
				fmt.Println(checksum)
			}
		}
		fmt.Println(checksum)
	}
}

type fileSpace struct {
	size int
	pos  int
}

func P9_1() {
	fi, err := os.Open("input9.txt")
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
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}
		line := scanner.Text()
		totalL := len(line)
		holesCount := totalL / 2
		filesCount := totalL / 2
		if totalL%2 == 1 {
			filesCount++
		}
		result := make([]int, 0, 4098)
		holes := make([]fileSpace, holesCount)
		files := make([]fileSpace, filesCount)
		for i := 0; i < totalL; i++ {
			size := int(line[i] - 48)
			if i%2 == 0 {
				files[i/2] = fileSpace{size, len(result)}
				for j := 0; j < size; j++ {
					result = append(result, i/2)
				}
			} else {
				holes[i/2] = fileSpace{size, len(result)}
				for j := 0; j < size; j++ {
					result = append(result, -1)
				}
			}
		}
		fn := filesCount - 1
		for ; fn >= 0; fn-- {
			file := files[fn]
			for i := 0; i < fn; i++ {
				if holes[i].size >= file.size {
					for a := 0; a < file.size; a++ {
						result[file.pos+a] = -1
						result[holes[i].pos+a] = fn
					}
					holes[i].pos = holes[i].pos + file.size
					holes[i].size = holes[i].size - file.size
					break
				}
			}
		}
		checksum := 0
		for i, val := range result {
			if val > 0 {
				checksum = checksum + (i * int(val))
			}
		}
		fmt.Println(checksum)
	}
}
