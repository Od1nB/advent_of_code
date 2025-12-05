package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	ranges := make([][2]int, 0, 190)
	ingridients := make([]int, 0, 1100)
	doingRanges := true
	for scnr.Scan() {
		t := scnr.Text()
		if t == " " || t == "\n" || t == "" {
			doingRanges = false
			continue
		}
		if doingRanges {
			ranges = append(ranges, parseRange(t))
		} else {
			n, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			ingridients = append(ingridients, n)
		}
	}

	for _, ing := range ingridients {
		for _, rng := range ranges {
			if inRange(rng, ing) {
				task1++
				break
			}
		}
	}

	ranges = mergeRanges(ranges)
	for _, rng := range ranges {
		task2 += rng[1] - rng[0] + 1
	}

	fmt.Println("task1: ", task1) // 868
	fmt.Println("task2: ", task2) // 354143734113772
}

func parseRange(s string) [2]int {
	spl := strings.Split(s, "-")
	if len(spl) != 2 {
		// fmt.Println(len())
		panic("not 2 numbers and -" + s)
	}

	first, err := strconv.Atoi(spl[0])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(spl[1])
	if err != nil {
		panic(err)
	}

	return [2]int{first, second}
}

func inRange(rng [2]int, cand int) bool {
	return cand >= rng[0] && cand <= rng[1]
}

func mergeRanges(ranges [][2]int) [][2]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		curr := ranges[i]

		if curr[0] <= last[1]+1 {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
