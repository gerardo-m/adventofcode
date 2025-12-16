package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Problem struct {
	Part1 func(*os.File)
	Part2 func(*os.File)
}

var problems = []Problem{
	{P1, P1_2},
	{P2, P2_2},
	{P3, P3_2},
	{P4, P4_2},
}

func main() {
	isTest := flag.Bool("t", false, "Execute with test data")
	part := flag.Int("p", 1, "Must be 1 or 2, defines the part of the problem to execute")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		printHelp()
		return
	}
	if args[0] == "l" || args[0] == "list" {
		printList()
		return
	}
	if args[0] == "h" || args[0] == "help" {
		printHelp()
		return
	}
	problemNumber, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("First argument must be help, list or a problem number. It was %v\n", args[0])
		return
	}
	if problemNumber < 1 || problemNumber > len(problems) {
		fmt.Println("Invalid problem number")
		return
	}
	sourceName := "p"
	if *isTest {
		sourceName = "t"
	}
	sourceName = sourceName + strconv.Itoa(problemNumber) + ".txt"
	fi, err := os.Open("./data/" + sourceName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	problemNumber = problemNumber - 1
	if *part < 1 || *part > 2 {
		*part = 1
	}
	if *part == 1 {
		problems[problemNumber].Part1(fi)
	} else {
		problems[problemNumber].Part2(fi)
	}

}

func printHelp() {
	// TODO
	fmt.Println("Help")
}

func printList() {
	// TODO
	fmt.Println("List")
}

func runP2() {
	fi, err := os.Open("./data/p2.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	P2_2(fi)
}
