package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var task1 int
	var task2 int

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

	for scnr.Scan() {
		t := scnr.Text()
		nums := stringToSlice(t)

		if isSafe(nums) {
			task1 += 1
		}
		if isSafeMissingOne(nums) {
			task2 += 1
		}

	}

	fmt.Println("task1: ", task1) // 224
	fmt.Println("task2: ", task2) // 293
}

func isSafeMissingOne(nums []int) bool {
	for ind := range nums {
		cp := slices.Delete(slices.Clone(nums), ind, ind+1)
		if isSafe(cp) {
			return true
		}

	}

	return false
}

func isSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	ascending := nums[1]-nums[0] > 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if ascending && diff <= 0 {
			return false
		}
		if !ascending && diff >= 0 {
			return false
		}
	}

	return true
}

func stringToSlice(input string) []int {
	parts := strings.Fields(input)
	res := make([]int, 0, len(parts))
	for s := range slices.Values(parts) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}

	return res
}
