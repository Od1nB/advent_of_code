package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	lo "github.com/samber/lo"
)

type Entry struct {
	id      int
	history []int
	extras  [][]int
}

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	entries := []Entry{}
	ind := 0
	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		entries = append(entries, Entry{
			id: ind,
			history: lo.Map(strings.Split(scnr.Text(), " "), func(s string, _ int) int {
				n, _ := strconv.Atoi(s)
				return n
			}),
		})
		ind++
	}
}
