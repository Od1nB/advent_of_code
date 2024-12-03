package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var task1 int
	var task2 int

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

	list1, list2 := []int{}, []int{}
	m2 := map[int]int{}
	for scnr.Scan() {
		t := scnr.Text()
		val := strings.Split(t, " ")
		v1, _ := strconv.Atoi(val[0])
		v2, _ := strconv.Atoi(val[len(val)-1])
		m2[v2]++
		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for ind, v1 := range list1 {
		task1 += getAbsDist(v1, list2[ind])
		task2 += v1 * m2[v1]
	}

	fmt.Println("task1: ", task1) // 1222801
	fmt.Println("task2: ", task2) // 22545250
}

func getAbsDist(v1, v2 int) int {
	if v1 >= v2 {
		return v1 - v2
	} else {
		return v2 - v1
	}
}
