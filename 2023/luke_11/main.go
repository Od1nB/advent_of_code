package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	universe := []string{}
	for scnr.Scan() {
		t := scnr.Text()
		universe = append(universe, t)
	}

	emptyRows, emptyColumns := getEmptyRows(universe), getEmptyColumns(universe)

	galaxies := getGalaxies(universe, emptyRows, emptyColumns)

	fmt.Println("task1: ", calcScore(galaxies))
	fmt.Println("task2: ", 0)
}

func calcScore(galaxies map[int]Coord) int {
	acc := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			acc += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}
	return acc
}

func getGalaxies(universe []string, emptyRows, emptyColumns map[int]bool) map[int]Coord {
	galaxies := map[int]Coord{}
	n := 0
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[i]); j++ {
			if string(universe[i][j]) == "#" {
				nx, ny := 0, 0
				for k := 0; k < i; k++ {
					if emptyRows[k] {
						ny++
					}
				}
				for k := 0; k < j; k++ {
					if emptyColumns[k] {
						nx++
					}
				}
				galaxies[n] = Coord{x: i + ny, y: j + nx}
				n++
			}
		}
	}
	return galaxies
}

func getEmptyRows(lines []string) map[int]bool {
	rowMap := map[int]bool{}
	for i, l := range lines {
		if strings.Contains(l, "#") {
			continue
		}
		rowMap[i] = true
	}
	return rowMap
}

func getEmptyColumns(lines []string) map[int]bool {
	columnMap := map[int]bool{}
	for i := range lines[0] {
		l := ""
		for j := range lines {
			l += string(lines[j][i])
		}
		if strings.Contains(l, "#") {
			continue
		}
		columnMap[i] = true
	}
	return columnMap

}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
