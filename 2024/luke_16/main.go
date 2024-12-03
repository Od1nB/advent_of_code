package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		fmt.Println(t)
	}

	fmt.Println("task1: ", task1)
	fmt.Println("task2: ", task2)
}
