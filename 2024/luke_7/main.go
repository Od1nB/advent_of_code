package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	sum  int
	nums []int
}

func parseRec(s string) record {
	before, after, _ := strings.Cut(s, ":")
	sum, err := strconv.Atoi(before)
	if err != nil {
		panic(err)
	}
	strs := strings.Split(strings.TrimSpace(after), " ")
	nums := []int{}
	for _, v := range strs {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return record{sum: sum, nums: nums}
}

func main() {
	records := []record{}
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
	task1, task2 := 0, 0

	for scnr.Scan() {
		t := scnr.Text()
		records = append(records, parseRec(t))
	}

	for _, r := range records {
		task1 += r.score([]operation{add, multiply})
		task2 += r.score([]operation{add, multiply, concat})
	}

	fmt.Println("task1: ", task1) // 4364915411363
	fmt.Println("task2: ", task2) // 38322057216320
}

func (r record) score(ops []operation) int {
	var valid bool
	calcAllCombs(r.nums, 1, r.nums[0], r.sum, ops, &valid)
	if !valid {
		return 0
	}
	return r.sum
}

func calcAllCombs(nums []int, ind, currRes, goal int, ops []operation, valid *bool) {
	if ind >= len(nums) {
		if currRes == goal {
			*valid = true
		}
		return
	}

	for _, op := range ops {
		newRes := op(currRes, nums[ind])
		calcAllCombs(nums, ind+1, newRes, goal, ops, valid)
	}
}

type operation func(a, b int) int

func add(a, b int) int {
	return a + b
}

func concat(a, b int) int {
	s := fmt.Sprintf("%d%d", a, b)
	i, _ := strconv.Atoi(s)
	return i
}

func multiply(a, b int) int {
	return a * b
}
