package main

import (
	"bufio"
	"fmt"
	"os"
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

	for scnr.Scan() {
		t := scnr.Text()
		for rng := range strings.SplitSeq(t, ",") {
			if len(rng) == 0 {
				continue
			}
			task1 += sum(invalidT1(parseRange(rng)))
			task2 += sum(invalidT2(parseRange(rng)))
		}
	}

	fmt.Println("task1: ", task1) // 32976912643
	fmt.Println("task2: ", task2) // 54446379122
}

func sum(ints []int) int {
	var res int
	for _, n := range ints {
		res += n
	}
	return res
}

func parseRange(s string) (int, int) {
	if s == " " {
		return 0, 0
	}
	ss := strings.Split(s, "-")
	if len(ss) != 2 {
		panic("not 2 nums" + s)
	}
	first, err := strconv.Atoi(ss[0])
	if err != nil {
		panic("nan" + ss[0])
	}
	last, err := strconv.Atoi(ss[1])
	if err != nil {
		panic("nan" + ss[1])
	}
	return first, last
}

func invalidT1(start, end int) []int {
	if start == 0 && end == 0 {
		return []int{}
	}
	invalids := []int{}
	for ind := start; ind <= end; ind++ {
		if !equalNum(ind) {
			continue
		}
		invalids = append(invalids, ind)
	}
	return invalids
}

func equalNum(i int) bool {
	s := strconv.Itoa(i)
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func invalidT2(start, end int) []int {
	if start == 0 && end == 0 {
		return []int{}
	}
	invalids := []int{}
	for ind := start; ind <= end; ind++ {
		if !isRepeatedPattern(strconv.Itoa(ind)) {
			continue
		}
		invalids = append(invalids, ind)
	}
	return invalids
}

func isRepeatedPattern(s string) bool {
	strLen := len(s)
	for candLen := 1; candLen <= strLen/2; candLen++ {
		if strLen%candLen != 0 {
			continue
		}
		pattern := s[:candLen]
		repeats := strLen / candLen

		if repeats < 2 {
			continue
		}
		if strings.Repeat(pattern, repeats) == s {
			return true
		}
	}
	return false
}
