package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := "example.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	var task1 int
	var task2 int
	do := true

	for scnr.Scan() {
		t := scnr.Text()
		for _, word := range parsePattern(t) {
			switch {
			case word == "do()":
				do = true
			case word == "don't()":
				do = false
			default:
				task1 += calcMultiply(word)
				if do {
					task2 += calcMultiply(word)
				}
			}
		}
	}

	fmt.Println("task1: ", task1) // 167090022
	fmt.Println("task2: ", task2) // 89823704
}

var mainPattern = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

func parsePattern(line string) []string {
	return mainPattern.FindAllString(line, -1)
}

var digPattern = regexp.MustCompile(`\d+`)

func calcMultiply(mul string) int {
	var numbers []int
	for _, s := range digPattern.FindAllString(mul, -1) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}
	return numbers[0] * numbers[1]
}
