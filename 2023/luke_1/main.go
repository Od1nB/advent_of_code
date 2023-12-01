package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var strToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var task1 int
	var task2 int

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)

	for scnr.Scan() {
		t1 := map[string]int{"f": 0, "f_ind": 0, "l": 0, "l_ind": 0}
		t2 := map[string]int{"f": 0, "f_ind": 0, "l": 0, "l_ind": 0}
		t := scnr.Text()
		c := findNumbersAndIndexes(t)
		c1 := findNumberLettersAndIndexes(t)
		mapFill(t1, c)
		mapFill(t2, append(c, c1...))
		r1, _ := strconv.Atoi(fmt.Sprintf("%d%d", t1["f"], t1["l"]))
		r2, _ := strconv.Atoi(fmt.Sprintf("%d%d", t2["f"], t2["l"]))
		task1 += r1
		task2 += r2
	}
	fmt.Println("task1: ", task1) //first 55447
	fmt.Println("task2: ", task2) // second 5470
}

func findNumbersAndIndexes(input string) [][]int {
	var numbers [][]int

	for i, char := range input {
		if unicode.IsDigit(char) {
			digit := int(char - '0')
			numbers = append(numbers, []int{digit, i})
		}
	}

	return numbers
}

func findNumberLettersAndIndexes(input string) [][]int {
	var numbers [][]int
	subs := getAllSubstringsWithIndex(input)

	for str, ind_arr := range subs {
		for _, v := range ind_arr {
			if val, ok := strToInt[str]; ok {
				numbers = append(numbers, []int{val, v})
			}
		}
	}
	return numbers
}

func getAllSubstringsWithIndex(input string) map[string][]int {
	substrings := make(map[string][]int)

	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ {
			substring := input[i:j]
			if arrval, ok := substrings[substring]; ok {
				substrings[substring] = append(arrval, i)
			} else {
				substrings[substring] = []int{i}
			}
		}
	}

	return substrings
}

func mapFill(m map[string]int, values [][]int) {
	for _, v := range values {
		if m["f_ind"] > v[1] || m["f"] == 0 {
			m["f_ind"] = v[1]
			m["f"] = v[0]
		}
		if m["l_ind"] < v[1] || m["l"] == 0 {
			m["l_ind"] = v[1]
			m["l"] = v[0]
		}
	}
}
