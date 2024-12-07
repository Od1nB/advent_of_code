package main

import (
	"bufio"
	"fmt"
	"os"
)

type Posistion struct {
	x     int
	y     int
	value rune
}

var (
	north = Posistion{x: 0, y: -1}
	south = Posistion{x: 0, y: 1}
	east  = Posistion{x: 1, y: 0}
	west  = Posistion{x: -1, y: 0}
)

var (
	m          [][]Posistion
	initialPos *Posistion
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
	for scnr.Scan() {
		t := scnr.Text()
		row := makePosRow(t, len(m))
		found, x := findGuard(row)
		if found {
			initialPos = &row[x]
			row[x].value = 'X'
		}
		m = append(m, row)

	}
	if initialPos != nil {
		fmt.Println("found guard at ", initialPos.x, " ", initialPos.y)
	}

	mCopy := copyRecur(m)
	moveThroughRec(mCopy, *initialPos, north, map[Posistion]Posistion{})
	task1 += calcScore(mCopy)

	for y, xPos := range m {
	inner:
		for x, pos := range xPos {
			if equalPos(pos, *initialPos) || pos.value == '#' {
				continue inner
			}
			cp := copyRecur(m)
			cp[y][x].value = '#'

			loop := moveThroughRec(cp, *initialPos, north, map[Posistion]Posistion{})
			if loop {
				task2 += 1
			}
		}
	}

	fmt.Println("task1: ", task1) // 5531
	fmt.Println("task2: ", task2) // 2165
}

func moveThroughRec(m [][]Posistion, currPos, direction Posistion, dejavu map[Posistion]Posistion) bool {
	if !within(m, currPos) {
		return false
	}

	if dir, ok := dejavu[currPos]; ok && equalPos(dir, direction) {
		return true
	}

	if blocked(m, currPos) {
		currPos.x = currPos.x - direction.x
		currPos.y = currPos.y - direction.y
		return moveThroughRec(m, currPos, turnRight(direction), dejavu)
	}

	if m[currPos.y][currPos.x].value == '.' {
		m[currPos.y][currPos.x].value = 'X'
		dejavu[currPos] = direction
	}
	currPos.x, currPos.y = currPos.x+direction.x, currPos.y+direction.y
	return moveThroughRec(m, currPos, direction, dejavu)
}

func within(m [][]Posistion, p Posistion) bool {
	return p.y >= 0 && p.y < len(m) && p.x >= 0 && p.x < len(m[p.y])
}

func blocked(m [][]Posistion, p Posistion) bool {
	return m[p.y][p.x].value == '#'
}

func copyRecur(src [][]Posistion) [][]Posistion {
	dst := make([][]Posistion, len(src))
	for ind := range src {
		dst[ind] = make([]Posistion, len(src[ind]))
		copy(dst[ind], src[ind])
	}
	return dst
}

func makePosRow(s string, y int) []Posistion {
	row := make([]Posistion, len(s))
	for x, v := range s {
		row[x].value = v
		row[x].x = x
		row[x].y = y
	}
	return row
}

func calcScore(m [][]Posistion) (score int) {
	for _, row := range m {
		for _, pos := range row {
			if pos.value != 'X' {
				continue
			}
			score += 1
		}
	}
	return
}

func turnRight(currDir Posistion) Posistion {
	switch {
	case equalPos(currDir, north):
		return east
	case equalPos(currDir, east):
		return south
	case equalPos(currDir, south):
		return west
	case equalPos(currDir, west):
		return north
	default:
		panic("No way to turn")
	}
}

func findGuard(positions []Posistion) (bool, int) {
	for ind, pos := range positions {
		if pos.value == '^' {
			return true, ind
		}
	}
	return false, 0
}

func equalPos(dirA, dirB Posistion) bool {
	return dirA.x == dirB.x && dirA.y == dirB.y
}
