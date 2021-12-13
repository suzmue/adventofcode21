package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabs := readCrabPos()
	fmt.Println(part1(crabs))
	fmt.Println(part2(crabs))
}

func part1(crabs []int) int {
	max, m := makeMap(crabs)
	_, e := search(len(crabs), 0, max, m, energy1)
	return e
}

func search(n, min, max int, m map[int]int, energy func(int, map[int]int) int) (int, int) {
	pos := -1
	minEnergy := math.MaxInt
	for i := min; i < max; i++ {
		e := energy(i, m)
		if e < minEnergy {
			pos = i
			minEnergy = e
		}
	}
	return pos, minEnergy
}

func energy1(val int, m map[int]int) int {
	total := 0
	for x, count := range m {
		perCrab := val - x
		if perCrab < 0 {
			perCrab = x - val
		}
		total += perCrab * count
	}
	return total
}

func part2(crabs []int) int {
	max, m := makeMap(crabs)
	expended := make([]int, max+1)
	for i := 0; i < max+1; i++ {
		if i == 0 {
			continue
		}
		expended[i] = i + expended[i-1]
	}
	var energy2 = func(val int, m map[int]int) int {
		total := 0
		for x, count := range m {
			dist := val - x
			if dist < 0 {
				dist = x - val
			}
			e := expended[dist]
			total += e * count
		}
		return total
	}

	_, e := search(len(crabs), 0, max, m, energy2)
	return e
}

func makeMap(crabs []int) (int, map[int]int) {
	var max = 0
	crabMap := make(map[int]int)
	for _, c := range crabs {
		if c > max {
			max = c
		}
		crabMap[c] += 1
	}
	return max, crabMap
}

func readCrabPos() []int {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var crabs []int
	for _, c := range strings.Split(string(contents), ",") {
		x, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		crabs = append(crabs, x)
	}
	return crabs
}
