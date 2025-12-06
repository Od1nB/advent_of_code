package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

var (
	task1 int
	task2 int
	width = 0
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

	var lines []string
	for scnr.Scan() {
		line := scnr.Text()
		if len(line) > width {
			width = len(line)
		}
		lines = append(lines, line)
	}

	empty := make(map[int]bool)
	for col := range width {
		isEmpty := true
		for row := range lines {
			if lines[row][col] != ' ' {
				isEmpty = false
				break
			}
		}
		empty[col] = isEmpty
	}

	task1 = solveMath(lines, empty, rows)
	task2 = solveMath(lines, empty, columnsReverse)

	fmt.Println("task1: ", task1) // 5335495999141
	fmt.Println("task2: ", task2) // 10142723156431
}

func solveMath(lines []string, empty map[int]bool, iterator func([]string, int, int) iter.Seq[string]) int {
	if len(lines) == 0 {
		return 0
	}

	col, total := 0, 0
	for col < width {
		for col < width && empty[col] {
			col++
		}

		if col >= width {
			break
		}

		start := col
		for col < width && !empty[col] {
			col++
		}

		numbers, operator := readProblem(lines, start, col, iterator)
		total += executeMath(numbers, operator)
	}

	return total
}

func executeMath(nums []int, operator rune) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for _, num := range nums[1:] {
		switch operator {
		case '*':
			result *= num
		case '+':
			result += num
		default:
			panic("invalid operator" + string(operator))
		}
	}
	return result
}

func readProblem(lines []string, start, end int, allNumbers func([]string, int, int) iter.Seq[string]) ([]int, rune) {
	nums := make([]int, 0, end-start+1)
	operator := rune(strings.TrimSpace(lines[len(lines)-1][start:end])[0])

	for token := range allNumbers(lines, start, end) {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		n, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	return nums, operator
}

func rows(lines []string, start, end int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for row := range len(lines) - 1 {
			if !yield(lines[row][start:end]) {
				return
			}
		}
	}
}

func columnsReverse(lines []string, start, end int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for col := end - 1; col >= start; col-- {
			var digits []rune
			for row := range len(lines) - 1 {
				ch := rune(lines[row][col])
				if ch >= '0' && ch <= '9' {
					digits = append(digits, ch)
				}
			}
			if len(digits) > 0 {
				if !yield(string(digits)) {
					return
				}
			}
		}
	}
}
