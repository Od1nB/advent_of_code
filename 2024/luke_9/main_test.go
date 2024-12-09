package main

import (
	"testing"
)

var (
	parsed = []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7,
		7, 7, -1, 8, 8, 8, 8, 9, 9}
)

func TestPart1(t *testing.T) {
	ordered := orderDisk(parsed)
	score := calcChecksum(ordered)
	if score != 1928 {
		t.Errorf("Wrong checksum, got: %d, expected: %d", score, 1928)
	}
}

func TestLastFile(t *testing.T) {
	arr := []int{1, 2, 3, -1}
	ind := lastFileInd(arr)
	if ind != 2 {
		t.Errorf("got index %d, but expected %d", ind, 2)
	}
}

func TestPart2(t *testing.T) {
	ordered := orderFiles(parsed)
	score := calcChecksum(ordered)
	if score != 2858 {
		t.Errorf("Wrong checksum, got: %d, expected: %d", score, 2858)
	}
}

func TestFindFreeIndex(t *testing.T) {
	ind := findFreeIndex(parsed, 2, 41) //9
	if ind != 2 {
		t.Errorf("got %d, but expected %d", ind, 2)
	}
}
