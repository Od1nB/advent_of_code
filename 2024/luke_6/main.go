package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

var m []string
var initialPos *Posistion

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
	var guard *Posistion

	for scnr.Scan() {
		t := scnr.Text()
		//fmt.Println(t)
		m = append(m, t)
		found, x := findGuard(t)
		if !found {
			continue
		}
		guard, initialPos = &Posistion{x: x, y: len(m) - 1}, &Posistion{x: x, y: len(m) - 1}
		m[len(m)-1] = strings.Replace(m[len(m)-1], "^", "X", 1)
	}
	if guard != nil {
		fmt.Println("found guard at ", guard.x, " ", guard.y)
	}

	posMap := createPosMap(m)
	direction := north
	for {
		if wallAhead(*guard, direction) {
			direction = turnRight(direction)
		}
		done := move(guard, &direction)
		if done {
			break
		}

	}
	task1 += calcScore(m)

	for y, xPos := range posMap {
	inner:
		for x, pos := range xPos {
			if equalPos(pos, *initialPos) || pos.value == '#' {
				continue inner
			}
			cp := copyRecur(posMap)
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

func createPosMap(m []string) [][]Posistion {
	posMap := make([][]Posistion, len(m))
	for y, xarr := range m {
		row := make([]Posistion, len(xarr))
		for x, r := range xarr {
			row[x].value = r
			row[x].x = x
			row[x].y = y
		}
		posMap[y] = row
	}

	return posMap
}

func calcScore(m []string) (score int) {
	for _, str := range m {
		score += strings.Count(str, "X")
	}
	return
}

func move(guard, posDelta *Posistion) bool {
	x := guard.x
	y := guard.y
	if !possibleMove(*guard, *posDelta) {
		return true
	}
	guard.x = x + posDelta.x
	guard.y = y + posDelta.y
	arr := []byte(m[guard.y])
	arr[guard.x] = 'X'
	m[guard.y] = string(arr)

	return false
}

func possibleMove(guard, posDelta Posistion) bool {
	x := guard.x
	y := guard.y
	if y+posDelta.y == len(m) || x+posDelta.x == len(m[y]) {
		return false
	}
	return true
}

func wallAhead(guard, posDelta Posistion) bool {
	if !possibleMove(guard, posDelta) {
		return false
	}
	x := guard.x
	y := guard.y
	if m[y][x+posDelta.x] == '#' ||
		m[y+posDelta.y][x] == '#' {
		return true
	}
	return false
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

func findGuard(s string) (bool, int) {
	if !strings.Contains(s, "^") {
		return false, 0
	}
	return true, strings.Index(s, "^")
}

func equalPos(dirA, dirB Posistion) bool {
	return dirA.x == dirB.x && dirA.y == dirB.y
}
