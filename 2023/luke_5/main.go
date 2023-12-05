package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Category int64

const (
	Seed Category = iota
	Soil
	Fertilizer
	Water
	Light
	Temperature
	Humidity
	Location
)

var (
	categoryMap = map[string]Category{
		"seed":        Seed,
		"soil":        Soil,
		"fertilizer":  Fertilizer,
		"water":       Water,
		"light":       Light,
		"temperature": Temperature,
		"humidity":    Humidity,
		"location":    Location,
	}
)

func (c Category) String() string {
	return [...]string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}[c]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	scnr.Scan()

	seeds := strings.Split(strings.Trim(strings.Split(scnr.Text(), ":")[1], " "), " ")
	scnr.Scan()
	fmt.Println(seeds)
	ind := 0
	maps := make([]*catMap, 7)
	var from Category
	var to Category
	for scnr.Scan() {
		t := scnr.Text()
		switch {
		case len(strings.Split(t, ":")) > 1:
			from = categoryMap[strings.Split(t, "-")[0]]
			to = categoryMap[strings.Split(strings.Split(t, "-")[2], " ")[0]]
			fmt.Println(from, " ", to)
			maps[ind] = &catMap{
				from:    from,
				to:      to,
				fromMap: map[int]int{},
				toMap:   map[int]int{},
			}
		case len(strings.Split(t, " ")) > 2:
			i := getInts(t)
			maps[ind].fromMap[i[1]] = i[2]
			maps[ind].toMap[i[0]] = i[2]
		default:
			ind += 1
		}
	}

}

type catMap struct {
	from    Category
	to      Category
	fromMap map[int]int
	toMap   map[int]int
}

func isInRange(inp int, start int, r int) bool {
	return (inp >= start && inp < start+r)
}

func getInts(s string) []int {
	stringNum := strings.Split(s, " ")
	var n = make([]int, 3)
	dest, _ := strconv.Atoi(stringNum[0])
	from, _ := strconv.Atoi(stringNum[1])
	r, _ := strconv.Atoi(stringNum[2])
	n[0], n[1], n[2] = dest, from, r
	return n
}
