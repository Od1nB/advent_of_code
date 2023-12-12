package main

import (
	"bufio"
	"fmt"
	"os"
)

type pipePart struct {
	typ    string
	x      int
	y      int
	xy     string
	n      *pipePart
	w      *pipePart
	e      *pipePart
	s      *pipePart
	inLoop bool
}

type pipeMaze struct {
	pipes map[string]*pipePart
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, task2 := 0, 0
	grid := []string{}
	for scnr.Scan() {
		t := scnr.Text()
		grid = append(grid, t)
	}
	pm := fillMaze(grid)
	task1, task2 = pm.moveThrough()/2, pm.calcInsideArea(grid)

	fmt.Println("task1: ", task1) //6768
	fmt.Println("task2: ", task2) //351
}

func (pm *pipeMaze) moveThrough() int {
	var start *pipePart
	for _, pipe := range pm.pipes {
		if pipe.typ == "S" {
			start = pipe
			break
		}
	}
	next := start.findNext(start)

	steps := 1
	last := start
	current := next

	for current != start {
		current.inLoop = true
		steps++
		prev := current
		current = current.findNext(last)
		last = prev
	}

	return steps
}

func (pm pipeMaze) calcInsideArea(grid []string) int {

	pm.replaceStart()
	score := 0
	for k, v := range pm.pipes {
		if !v.inLoop {
			pm.pipes[k].typ = "."
		}
	}
	for y, line := range grid {
		linesscore := 0
		last := ""
		lines := 0
		tempscore := 0
		for x := range line {
			if pipe := pm.pipes[intsToCord(x, y)]; pipe.inLoop {
				switch {
				case pipe.typ == "|":
					lines++
					last = ""
				case pipe.typ == "7" && last == "L":
					lines++
					last = ""
				case pipe.typ == "J" && last == "F":
					lines++
					last = ""
				case pipe.typ == "F" || pipe.typ == "L":
					last = pipe.typ
				}
			} else if lines%2 == 0 && lines > 0 && tempscore != 0 {
				linesscore += tempscore
				tempscore = 0

			} else if lines%2 != 0 && pipe.typ == "." {
				tempscore++
			}
		}
		score += linesscore
	}
	return score
}

// Made this to output the graph to look at it
/*
for k, v := range pm.pipes {
	if !v.inLoop {
		pm.pipes[k].typ = "."
	}
}
out := make([]string, 140)
var sb strings.Builder
for ind, _ := range out {
	sb.Grow(140)
	for i := 0; i < 140; i++ {
		sb.WriteString(pm.pipes[strconv.Itoa(i)+":"+strconv.Itoa(ind)].typ)
	}
	out[ind] = sb.String()
	sb.Reset()
}
f, _ := os.Create("out.txt")
defer f.Close()
for _, v2 := range out {
	fmt.Fprintln(f, v2)
}
*/

func (pipe *pipePart) findNext(last *pipePart) *pipePart {
	switch {
	case pipe.n != nil && pipe.n != last:
		return pipe.n
	case pipe.s != nil && pipe.s != last:
		return pipe.s
	case pipe.w != nil && pipe.w != last:
		return pipe.w
	case pipe.e != nil && pipe.e != last:
		return pipe.e
	default:
		fmt.Println("bad stuff")
		return nil
	}
}

func fillMaze(input []string) *pipeMaze {
	pm := pipeMaze{pipes: map[string]*pipePart{}}
	for y, line := range input {
		for x, typ := range line {
			pm.pipes[intsToCord(x, y)] = &pipePart{
				typ:    string(typ),
				x:      x,
				y:      y,
				xy:     intsToCord(x, y),
				inLoop: string(typ) == "S",
			}
		}
	}
	for _, pipe := range pm.pipes {
		switch pipe.typ {
		case "|":
			pipe.n = pm.pipes[intsToCord(pipe.x, pipe.y-1)]
			pipe.s = pm.pipes[intsToCord(pipe.x, pipe.y+1)]
		case "-":
			pipe.w = pm.pipes[intsToCord(pipe.x-1, pipe.y)]
			pipe.e = pm.pipes[intsToCord(pipe.x+1, pipe.y)]
		case "L":
			pipe.n = pm.pipes[intsToCord(pipe.x, pipe.y-1)]
			pipe.e = pm.pipes[intsToCord(pipe.x+1, pipe.y)]
		case "J":
			pipe.n = pm.pipes[intsToCord(pipe.x, pipe.y-1)]
			pipe.w = pm.pipes[intsToCord(pipe.x-1, pipe.y)]
		case "7":
			pipe.s = pm.pipes[intsToCord(pipe.x, pipe.y+1)]
			pipe.w = pm.pipes[intsToCord(pipe.x-1, pipe.y)]
		case "F":
			pipe.s = pm.pipes[intsToCord(pipe.x, pipe.y+1)]
			pipe.e = pm.pipes[intsToCord(pipe.x+1, pipe.y)]
		}
		if pipe.n != nil && pipe.n.typ == "S" {
			pipe.inLoop = true
			pipe.n.s = pipe
		}
		if pipe.s != nil && pipe.s.typ == "S" {
			pipe.inLoop = true
			pipe.s.n = pipe
		}
		if pipe.e != nil && pipe.e.typ == "S" {
			pipe.inLoop = true
			pipe.e.w = pipe
		}
		if pipe.w != nil && pipe.w.typ == "S" {
			pipe.inLoop = true
			pipe.w.e = pipe
		}
	}
	return &pm
}

func (pm pipeMaze) replaceStart() {
	for _, pipe := range pm.pipes {
		if pipe.typ == "S" {
			switch {
			case pipe.w != nil && pipe.s != nil:
				pipe.typ = "/"
			case pipe.e != nil && pipe.s != nil:
				pipe.typ = "F"
			case pipe.e != nil && pipe.w != nil:
				pipe.typ = "-"
			case pipe.n != nil && pipe.w != nil:
				pipe.typ = "J"
			case pipe.n != nil && pipe.s != nil:
				pipe.typ = "|"
			}
		}
	}
}

func intsToCord(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
