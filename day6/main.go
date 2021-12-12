package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	state := readInitialState()
	fmt.Println(part1(state))
	fmt.Println(part2(state))
}

func part1(state []int) int {
	return runSimulation(state, 80)
}

func runSimulation(state []int, days int) int {
	var jellies [9]int
	for _, s := range state {
		jellies[s] += 1
	}
	res := simulate(jellies, days)
	var total int
	for _, r := range res {
		total += r
	}
	return total
}

func part2(state []int) int {
	return runSimulation(state, 256)
}

func simulate(jellies [9]int, days int) [9]int {
	for i := 0; i < days; i++ {
		var next [9]int
		for j := 1; j < len(jellies); j++ {
			next[j-1] = jellies[j]
		}
		next[6] += jellies[0]
		next[8] = jellies[0]
		jellies = next
	}
	return jellies
}

func readInitialState() []int {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var state []int
	for _, v := range strings.Split(string(contents), ",") {
		s, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		state = append(state, s)
	}
	return state
}
