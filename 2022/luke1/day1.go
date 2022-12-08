package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(getNum())
}

func getNum() int {
	f, err := os.Open("input.txt")
	if err != nil {
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	temp, biggest, second, third := 0, 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if temp > third {
				out := biggestThree([]int{temp, biggest, second, third})
				biggest, second, third = out[2], out[1], out[0]
			}
			temp = 0
		}
		if num, ok := strconv.Atoi(line); ok == nil {
			temp += num
		}
	}
	return biggest + second + third
}

func biggestThree(s []int) []int {
	sort.Ints(s)
	return s[len(s)-3:]
}
