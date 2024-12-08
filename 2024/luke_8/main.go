package main

import (
	"bufio"
	"fmt"
	"os"
)

type posistion struct {
	x     int
	y     int
	value rune
}

var m [][]posistion

func main() {
	freqs := map[rune][]posistion{}
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
		m = append(m, row)
		putFreqs(freqs, row)
	}

	antiNodesP1, antiNodesP2 := map[posistion]bool{}, map[posistion]bool{}
	for _, freqPoss := range freqs {
		for _, f1 := range freqPoss {
			for _, f2 := range freqPoss {
				if f2 == f1 {
					continue
				}
				if p := f2.Add(f2.Sub(f1)); p.within(m) {
					antiNodesP1[p] = true
				}
				for pdelta := f2.Sub(f1); f2.within(m); f2 = f2.Add(pdelta) {
					antiNodesP2[f2] = true
				}
			}
		}

	}
	antiNodesP2 = cleanMap(antiNodesP2)
	task1 += len(antiNodesP1)
	task2 += len(antiNodesP2)

	fmt.Println("task1: ", task1) // 295
	fmt.Println("task2: ", task2) // 1034
}

func makePosRow(s string, y int) []posistion {
	row := make([]posistion, len(s))
	for x, v := range s {
		row[x].value = v
		row[x].x = x
		row[x].y = y
	}
	return row
}

func putFreqs(freqMap map[rune][]posistion, row []posistion) {
	for _, pos := range row {
		if pos.value != '.' {
			if len(freqMap[pos.value]) == 0 {
				freqMap[pos.value] = []posistion{pos}
				continue
			}
			freqMap[pos.value] = append(freqMap[pos.value], pos)
		}
	}
}

func (p posistion) Add(delta posistion) posistion {
	var cpPoint posistion
	cpPoint.x = p.x + delta.x
	cpPoint.y = p.y + delta.y
	return cpPoint
}

func (p posistion) Sub(delta posistion) posistion {
	var cpPoint posistion
	cpPoint.x = p.x - delta.x
	cpPoint.y = p.y - delta.y
	return cpPoint
}

func cleanMap(m map[posistion]bool) map[posistion]bool {
	seen := []posistion{}
	for k := range m {
		for _, v := range seen {
			if equalPos(k, v) {
				delete(m, k)
			}
		}
		seen = append(seen, k)
	}
	return m
}

func (p posistion) within(m [][]posistion) bool {
	val := p.y >= 0 && p.y < len(m) && p.x >= 0 && p.x < len(m[p.y])
	return val
}

func equalPos(pA, pB posistion) bool {
	return pA.x == pB.x && pA.y == pB.y
}
