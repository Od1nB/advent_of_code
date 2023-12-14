package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	lo "github.com/samber/lo"
)

type Direction int

const (
	FORWARD Direction = iota
	BACKWARDS
)

type Record struct {
	str    string
	groups []int
}

var cache sync.Map

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	task1, task2 := 0, 0
	records_p1 := []*Record{}
	records_p2 := []*Record{}
	var sb_recordmultiplier strings.Builder

	for scnr.Scan() {
		t := scnr.Text()
		split := strings.Split(t, " ")
		record := &Record{split[0], lo.Map(strings.Split(split[1], ","), func(s string, _ int) int {
			num, _ := strconv.Atoi(s)
			return num
		})}
		records_p1 = append(records_p1, record)

		sb_recordmultiplier.Grow(5 + len(record.str)*5)
		for i := 0; i < 5; i++ {
			sb_recordmultiplier.WriteString(record.str)
			if i < 4 {
				sb_recordmultiplier.WriteString("?")
			}
		}
		records_p2 = append(records_p2, &Record{sb_recordmultiplier.String(), expandSlice(record.groups, 5)})
		sb_recordmultiplier.Reset()
	}

	t1 := time.Now()
	for _, rec := range records_p1 {
		task1 += numberOfArr(rec.str, rec.groups)
	}
	fmt.Println("time taken: ", time.Since(t1))
	fmt.Println("task1: ", task1) //7771

	wg := sync.WaitGroup{}
	t2 := time.Now()
	for _, record := range records_p2 {
		wg.Add(1)
		s := record
		go func() {
			defer wg.Done()
			task2 += numberOfArr(s.str, s.groups)
		}()
	}
	wg.Wait()

	fmt.Println("time taken: ", time.Since(t2))
	fmt.Println("task2: ", task2) // 10861030975833
}

func numberOfArr(str string, ints []int) int {
	key := keyHash(str, ints)
	if ans, ok := cache.Load(key); ok {
		return ans.(int)
	}

	if len(str) == 0 {
		if len(ints) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if strings.HasPrefix(str, "?") {
		return numberOfArr(strings.Replace(str, "?", ".", 1), ints) + numberOfArr(strings.Replace(str, "?", "#", 1), ints)
	}
	if strings.HasPrefix(str, ".") {
		intrmdt := numberOfArr(strings.TrimPrefix(str, "."), ints)
		cache.Store(key, intrmdt)
		return intrmdt
	}

	if strings.HasPrefix(str, "#") {
		if len(ints) == 0 {
			cache.Store(key, 0)
			return 0
		}
		if len(str) < ints[0] {
			cache.Store(key, 0)
			return 0
		}
		if strings.Contains(str[0:ints[0]], ".") {
			cache.Store(key, 0)
			return 0
		}
		if len(ints) > 1 {
			if len(str) < ints[0]+1 || string(str[ints[0]]) == "#" {
				cache.Store(key, 0)
				return 0
			}
			intrmdt := numberOfArr(str[ints[0]+1:], ints[1:])
			cache.Store(key, intrmdt)
			return intrmdt
		} else {
			intrmdt := numberOfArr(str[ints[0]:], ints[1:])
			cache.Store(key, intrmdt)
			return intrmdt
		}
	}
	cache.Store(key, 0)
	return 0
}

func expandSlice(slice []int, n int) []int {
	newSlice := make([]int, len(slice)*n)

	for i := 0; i < n; i++ {
		copy(newSlice[i*len(slice):(i+1)*len(slice)], slice)
	}

	return newSlice
}

func keyHash(str string, slice []int) string {
	h := sha256.New()
	h.Write([]byte(str))

	for _, v := range slice {
		h.Write([]byte(string(rune(v))))
	}
	return string(h.Sum(nil)[:])
}
