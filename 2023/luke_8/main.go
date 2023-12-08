package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Direction string

const (
	LEFT  Direction = "L"
	RIGHT Direction = "R"
)

type Entry struct {
	Position string
	Left     string
	Right    string
}

func main() {
	startTime := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	scnr.Scan()

	dir := scnr.Text()
	entryMap := map[string]Entry{}
	position := ""
	for scnr.Scan() {
		if scnr.Text() == "" {
			continue
		}
		t := scnr.Text()
		entryMap[strings.Split(t, " ")[0]] = Entry{
			Position: strings.Split(t, " ")[0],
			Left:     strings.Split(strings.Split(t, "(")[1], ",")[0],
			Right:    strings.TrimRight(strings.Trim(strings.Split(strings.Split(t, "(")[1], ",")[1], " "), ")"),
		}
		if position == "" {
			position = strings.Split(t, " ")[0]
		}
	}
	followMap := func(from, to string) int {
		shortest := []int{}
		for e := range entryMap {
			if !strings.HasSuffix(e, from) {
				continue
			}
			steps := 0
			for !strings.HasSuffix(e, to) {
				switch Direction(string(dir[steps%len(dir)])) {
				case LEFT:
					e = entryMap[e].Left
				case RIGHT:
					e = entryMap[e].Right
				}
				steps++
			}
			shortest = append(shortest, steps)

		}
		return lcm(shortest...)
	}

	fmt.Print("Parsing time: ", time.Since(startTime), "\n\n")

	startTime = time.Now()
	fmt.Println("-- task 1 --")
	fmt.Println("Answer: ", followMap("AAA", "ZZZ"))
	fmt.Print("Time: ", time.Since(startTime), "\n\n")

	startTime = time.Now()
	fmt.Println("-- task 2 --")
	fmt.Println("Answer: ", followMap("A", "Z"))
	fmt.Println("Time: ", time.Since(startTime))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(ints ...int) int {
	if len(ints) == 1 {
		return ints[0]
	} else if len(ints) > 2 {
		return lcm(ints[0], lcm(ints[1:]...))
	}
	return ints[0] * ints[1] / gcd(ints[0], ints[1])
}
