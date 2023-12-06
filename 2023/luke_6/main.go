package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	Time "time"

	"github.com/samber/lo"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile("[0-9]+")
	scnr := bufio.NewScanner(f)

	scnr.Scan()
	l1 := re.FindAllString(scnr.Text(), -1)
	time := lo.Map(l1, func(x string, _ int) int {
		n, _ := strconv.Atoi(x)
		return n
	})
	t2 := lo.Reduce(l1, func(n int, str string, _ int) int {
		i, _ := strconv.Atoi(str)
		newAgg, _ := strconv.Atoi(fmt.Sprintf("%d%d", n, i))
		return newAgg
	}, 0)

	scnr.Scan()
	l2 := re.FindAllString(scnr.Text(), -1)
	distance := lo.Map(l2, func(x string, _ int) int {
		n, _ := strconv.Atoi(x)
		return n
	})
	d2 := lo.Reduce(l2, func(n int, str string, _ int) int {
		i, _ := strconv.Atoi(str)
		newAgg, _ := strconv.Atoi(fmt.Sprintf("%d%d", n, i))
		return newAgg
	}, 0)

	ans := make([]int, len(time))
	for ind, _ := range ans {
		ans[ind] = 1
	}
	for ind, t := range time {
		ans[ind] = getPossibleSolutions(t, distance[ind])
	}
	task1 := 1
	for _, v := range ans {
		task1 = task1 * v
	}
	fmt.Println("task1 ", task1) //1413720
	startT := Time.Now()
	//fmt.Println("task2 threads", getPossibleSolutions(t2, d2)) //30565288
	fmt.Println("task2 threads", threadPossible(t2, d2, 2000)) //30565288
	fmt.Println(Time.Since(startT))
}

func getPossibleSolutions(t, d int) int {
	m := map[int]bool{}
	for i := 0; i < t; i++ {
		if l := getLenght(i+1, t); l > d {
			m[i+1] = true
		}
	}
	return len(m)
}

func threadPossible(t, d, threads int) int {
	wg := sync.WaitGroup{}
	a := make([]int, threads)
	for i := 0; i < threads; i++ {
		itr := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			threadThatShit(itr*(t/threads), (itr+1)*(t/threads), d, t, &a, itr)
		}()

	}
	wg.Wait()
	return lo.Reduce(a, func(agg int, i int, _ int) int {
		return agg + i
	}, 0)
}

func getLenght(wait, time int) int {
	return (time - wait) * wait
}

func threadThatShit(from, to, dist, time int, arr *[]int, ind int) {
	m := map[int]bool{}
	for i := from; i < to; i++ {
		if l := getLenght(i+1, time); l > dist {
			m[i+1] = true
		}
	}
	(*arr)[ind] = len(m)
}
