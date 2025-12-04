package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	task1   int
	task2   int
	factory map[int][]rune
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

	ind := 0
	factory = make(map[int][]rune, 140)
	for scnr.Scan() {
		factory[ind] = parsePaperRow(scnr.Text())
		ind++
	}

	for y := range factory {
		for x := range factory[y] {
			if factory[y][x] != '@' {
				continue
			}
			num := adjecentRolls(x, y)
			if num < 4 {
				task1++
			}
		}
	}

	for changed := true; changed; {
		changed = false
		for y := range factory {
			for x := range factory[y] {
				if factory[y][x] != '@' {
					continue
				}
				num := adjecentRolls(x, y)
				if num < 4 {
					task2++
					factory[y][x] = '.'
					changed = true
				}
			}
		}
	}

	fmt.Println("task1: ", task1) // 1547
	fmt.Println("task2: ", task2) // 8948
}

func parsePaperRow(s string) []rune {
	rr := make([]rune, 0, len(s))
	for _, ru := range s {
		rr = append(rr, ru)
	}
	return rr
}

func adjecentRolls(x, y int) int {
	if y >= len(factory) {
		panic("too big y")
	}

	if x >= len(factory[y]) {
		panic("too big x")
	}

	res := 0

	// BL
	if y+1 < len(factory) && x-1 >= 0 && factory[y+1][x-1] == '@' {
		res++
	}
	// BM
	if y+1 < len(factory) && factory[y+1][x] == '@' {
		res++
	}
	// BR
	if y+1 < len(factory) && x+1 < len(factory[y+1]) && factory[y+1][x+1] == '@' {
		res++
	}
	// ML
	if x-1 >= 0 && factory[y][x-1] == '@' {
		res++
	}
	// MR
	if x+1 < len(factory[y]) && factory[y][x+1] == '@' {
		res++
	}
	// TL
	if y-1 >= 0 && x-1 >= 0 && factory[y-1][x-1] == '@' {
		res++
	}
	// TM
	if y-1 >= 0 && factory[y-1][x] == '@' {
		res++
	}
	// TR
	if y-1 >= 0 && x+1 < len(factory[y-1]) && factory[y-1][x+1] == '@' {
		res++
	}

	return res
}
