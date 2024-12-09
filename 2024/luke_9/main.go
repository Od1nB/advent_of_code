package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	var disk []int

	for scnr.Scan() {
		t := scnr.Text()
		isFile := true
		fileID := 0
		for _, v := range t {
			num, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			if isFile {
				disk = append(disk, slices.Repeat([]int{fileID}, num)...)
				fileID++
				isFile = !isFile
			} else {
				disk = append(disk, slices.Repeat([]int{-1}, num)...)
				isFile = !isFile
			}
		}
	}

	disk2 := make([]int, len(disk))
	copy(disk2, disk)

	disk = orderDisk(disk)
	disk2 = orderFiles(disk2)

	task1 += calcChecksum(disk)
	task2 += calcChecksum(disk2)

	fmt.Println("task1: ", task1) // 6446899523367
	fmt.Println("task2: ", task2) // 6478232739671
}

func lastFileInd(d []int) int {
	for ind, num := range slices.Backward(d) {
		if num != -1 {
			return ind
		}
	}
	return -1
}

func orderFiles(d []int) []int {
	fileID := d[len(d)-1] // input and example always end in file (odd)
	for fileID > 0 {
		length := fileIDLength(d, fileID)
		idInd := slices.Index(d, fileID)
		freeInd := findFreeIndex(d, length, idInd)
		if freeInd > -1 {
			for offset := range length {
				d[freeInd+offset] = d[idInd+offset]
				d[idInd+offset] = -1
			}
		}
		fileID--
	}

	return d
}

func orderDisk(d []int) []int {
	lastFile := lastFileInd(d)
	fmt.Println(lastFile)
	for i := 0; i < lastFile; i++ {
		if d[i] == -1 {
			d[i] = d[lastFile]
			d[lastFile] = -1
			lastFile = lastFileInd(d)
		}
	}
	return d
}

func calcChecksum(disk []int) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] > -1 {
			sum += i * disk[i]
		}
	}
	return sum
}

func fileIDLength(d []int, id int) int {
	length := 0
	var seen bool
	var done bool

	for ind := 0; ind < len(d); ind++ {
		if seen && done {
			return length
		}
		if d[ind] == id {
			length++
			seen = true
		}

		if seen && d[ind] != id {
			done = true
		}

	}
	return length
}

func findFreeIndex(disk []int, length int, maxID int) int {
	free := slices.Index(disk, -1)
	if maxID <= 0 || free == -1 {
		return -1
	}
	count := 0

	for i := free; i < len(disk) && i < maxID; i++ {
		if disk[i] == -1 {
			count++
			if count == length {
				return i - length + 1
			}
		} else {
			count = 0
		}
	}

	return -1
}
