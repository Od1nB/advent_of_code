package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	lo "github.com/samber/lo"
)

type Direction int

const (
	FORWARD Direction = iota
	BACKWARDS
)

type Record struct {
	str    string
	groups []int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, task2 := 0, 0
	records_p1 := []*Record{}

	for scnr.Scan() {
		t := scnr.Text()
		split := strings.Split(t, " ")
		record := &Record{split[0], lo.Map(strings.Split(split[1], ","), func(s string, _ int) int {
			num, _ := strconv.Atoi(s)
			return num
		})}
		records_p1 = append(records_p1, record)
	}
	t1 := time.Now()
	for _, rec := range records_p1 {
		task1 += numberOfArr(rec.str, rec.groups)
	}
	fmt.Println("time taken: ", time.Since(t1))

	fmt.Println("task1: ", task1) //7771
	fmt.Println("task2: ", task2)
}

func numberOfArr(str string, ints []int) int {
	if len(str) == 0 {
		if len(ints) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if strings.HasPrefix(str, "?") {
		return numberOfArr(strings.Replace(str, "?", ".", 1), ints) + numberOfArr(strings.Replace(str, "?", "#", 1), ints)
	}
	if strings.HasPrefix(str, ".") {
		return numberOfArr(strings.TrimPrefix(str, "."), ints)
	}

	if strings.HasPrefix(str, "#") {
		if len(ints) == 0 {
			return 0
		}
		if len(str) < ints[0] {
			return 0
		}
		if strings.Contains(str[0:ints[0]], ".") {
			return 0
		}
		if len(ints) > 1 {
			if len(str) < ints[0]+1 || string(str[ints[0]]) == "#" {
				return 0
			}
			return numberOfArr(str[ints[0]+1:], ints[1:])
		} else {
			return numberOfArr(str[ints[0]:], ints[1:])
		}
	}
	return 0
}
