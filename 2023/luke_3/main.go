package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, ind, task2 := 0, 0, 0

	slide := &slidingWindow{}
	slide.addRow(&row{str: "", lenght: 0, indToVal: map[int]int{}, numbers: []*numba{}, symbols: []*symbol{}})
	slide.addRow(&row{str: "", lenght: 0, indToVal: map[int]int{}, numbers: []*numba{}, symbols: []*symbol{}})
	for scnr.Scan() {
		t := scnr.Text()
		r := &row{ind: ind, str: t, lenght: len(t), indToVal: map[int]int{}}
		r.findNumbers()
		r.fillMap()
		slide.addRow(r)
		if slide.top != nil && slide.middle != nil && slide.bottom != nil {
			task1 += slide.calcTask1()
			task2 += slide.calcGear()
		}
		ind++
	}
	slide.addRow(&row{str: "", lenght: 0, indToVal: map[int]int{}, numbers: []*numba{}, symbols: []*symbol{}})
	task1 += slide.calcTask1()
	task2 += slide.calcGear()

	fmt.Println("task1: ", task1) // 521601
	fmt.Println("task2: ", task2) // 80694070
}

func (w *slidingWindow) calcGear() int {
	gearScore := 0
	for _, s := range w.middle.symbols {
		if s.str == "*" {
			m := w.findAdjecent(s.ind)
			if len(m) > 1 {
				gear := 1
				for k := range m {
					gear = gear * k
				}
				gearScore += gear
			}
		}
	}
	return gearScore
}

func (w *slidingWindow) calcTask1() int {
	score := 0
	for _, sym := range w.middle.symbols {
		m := w.findAdjecent(sym.ind)
		for val := range m {
			score += val
		}
	}
	return score
}

func (w *slidingWindow) findAdjecent(index int) map[int]int {
	m := map[int]int{}
	if topleft, tl := w.top.indToVal[index-1]; tl {
		m[topleft] = index
	}
	if top, t := w.top.indToVal[index]; t {
		m[top] = index
	}
	if topright, tr := w.top.indToVal[index+1]; tr {
		m[topright] = index
	}
	if middleleft, ml := w.middle.indToVal[index-1]; ml {
		m[middleleft] = index
	}
	if middleright, mr := w.middle.indToVal[index+1]; mr {
		m[middleright] = index
	}
	if bottomleft, bl := w.bottom.indToVal[index-1]; bl {
		m[bottomleft] = index
	}
	if bottom, b := w.bottom.indToVal[index]; b {
		m[bottom] = index
	}
	if bottomright, br := w.bottom.indToVal[index+1]; br {
		m[bottomright] = index
	}
	return m
}

type row struct {
	ind        int
	str        string
	lenght     int
	numbers    []*numba
	numbersInd []int
	symbols    []*symbol
	symbolsInd []int
	indToVal   map[int]int
}

func (r *row) findNumbers() {
	lastNum := false
	for ind, sym := range r.str {
		if string(sym) != "." {
			if num, err := strconv.Atoi(string(sym)); lastNum && err == nil {
				c_numba := r.numbers[len(r.numbers)-1]
				new_num, _ := strconv.Atoi(fmt.Sprintf("%d%d", c_numba.num, num))
				c_numba.num = new_num
				c_numba.ind = append(c_numba.ind, ind)
				r.numbersInd = append(r.numbersInd, ind)
				r.numbers[len(r.numbers)-1] = c_numba
				lastNum = true
			} else if num, err := strconv.Atoi(string(sym)); err == nil && !lastNum {
				r.numbers = append(r.numbers, &numba{num: num, ind: []int{ind}})
				r.numbersInd = append(r.numbersInd, ind)
				lastNum = true
			} else {
				r.symbols = append(r.symbols, &symbol{str: string(sym), ind: ind})
				r.symbolsInd = append(r.symbolsInd, ind)
				lastNum = false
			}
		} else {
			lastNum = false
		}
	}
}

func (r *row) fillMap() {
	for _, n := range r.numbers {
		for _, ind := range n.ind {
			r.indToVal[ind] = n.num
		}
	}
}

type symbol struct {
	str string
	ind int
}

type numba struct {
	num int
	ind []int
}

type slidingWindow struct {
	top    *row
	middle *row
	bottom *row
}

func (w *slidingWindow) addRow(r *row) {
	switch {
	case w.top == nil:
		w.top = r
	case w.middle == nil:
		w.middle = r
	case w.bottom == nil:
		w.bottom = r
	default:
		w.top = w.middle
		w.middle = w.bottom
		w.bottom = r
	}
}
