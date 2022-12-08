package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(getNum())
}

func getNum() int {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	temp, biggest := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if temp > biggest {
				biggest = temp
			}
			temp = 0
		}
		if num, ok := strconv.Atoi(line); ok == nil {
			temp += num
		}
	}
	return biggest
}
