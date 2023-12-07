package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Composition int

const (
	HighCard Composition = iota
	OnePair
	TwoPair
	ThreeofaKind
	FullHouse
	FourofaKind
	FiveofaKind
)

func (c Composition) String() string {
	switch c {
	case 0:
		return "HighCard"
	case 1:
		return "OnePair"
	case 2:
		return "TwoPair"
	case 3:
		return "ThreeofaKind"
	case 4:
		return "FullHouse"
	case 5:
		return "FourofaKind"
	case 6:
		return "FiveofaKind"
	}
	return ""
}

var rankToVal = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}
var rankwJoker = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type Hand struct {
	Cards string
	Bid   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	a := []Hand{}
	tot := 0
	for scnr.Scan() {
		t := strings.Split(scnr.Text(), " ")
		b, _ := strconv.Atoi(t[1])
		a = append(a, Hand{Cards: string(t[0]), Bid: b})
	}
	sort.Slice(a, Less(a))
	for ind, hand := range a {
		tot += (ind + 1) * hand.Bid
	}
	tot2 := 0
	sort.Slice(a, LesswJkr(a))
	for ind, hand := range a {
		tot2 += (ind + 1) * hand.Bid
	}
	fmt.Println("task1 ", tot)
	fmt.Println("task2 ", tot2)

}

func (h Hand) CardComp() Composition {
	m := map[string]int{}
	for _, v := range h.Cards {
		m[string(v)]++
	}
	switch {
	case len(m) == 1:
		return FiveofaKind
	case len(m) == 2:
		for _, v := range m {
			if v == 4 {
				return FourofaKind
			}
		}
		return FullHouse
	case len(m) == 3:
		for _, v := range m {
			if v == 3 {
				return ThreeofaKind
			}
		}
		return TwoPair
	case len(m) == 4:
		return OnePair
	default:
		return HighCard
	}
}
func (h Hand) CardCompJoker() Composition {
	m := map[string]int{}
	for _, v := range h.Cards {
		m[string(v)]++
	}
	switch {
	case len(m) == 1:
		return FiveofaKind
	case len(m) == 2:
		if _, jkr := m["J"]; jkr {
			return FiveofaKind
		}
		for _, v := range m {
			if v == 4 {
				return FourofaKind
			}
		}
		return FullHouse
	case len(m) == 3:
		numJkr, jkr := m["J"]
		for _, v := range m {
			if v == 3 && jkr {
				return FourofaKind
			} else if v == 3 && !jkr {
				return ThreeofaKind
			}
		}
		if jkr {
			if numJkr == 1 {
				return FullHouse
			}
			return FourofaKind
		}
		return TwoPair
	case len(m) == 4:
		if _, jkr := m["J"]; jkr {
			return ThreeofaKind
		}
		return OnePair
	default:
		if _, jkr := m["J"]; jkr {
			return OnePair
		}
		return HighCard
	}
}

func cardValue(s string) int {
	return rankToVal[s]
}

func jkrCardValue(s string) int {
	return rankwJoker[s]
}

func highCard(h, j Hand) bool {
	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i] == j.Cards[i] {
			continue
		}
		return cardValue(string(h.Cards[i])) < cardValue(string(j.Cards[i]))
	}
	return false
}

func jkrHighCard(h, j Hand) bool {
	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i] == j.Cards[i] {
			continue
		}
		return jkrCardValue(string(h.Cards[i])) < jkrCardValue(string(j.Cards[i]))
	}
	return false
}

func Sort(less func(hands []Hand) func(i int, j int) bool) func(hands []Hand) []Hand {
	return func(hands []Hand) []Hand {
		sort.Slice(hands, less(hands))
		return hands
	}
}

func Less(hands []Hand) func(i, j int) bool {
	return func(i, j int) bool {
		c1, c2 := hands[i].CardComp(), hands[j].CardComp()
		if c1 != c2 {
			return c1 < c2
		}
		return highCard(hands[i], hands[j])
	}
}

func LesswJkr(hands []Hand) func(i, j int) bool {
	return func(i, j int) bool {
		c1, c2 := hands[i].CardCompJoker(), hands[j].CardCompJoker()
		if c1 != c2 {
			return c1 < c2
		}
		return jkrHighCard(hands[i], hands[j])
	}
}
