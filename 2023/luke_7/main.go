package main

import (
	"bufio"
	"fmt"
	camel "luke_7/camelgame"
	"os"
	"sort"
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
	a := []camel.Hand{}
	tot, tot2 := 0, 0
	for scnr.Scan() {
		t := strings.Split(scnr.Text(), " ")
		b, _ := strconv.Atoi(t[1])
		a = append(a, camel.Hand{Cards: string(t[0]), Bid: b})
	}
	sort.Slice(a, camel.Less(a, camel.CardValue, camel.CardComp))
	for ind, hand := range a {
		tot += (ind + 1) * hand.Bid
	}
	sort.Slice(a, camel.Less(a, camel.JkrCardValue, camel.CardCompJoker))
	for ind, hand := range a {
		tot2 += (ind + 1) * hand.Bid
	}
	fmt.Println("task1 ", tot)
	fmt.Println("task2 ", tot2)

}
