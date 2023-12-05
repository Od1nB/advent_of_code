package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var (
	seedToSoil = []ConvMap{}
	soilToFert = []ConvMap{}
	fertToWtr  = []ConvMap{}
	wtrToLght  = []ConvMap{}
	lghtToTmp  = []ConvMap{}
	tmpToHum   = []ConvMap{}
	humToLoc   = []ConvMap{}
)

type ConvMap struct {
	src int
	dst int
	rng int
}

func fromLine(line string) ConvMap {
	cMap := &ConvMap{}

	split := lo.Map(strings.Split(line, " "), func(x string, _ int) int {
		num, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		return num
	})

	cMap.src = split[1]
	cMap.dst = split[0]
	cMap.rng = split[2]

	return *cMap
}

func Solution1() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)

	scnr.Scan()
	line := scnr.Text()
	seeds := lo.Map(strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " "), func(x string, _ int) int {
		num, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		return num
	})

	readMaps(scnr, &seedToSoil)
	readMaps(scnr, &soilToFert)
	readMaps(scnr, &fertToWtr)
	readMaps(scnr, &wtrToLght)
	readMaps(scnr, &lghtToTmp)
	readMaps(scnr, &tmpToHum)
	readMaps(scnr, &humToLoc)

	locs := make([]int, 0)

	for _, seed := range seeds {
		soil := srcToDst(seed, seedToSoil)
		fert := srcToDst(soil, soilToFert)
		wtr := srcToDst(fert, fertToWtr)
		lght := srcToDst(wtr, wtrToLght)
		tmp := srcToDst(lght, lghtToTmp)
		hum := srcToDst(tmp, tmpToHum)
		loc := srcToDst(hum, humToLoc)
		locs = append(locs, loc)
	}

	fmt.Println(lo.Min(locs))
	return nil
}
func Solution2() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)

	scnr.Scan()
	line := scnr.Text()
	seeds := lo.Map(strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " "), func(x string, _ int) int {
		num, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		return num
	})

	readMaps(scnr, &seedToSoil)
	readMaps(scnr, &soilToFert)
	readMaps(scnr, &fertToWtr)
	readMaps(scnr, &wtrToLght)
	readMaps(scnr, &lghtToTmp)
	readMaps(scnr, &tmpToHum)
	readMaps(scnr, &humToLoc)

	locs := make([]int, 0)

	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			soil := srcToDst(seed, seedToSoil)
			fert := srcToDst(soil, soilToFert)
			wtr := srcToDst(fert, fertToWtr)
			lght := srcToDst(wtr, wtrToLght)
			tmp := srcToDst(lght, lghtToTmp)
			hum := srcToDst(tmp, tmpToHum)
			loc := srcToDst(hum, humToLoc)
			locs = append(locs, loc)
		}
	}

	fmt.Println(lo.Min(locs))
	return nil
}

func srcToDst(val int, maps []ConvMap) int {
	for _, m := range maps {
		if val >= m.src && val < m.src+m.rng {
			return m.dst + (val - m.src)
		}
	}
	return val
}

func readMaps(scnr *bufio.Scanner, maps *[]ConvMap) {
	var line string
	for scnr.Scan() {
		line := scnr.Text()
		if strings.Contains(line, ":") {
			break
		}
	}

	for scnr.Scan() {
		line = scnr.Text()
		if line == "" {
			break
		}

		*maps = append(*maps, fromLine(line))
	}
}

func main() {
	//s := time.Now()
	var e error
	e = Solution1() //322500873
	e = Solution2() //108956227
	//fmt.Println(time.Since(s))
	if e != nil {
		panic(e)
	}
}
