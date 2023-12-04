package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, task2, ind := 0, 0, 0
	cards := make([]*Card, 192)
	m := map[int]int{}
	for scnr.Scan() {
		t := scnr.Text()
		c := makeCard(t, ind)
		cards[ind] = c
		m[ind] = 1
		task1 += c.scoreTask1()
		ind += 1
	}

	for ind, c := range cards {
		copies := m[ind]
		for i := 0; i < copies; i++ {
			AddCopies(m, c)
		}
	}
	for _, v := range m {
		task2 += v
	}

	fmt.Println("task1: ", task1) // 21213
	fmt.Println("task2: ", task2) // 8549735
}

func AddCopies(m map[int]int, c *Card) {
	matches := c.getMatches()
	if matches != 0 {
		for i := c.id + 1; i <= c.id+matches; i++ {
			m[i] += 1
		}
	}
}

type Card struct {
	id      int
	str     string
	input   *Side
	winning *Side
	points  int
}

func makeCard(s string, i int) *Card {
	split1 := strings.Split(s, ":")
	sides := strings.Split(split1[1], "|")
	c := &Card{id: i, str: split1[1], points: 0, input: &Side{str: sides[0], numsMap: map[int]int{}, nums: []int{}}, winning: &Side{str: sides[1], numsMap: map[int]int{}, nums: []int{}}}
	c.input.fillSide()
	c.winning.fillSide()
	return c
}

type Side struct {
	str     string
	numsMap map[int]int
	nums    []int
}

func (c *Card) scoreTask1() int {
	for n := range c.winning.numsMap {
		if _, ok := c.input.numsMap[n]; ok {
			c.doublePoints()
		}
	}
	return c.points
}
func (c *Card) getMatches() int {
	out := 0
	for n, _ := range c.winning.numsMap {
		if _, ok := c.input.numsMap[n]; ok {
			out += 1
		}
	}
	return out
}

func (s *Side) fillSide() {
	s.str = strings.Trim(s.str, " ")
	nums := strings.Split(s.str, " ")
	for ind, n := range nums {
		if n != "" {
			num, _ := strconv.Atoi(strings.Trim(n, " "))
			s.numsMap[num] = ind
			s.nums = append(s.nums, num)
		}
	}
}

func (c *Card) doublePoints() {
	if c.points == 0 {
		c.points = 1
	} else {
		c.points = c.points * 2
	}
}
