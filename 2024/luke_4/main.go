package main

import (
	"bufio"
	"fmt"
	"os"
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

	var inp []string
	for scnr.Scan() {
		t := scnr.Text()
		inp = append(inp, t)
	}
	part1(inp, &task1)
	part2(inp, &task2)

	fmt.Println("task1: ", task1) // 2685
	fmt.Println("task2: ", task2) // 2048
}

func checkXMAS(inp []string, x, y, dx, dy int) int {
	var res int
	if y+3*dy >= 0 &&
		x+3*dx >= 0 &&
		y+3*dy < len(inp) &&
		x+3*dx < len(inp[y+3*dy]) &&
		inp[y+dy][x+dx] == 'M' &&
		inp[y+2*dy][x+2*dx] == 'A' &&
		inp[y+3*dy][x+3*dx] == 'S' {
		res++
	}
	return res
}

func part1(inp []string, res *int) {
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			if inp[y][x] == 'X' {
				*res += checkXMAS(inp, x, y, -1, -1)
				*res += checkXMAS(inp, x, y, -1, 0)
				*res += checkXMAS(inp, x, y, -1, 1)
				*res += checkXMAS(inp, x, y, 0, -1)
				*res += checkXMAS(inp, x, y, 0, 1)
				*res += checkXMAS(inp, x, y, 1, -1)
				*res += checkXMAS(inp, x, y, 1, 0)
				*res += checkXMAS(inp, x, y, 1, 1)
			}
		}
	}
}

func part2(inp []string, res *int) {
	for y := 1; y < len(inp)-1; y++ {
		for x := 1; x < len(inp[y])-1; x++ {
			if inp[y][x] == 'A' {
				var c int
				if inp[y-1][x-1] == 'M' && inp[y+1][x+1] == 'S' {
					c++
				}
				if inp[y+1][x+1] == 'M' && inp[y-1][x-1] == 'S' {
					c++
				}
				if inp[y+1][x-1] == 'M' && inp[y-1][x+1] == 'S' {
					c++
				}
				if inp[y-1][x+1] == 'M' && inp[y+1][x-1] == 'S' {
					c++
				}
				if c == 2 {
					*res++
				}
			}
		}
	}
}
