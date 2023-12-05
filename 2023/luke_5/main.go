package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	scnr.Scan()

	seeds := strings.Split(strings.Trim(strings.Split(scnr.Text(), ":")[1], " "), " ")
	fmt.Println(seeds)
	for scnr.Scan() {

	}

}
