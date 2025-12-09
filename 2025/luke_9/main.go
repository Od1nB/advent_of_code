package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	task1       int
	task2       int
	points      = make([]point2D, 0, 500)
	greenTiles  = make(map[point2D]bool)
	insideCache = make(map[point2D]bool)
)

type point2D struct {
	X, Y int
}

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
		points = append(points, parsePoint(scnr.Text()))
	}

	task1 = biggestArea()
	buildGreenEdges()
	task2 = biggestAreaWithGreen()

	fmt.Println("task1: ", task1) // 4777824480
	fmt.Println("task2: ", task2) // 1542119040
}

func biggestArea() int {
	biggest := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := (intAbs(points[i].X-points[j].X) + 1) * (intAbs(points[i].Y-points[j].Y) + 1)
			if area <= biggest {
				continue
			}
			biggest = area
		}
	}
	return biggest
}

func parsePoint(s string) point2D {
	splits := strings.Split(s, ",")
	if len(splits) != 2 {
		panic("not 2 vars")
	}

	x, err := strconv.Atoi(splits[0])
	if err != nil {
		panic("x not number" + splits[0])
	}
	y, err := strconv.Atoi(splits[1])
	if err != nil {
		panic("y not number" + splits[1])
	}

	return point2D{x, y}
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func buildGreenEdges() {
	for _, p := range points {
		greenTiles[p] = true
	}

	for i := range points {
		next := (i + 1) % len(points)
		addLineBetween(points[i], points[next])
	}
}

func addLineBetween(p1, p2 point2D) {
	if p1.X == p2.X {
		start, end := p1.Y, p2.Y
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			greenTiles[point2D{p1.X, y}] = true
		}
	} else {
		start, end := p1.X, p2.X
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			greenTiles[point2D{x, p1.Y}] = true
		}
	}
}

func isInsidePolygon(p point2D) bool {
	crossings := 0
	rays := len(points)

	for i := range rays {
		j := (i + rays - 1) % rays

		start := points[j]
		end := points[i]

		crossY := (end.Y > p.Y) != (start.Y > p.Y)
		if !crossY {
			continue
		}

		slope := float64(end.X-start.X) / float64(end.Y-start.Y)
		crossX := float64(start.X) + slope*float64(p.Y-start.Y)

		if float64(p.X) < crossX {
			crossings++
		}
	}

	return crossings%2 == 1
}

func biggestAreaWithGreen() int {
	biggest := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			if !rectangleValid(points[i], points[j]) {
				continue
			}
			area := (intAbs(points[i].X-points[j].X) + 1) * (intAbs(points[i].Y-points[j].Y) + 1)
			if area > biggest {
				biggest = area
			}
		}
	}
	return biggest
}

func rectangleValid(p1, p2 point2D) bool {
	minX, maxX := p1.X, p2.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := p1.Y, p2.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	area := (maxX - minX + 1) * (maxY - minY + 1)

	if area <= 10000 {
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				if !isGreenOrRed(point2D{x, y}) {
					return false
				}
			}
		}
		return true
	}

	c3 := point2D{p1.X, p2.Y}
	c4 := point2D{p2.X, p1.Y}

	if !isGreenOrRed(c3) || !isGreenOrRed(c4) {
		return false
	}

	for x := minX; x <= maxX; x++ {
		if !isGreenOrRed(point2D{x, minY}) || !isGreenOrRed(point2D{x, maxY}) {
			return false
		}
	}
	for y := minY + 1; y < maxY; y++ {
		if !isGreenOrRed(point2D{minX, y}) || !isGreenOrRed(point2D{maxX, y}) {
			return false
		}
	}

	return true
}

func isGreenOrRed(p point2D) bool {
	if greenTiles[p] {
		return true
	}
	if result, ok := insideCache[p]; ok {
		return result
	}
	result := isInsidePolygon(p)
	insideCache[p] = result
	return result
}
