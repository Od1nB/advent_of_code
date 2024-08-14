package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	BOLDER  = 79
	NOTHING = 46
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, task2 := 0, 0
	lines := []string{}
	byteLines := [][]byte{}
	cubeRocks := map[string]bool{}
	roundRocks := map[string]bool{}
	ind := 0
	for scnr.Scan() {
		t := scnr.Text()
		fmt.Println(t)
		byteLines = append(byteLines, []byte(t))
		lines = append(lines, t)
		for x, str := range t {
			switch string(str) {
			case "#":
				cubeRocks[intsToCoord(x, ind)] = true
			case "O":
				roundRocks[intsToCoord(x, ind)] = true
			}
		}
		ind++
	}

	slidNorth := slideNorth(byteLines)
	fmt.Println((bytes.Join(slidNorth, []byte{})))
	task1 = sumPuzzleWeight(slidNorth)

	//fmt.Println(cubeRocks)

	fmt.Println("task1: ", task1)
	fmt.Println("task2: ", task2)
}

func sumPuzzleWeight(puzzle [][]byte) int {
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])

	sum := 0
	for line := 0; line < puzzleH; line++ {
		for col := 0; col < puzzleW; col++ {
			if puzzle[line][col] == BOLDER {
				sum += puzzleH - line
			}
		}
	}
	return sum
}

func slideNorth(lines [][]byte) [][]byte {
	/*
		height, width := len(lines), len(lines[0])

		moved := 0
		for line := 1; line < height; line++ {
			for col := 0; col < width; col++ {
				if string(lines[line][col]) == "O" {
					if lines[line-1][col] == POINT {
						lines[line-1][col] = CUBE
						lines[line][col] = POINT
						moved++
					}

				}
			}
		}
		if moved == 0 {
			return lines
		} else {
			return slideNorth(lines)
		}
	*/

	puzzle := lines
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])
	move := 0
	for line := 1; line < puzzleH; line++ {
		for col := 0; col < puzzleW; col++ {
			if string(puzzle[line][col]) == "O" {
				// Rolling bolder can go up
				switch puzzle[line-1][col] {
				case NOTHING:
					puzzle[line-1][col] = BOLDER
					puzzle[line][col] = NOTHING
					move++
				}
			}
		}
	}
	if move == 0 {
		return puzzle
	} else {
		return slideNorth(puzzle)
	}

}
func intsToCoord(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
