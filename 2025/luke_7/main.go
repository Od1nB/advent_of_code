package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var (
	task1           int
	task2           int
	TachyonManiFold = make([][]rune, 0, 142)
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
		TachyonManiFold = append(TachyonManiFold, []rune(scnr.Text()))
	}
	startX := slices.Index(TachyonManiFold[0], 'S')
	TachyonManiFold[1][startX] = '|'

	for y := range len(TachyonManiFold) {
		for x, symb := range TachyonManiFold[y] {
			if symb != '|' {
				continue
			}
			blastDownward(x, y)
		}
	}

	reachable := map[int]int{startX: 1}
	for y := range len(TachyonManiFold) {
		next := make(map[int]int, len(reachable)+len(reachable)/2)

		for x, paths := range reachable {
			if TachyonManiFold[y][x] != '^' {
				next[x] += paths
				continue
			}
			if inside(x-1, y) {
				next[x-1] += paths
			}
			if inside(x+1, y) {
				next[x+1] += paths
			}
		}
		reachable = next
	}

	for _, count := range reachable {
		task2 += count
	}

	fmt.Println("task1: ", task1)                          // 1619
	fmt.Println("task2: ", task2, task2 == 23607984027985) // 23607984027985
}

func blastDownward(x, y int) {
	if !inside(y+1, x) {
		return
	}
	switch TachyonManiFold[y+1][x] {
	case '.':
		TachyonManiFold[y+1][x] = '|'
		return
	case '^':
		if inside(y+1, x+1) {
			TachyonManiFold[y+1][x+1] = '|'
		}
		if inside(y+1, x-1) {
			TachyonManiFold[y+1][x-1] = '|'
		}
		task1++
	default:
		return
	}
}

func inside(x, y int) bool {
	return y < len(TachyonManiFold) && x < len(TachyonManiFold[y])
}
