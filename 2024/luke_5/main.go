package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	cmp "github.com/Od1nB/advent_of_code/tools/comparison"
	"github.com/samber/lo"
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
	task1, task2 := 0, 0
	var afterNewline bool
	var correct, incorrect [][]int
	orders := map[int][]int{}

	for scnr.Scan() {
		t := scnr.Text()
		if t == "" {
			afterNewline = true
			continue
		}

		if !afterNewline {
			s1, s2, _ := strings.Cut(t, "|")
			n1, err := strconv.Atoi(s1)
			n2, err := strconv.Atoi(s2)
			if err != nil {
				panic(err)
			}
			orders[n1] = append(orders[n1], n2)
		}

		if afterNewline {
			nums := lo.Map(strings.Split(t, ","), func(s string, _ int) int {
				num, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				return num
			})
			if correctOrder(orders, nums) {
				correct = append(correct, nums)
				continue
			}
			incorrect = append(incorrect, nums)
		}
	}

	for _, list := range correct {
		calcScore(list, &task1)
	}
	for _, list := range incorrect {
		calcScore(repair(orders, list), &task2)
	}

	fmt.Println("task1: ", task1) // 7365
	fmt.Println("task2: ", task2) // 5770
}

func correctOrder(orders map[int][]int, update []int) bool {
	seen := []int{}
	for _, num := range update {
		if os, ok := orders[num]; ok {
			if wrong := cmp.Intersection(os, append(seen, num)); len(wrong) != 0 {
				return false
			}
			seen = append(seen, num)

		}
	}
	return true
}

func calcScore(nums []int, res *int) {
	*res += nums[len(nums)/2]
}

func repair(orders map[int][]int, update []int) []int {
	fixed := make([]int, 0, len(update))
	for ind, num := range update {
		fixed = append(fixed, num)
		if wrong := cmp.Intersection(orders[num], fixed[:ind]); len(wrong) != 0 {
			fixed = slices.DeleteFunc(fixed, func(i int) bool {
				return slices.Contains(wrong, i)
			})
			fixed = append(fixed, wrong...)
		}
	}

	return fixed
}
