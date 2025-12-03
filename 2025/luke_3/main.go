package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	task1 int
	task2 int
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)

	for scnr.Scan() {
		t := scnr.Text()
		ints := parseBatteries(t)
		task1 += findBiggest(ints)
		task2 += findBiggestN(ints, 12)
	}

	fmt.Println("task1: ", task1) // 17535
	fmt.Println("task2: ", task2) // 173577199527257
}

func parseBatteries(s string) []int {
	ints := make([]int, 0, len(s))
	for _, r := range s {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			panic("NaN " + string(r))
		}
		if i <= 0 || i > 9 {
			panic("outside of range " + string(r))
		}
		ints = append(ints, i)
	}
	return ints
}

func findBiggest(ints []int) int {
	first := ints[0]
	var second int
	for ind, n := range ints[1:] {
		switch {
		case n > first && ind+1 != len(ints[1:]):
			first = n
			second = 0
		case n > second:
			second = n
		}
	}

	return bigNumber(first, second)
}

func findBiggestN(ints []int, count int) int {
	if count > len(ints) {
		count = len(ints)
	}

	result := 0
	remaining := make([]int, len(ints))
	copy(remaining, ints)

	for i := 0; i < count; i++ {
		bestVal := 0
		bestPos := 0

		for pos := 0; pos <= len(remaining)-count+i; pos++ {
			if remaining[pos] > bestVal {
				bestVal = remaining[pos]
				bestPos = pos
			}
		}

		result = result*10 + bestVal
		remaining = remaining[bestPos+1:]
	}

	return result
}

func bigNumber(first, second int) int {
	n, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
	if err != nil {
		panic(err)
	}
	return n
}
