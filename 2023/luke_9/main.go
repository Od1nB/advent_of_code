package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	lo "github.com/samber/lo"
)

type Entry struct {
	id      int
	history []int
	extras  [][]int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	entries := []Entry{}
	ind := 0
	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		entries = append(entries, Entry{
			id: ind,
			history: lo.Map(strings.Split(scnr.Text(), " "), func(s string, _ int) int {
				n, _ := strconv.Atoi(s)
				return n
			}),
		})
		fillExtras(&entries[ind])
		ind++
	}
	fmt.Println("task1: ", appendHistory(entries))  //2175229206
	fmt.Println("task2: ", prePendHistory(entries)) //942
}

func appendHistory(entries []Entry) int {
	score := 0
	for _, v := range entries {
		for i := len(v.extras) - 1; i >= 0; i-- {
			if isAllZeroes(v.extras[i]) {
				v.extras[i] = append(v.extras[i], 0)
			} else {
				v.extras[i] = append(v.extras[i], v.extras[i][len(v.extras[i])-1]+v.extras[i+1][len(v.extras[i])-1])
			}
		}
		num := v.history[len(v.history)-1] + v.extras[0][len(v.extras[0])-1]
		score += num
	}

	return score
}

func prePendHistory(entries []Entry) int {
	score := 0
	for _, v := range entries {
		for i := len(v.extras) - 1; i >= 0; i-- {
			if isAllZeroes(v.extras[i]) {
				v.extras[i] = append([]int{0}, v.extras[i]...)
			} else {
				v.extras[i] = append([]int{v.extras[i][0] - v.extras[i+1][0]}, v.extras[i]...)
			}
		}
		num := v.history[0] - v.extras[0][0]
		score += num
	}
	return score
}
func fillExtras(e *Entry) {
	intrmdt := make([]int, len(e.history)-1)
	for i := 1; i < len(e.history); i++ {
		intrmdt[i-1] = e.history[i] - (e.history[i-1])
	}
	e.extras = append(e.extras, intrmdt)
	for !isAllZeroes(intrmdt) {
		loopintrmdt := make([]int, len(intrmdt)-1)
		for i := 1; i < len(intrmdt); i++ {
			loopintrmdt[i-1] = intrmdt[i] - (intrmdt[i-1])
		}
		e.extras = append(e.extras, loopintrmdt)
		intrmdt = loopintrmdt
	}
}

func isAllZeroes(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}
