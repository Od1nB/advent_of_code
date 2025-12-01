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

const size = 100

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

	instructions := make([]int, 0, 4659)
	for scnr.Scan() {
		instructions = append(instructions, createInst(scnr.Text()))
	}

	pos := 50
	for _, inst := range instructions {
		pos = move(pos, inst)
		if pos == 0 {
			task1++
		}
	}

	fmt.Println("task1: ", task1) // 1158
	fmt.Println("task2: ", task2) // 6860
}

func move(pos, move int) int {
	switch {
	case move > 0:
		return moveR(pos, move)
	case move < 0:
		return moveL(pos, move)
	default:
		return pos
	}
}

func moveR(pos, move int) int {
	end := pos + move
	crossings := end / size
	task2 += crossings
	return end % size
}

func moveL(pos, move int) int {
	for range -move {
		// modulo arithmetic
		// https://dev.to/avocoaster/how-to-wrap-around-a-range-of-numbers-with-the-modulo-cdo
		pos = (pos - 1 + size) % size
		if pos == 0 {
			task2++
		}
	}
	return pos
}

func createInst(s string) int {
	switch {
	case strings.HasPrefix(s, "L"):
		a, _ := strings.CutPrefix(s, "L")
		return -mustInt(a)
	case strings.HasPrefix(s, "R"):
		a, _ := strings.CutPrefix(s, "R")
		return mustInt(a)
	default:
		panic("not what I know " + s)
	}
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
