package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	task1  int
	task2  int
	points = make([]Point3D, 0, 1000)
	pairs  = make([]pair, 0, 500000)
)

type Point3D struct {
	X, Y, Z int
}

type pair struct {
	i, j     int
	distance float64
}

// UnionFind is an implementation of a disjointed-set datastructure
// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type UnionFind struct {
	parent    []int
	connected []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	connected := make([]int, n)
	for i := range n {
		parent[i] = i
		connected[i] = 1
	}
	return &UnionFind{parent, connected}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}
	if uf.connected[rootX] < uf.connected[rootY] {
		rootX, rootY = rootY, rootX
	}
	uf.parent[rootY] = rootX
	uf.connected[rootX] += uf.connected[rootY]
	return true
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)

	for scnr.Scan() {
		p := parse(scnr.Text())
		points = append(points, p)
	}

	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, pair{i, j, distance(points[i], points[j])})
		}
	}
	sort.Slice(pairs, func(a, b int) bool { return pairs[a].distance < pairs[b].distance })

	uf := NewUnionFind(len(points))
	var lastI, lastJ int
	targetAttempts := len(points)
	if len(points) <= 20 {
		targetAttempts = 10
	}

	attempts := 0
	for _, p := range pairs {
		united := uf.Union(p.i, p.j)
		attempts++

		if united {
			lastI, lastJ = p.i, p.j
		}

		if attempts == targetAttempts {
			sizes := make(map[int]int)
			for i := range len(points) {
				root := uf.Find(i)
				sizes[root] = uf.connected[root]
			}

			circuitSizes := make([]int, 0, len(sizes))
			for _, size := range sizes {
				circuitSizes = append(circuitSizes, size)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))
			task1 = circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
		}

		if attempts > targetAttempts {
			sizes := make(map[int]int)
			for i := range len(points) {
				root := uf.Find(i)
				sizes[root] = uf.connected[root]
			}
			if len(sizes) == 1 {
				task2 = points[lastI].X * points[lastJ].X
				break
			}
		}
	}

	if task2 == 0 {
		panic("task2 not complete. A full merged circuit not accomplished")
	}

	fmt.Println("task1: ", task1) // 57970
	fmt.Println("task2: ", task2) // 8520040659
}

func distance(a, b Point3D) float64 {
	return math.Sqrt(
		math.Pow(float64(a.X)-float64(b.X), 2) +
			math.Pow(float64(a.Y)-float64(b.Y), 2) +
			math.Pow(float64(a.Z)-float64(b.Z), 2),
	)
}

func parse(s string) Point3D {
	splits := strings.Split(s, ",")
	if len(splits) != 3 {
		panic("not 3 vars")
	}

	x, err := strconv.Atoi(splits[0])
	if err != nil {
		panic("x not number" + splits[0])
	}
	y, err := strconv.Atoi(splits[1])
	if err != nil {
		panic("y not number" + splits[1])
	}
	z, err := strconv.Atoi(splits[2])
	if err != nil {
		panic("z not number" + splits[2])
	}

	return Point3D{x, y, z}
}
