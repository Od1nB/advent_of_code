package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	task1 := 0
	task2 := 0
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)

	for scnr.Scan() {
		t := scnr.Text()
		g := splitLine(t)
		setInit(g)
		fillSets(g)
		if task1Condition(g) {
			task1 += g.id
		}
		task2 += powerOfSmallest(g)
	}
	fmt.Println("Task1 answer: ", task1) // 2176
	fmt.Println("Task2 answer: ", task2) // 63700

}

func task1Condition(g *Game) bool {
	fmt.Println(g.str)
	red, green, blue := 12, 13, 14
	for _, set := range g.sets {
		g_red := set.red
		g_green := set.green
		g_blue := set.blue
		if red < g_red || green < g_green || blue < g_blue {
			return false
		}
	}
	return true
}

func powerOfSmallest(g *Game) int {
	red, blue, green := 0, 0, 0
	for _, set := range g.sets {
		if set.red > red {
			red = set.red
		}
		if set.blue > blue {
			blue = set.blue
		}
		if set.green > green {
			green = set.green
		}
	}
	return red * blue * green
}

func splitLine(s string) *Game {
	split := strings.Split(s, ":")
	g := strings.Split(split[0], " ")
	i, _ := strconv.Atoi(g[1])
	return &Game{id: i, str: split[1]}
}

func setInit(g *Game) {
	setsSplit := strings.Split(g.str, ";")
	sets := make([]*Set, len(setsSplit))
	for i, s := range setsSplit {
		sets[i] = &Set{str: s}
	}
	g.sets = sets
}

func fillSets(g *Game) {
	for _, set := range g.sets {
		fillSet(set)
	}
}

func fillSet(s *Set) {
	colors := strings.Split(s.str, ",")
	for _, c := range colors {
		trim_c := strings.Trim(c, " ")
		colorFill(s, strings.Split(trim_c, " "))
	}
}

type Game struct {
	id   int
	str  string
	sets []*Set
}

type Set struct {
	str   string
	blue  int
	green int
	red   int
}

func colorFill(s *Set, str []string) {
	num, _ := strconv.Atoi(str[0])
	switch str[1] {
	case "green":
		s.green = num
	case "red":
		s.red = num
	case "blue":
		s.blue = num
	}
}
