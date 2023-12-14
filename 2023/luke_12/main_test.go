package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfArr(t *testing.T) {
	assert.Equal(t, 1, numberOfArr("???.###", []int{1, 1, 3}))
	assert.Equal(t, 4, numberOfArr(".??..??...?##.", []int{1, 1, 3}))
	assert.Equal(t, 1, numberOfArr("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}))
	assert.Equal(t, 1, numberOfArr("????.#...#...", []int{4, 1, 1}))
	assert.Equal(t, 4, numberOfArr("????.######..#####.", []int{1, 6, 5}))
	assert.Equal(t, 10, numberOfArr("?###????????", []int{3, 2, 1}))
	assert.Equal(t, 1, numberOfArr("???.###????.###????.###????.###????.###", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}))
	assert.Equal(t, 16384, numberOfArr(".??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}))
}

func TestKeyHash(t *testing.T) {
	h1, h2, h3 := keyHash("123", []int{1, 2, 3}), keyHash("123", []int{1, 2, 3}), keyHash("123", []int{1, 0, 3})
	assert.Equal(t, 0, strings.Compare(h2, h1))
	assert.Equal(t, true, strings.Compare(h2, h3) != 0)
}
