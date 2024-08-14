package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	length  int
	str     string
	letters string
	sign    string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, task2 := 0, 0
	vals := []string{}
	lenses := [256][]Lens{}

	for scnr.Scan() {
		t := scnr.Text()
		vals = strings.Split(t, ",")
	}
	re := regexp.MustCompile(`(\w+)([-=])(\d*)`)
	for _, v := range vals {
		task1 += HashThatShit(v)
		some := re.FindAllStringSubmatch(v, -1)
		//fmt.Println(some[0])
		lens := &Lens{str: some[0][0], letters: some[0][1], sign: some[0][2]}
		if num, err := strconv.Atoi(some[0][3]); err == nil {
			lens.length = num
		}
		currBox := &lenses[HashThatShit(lens.letters)]
		index := slices.IndexFunc(*currBox, func(l Lens) bool { return l.letters == lens.letters })
		if lens.sign == "-" && index != -1 {
			*currBox = slices.Delete(*currBox, index, index+1)
		} else if lens.sign == "=" && index != -1 {
			(*currBox)[index] = *lens
		} else if lens.sign == "=" {
			*currBox = append(*currBox, *lens)
		}
		//lens := &Lens{length: v}
	}
	for i, b := range lenses {
		for j, l := range b {
			task2 += (i + 1) * (j + 1) * l.length
		}
	}
	//fmt.Println(HashThatShit("ot"))
	fmt.Println("task1: ", task1)
	fmt.Println("task2: ", task2)
}

func HashThatShit(s string) int {
	value := 0
	for i := range s {
		value += int(s[i])
		value *= 17
		value = value % 256
	}
	return value
}
