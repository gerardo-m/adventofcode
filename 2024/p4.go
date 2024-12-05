package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	XMAS	SAMX	X	  X		S	  S
	M		A		 M	 M		 A   A
	A		M		  A A		  M M
	S		X		   S		   X

	{
		pos: expect M [pos+1, pos + lineL, pos + lineL + 1,  pos+ lineL - 1]
	}

	pos =>
	struct{
		expected,
		next,

	}

	X or M
		CHECK if next is -1. COUNT + 1
	-2 if node is going to be ignored
*/

type node struct {
	expected byte
	next     int
	previous int
}

func P4(readLine func(bufio.Reader) (string, error)) {
	fi, err := os.Open("input4.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)
	count := 0
	futureNodes := make(map[int][]node)
	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lineL := len(line)
		for i := 0; i < lineL; i++ {
			currPos := lineNumber*lineL + i
			if line[i] == 'X' {
				if i+3 < lineL {
					futureNodes[currPos+1] = append(futureNodes[currPos+1], node{'M', currPos + 2, currPos})
					futureNodes[currPos+2] = append(futureNodes[currPos+2], node{'A', currPos + 3, currPos + 1})
					futureNodes[currPos+3] = append(futureNodes[currPos+3], node{'S', -1, currPos + 2})

					futureNodes[currPos+lineL+1] = append(futureNodes[currPos+lineL+1], node{'M', currPos + lineL*2 + 2, currPos})
					futureNodes[currPos+lineL*2+2] = append(futureNodes[currPos+lineL*2+2], node{'A', currPos + lineL*3 + 3, currPos + lineL + 1})
					futureNodes[currPos+lineL*3+3] = append(futureNodes[currPos+lineL*3+3], node{'S', -1, currPos + lineL*2 + 2})
				}
				if i > 2 {
					futureNodes[currPos+lineL-1] = append(futureNodes[currPos+lineL-1], node{'M', currPos + lineL*2 - 2, currPos})
					futureNodes[currPos+lineL*2-2] = append(futureNodes[currPos+lineL*2-2], node{'A', currPos + lineL*3 - 3, currPos + lineL - 1})
					futureNodes[currPos+lineL*3-3] = append(futureNodes[currPos+lineL*3-3], node{'S', -1, currPos + lineL*2 - 2})
				}
				futureNodes[currPos+lineL] = append(futureNodes[currPos+lineL], node{'M', currPos + lineL*2, currPos})
				futureNodes[currPos+lineL*2] = append(futureNodes[currPos+lineL*2], node{'A', currPos + lineL*3, currPos + lineL})
				futureNodes[currPos+lineL*3] = append(futureNodes[currPos+lineL*3], node{'S', -1, currPos + lineL*2})
			}
			if line[i] == 'S' {
				if i+3 < lineL {
					futureNodes[currPos+1] = append(futureNodes[currPos+1], node{'A', currPos + 2, currPos})
					futureNodes[currPos+2] = append(futureNodes[currPos+2], node{'M', currPos + 3, currPos + 1})
					futureNodes[currPos+3] = append(futureNodes[currPos+3], node{'X', -1, currPos + 2})

					futureNodes[currPos+lineL+1] = append(futureNodes[currPos+lineL+1], node{'A', currPos + lineL*2 + 2, currPos})
					futureNodes[currPos+lineL*2+2] = append(futureNodes[currPos+lineL*2+2], node{'M', currPos + lineL*3 + 3, currPos + lineL + 1})
					futureNodes[currPos+lineL*3+3] = append(futureNodes[currPos+lineL*3+3], node{'X', -1, currPos + lineL*2 + 2})
				}
				if i > 2 {
					futureNodes[currPos+lineL-1] = append(futureNodes[currPos+lineL-1], node{'A', currPos + lineL*2 - 2, currPos})
					futureNodes[currPos+lineL*2-2] = append(futureNodes[currPos+lineL*2-2], node{'M', currPos + lineL*3 - 3, currPos + lineL - 1})
					futureNodes[currPos+lineL*3-3] = append(futureNodes[currPos+lineL*3-3], node{'X', -1, currPos + lineL*2 - 2})
				}
				futureNodes[currPos+lineL] = append(futureNodes[currPos+lineL], node{'A', currPos + lineL*2, currPos})
				futureNodes[currPos+lineL*2] = append(futureNodes[currPos+lineL*2], node{'M', currPos + lineL*3, currPos + lineL})
				futureNodes[currPos+lineL*3] = append(futureNodes[currPos+lineL*3], node{'X', -1, currPos + lineL*2})
			}
			posibleNodes := futureNodes[currPos]
			for _, aNode := range posibleNodes {
				if aNode.previous == -2 {
					if aNode.next >= 0 {
						for j := 0; j < len(futureNodes[aNode.next]); j++ {
							if futureNodes[aNode.next][j].previous == currPos {
								futureNodes[aNode.next][j].previous = -2
								break
							}
						}
					}
					continue
				}
				if aNode.expected != line[i] {
					if aNode.next >= 0 {
						for j := 0; j < len(futureNodes[aNode.next]); j++ {
							if futureNodes[aNode.next][j].previous == currPos {
								futureNodes[aNode.next][j].previous = -2
								break
							}
						}
					}
				} else if aNode.next == -1 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

type expectedLetter struct {
	word   string
	lPos   int
	offset int
}

/*
S..S..S
.A.A.A.
..MMM..
SAMXMAS
..MMM..
.A.A.A.
S..S..S
7 3
...XMAS ...SMAX
..MMM.. ..AAA..
.A.A.A. .M.M.M.
S..S..S X..X..X

	struct{
		word,
		lPos,
		offset,
	}

	1
	lineL - 1
	lineL
	lineL + 1

XMAS, SAMX
expectedLetter

	for every letter in the puzzle{
		if letter is beggining of a word
			expectedLetter[pos+offset] << struct(word, 1, offset)
		if letter is in expectedLetter[pos]
			if it is the last letter in the word
				success
			else
				expectedLetter[pos+offset] << struct(word, previous lPos + 1, offset)
	}
*/
func P4_o() {
	fi, err := os.Open("input4.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(fi)
	words := []string{
		"XMAS",
		"SAMX",
	}
	count := 0
	expectedLetters := make(map[int][]expectedLetter)
	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		lineL := len(line)
		for i := 0; i < lineL; i++ {
			currPos := lineNumber*lineL + i
			for _, word := range words {
				if line[i] == word[0] {
					if i+len(word) <= lineL {
						expectedLetters[currPos+1] = append(expectedLetters[currPos+1], expectedLetter{word, 1, 1})
						expectedLetters[currPos+lineL+1] = append(expectedLetters[currPos+lineL+1], expectedLetter{word, 1, lineL + 1})
					}
					if i >= len(word)-1 {
						expectedLetters[currPos+lineL-1] = append(expectedLetters[currPos+lineL-1], expectedLetter{word, 1, lineL - 1})
					}
					expectedLetters[currPos+lineL] = append(expectedLetters[currPos+lineL], expectedLetter{word, 1, lineL})
				}
			}
			for _, expLetter := range expectedLetters[currPos] {
				if expLetter.word[expLetter.lPos] == line[i] {
					if expLetter.lPos == len(expLetter.word)-1 {
						count++
					} else {
						expectedLetters[currPos+expLetter.offset] = append(expectedLetters[currPos+expLetter.offset], expectedLetter{expLetter.word, expLetter.lPos + 1, expLetter.offset})
					}
				}
			}
		}
	}
	fmt.Println(count)
}
