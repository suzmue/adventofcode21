package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	depths := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		depths[i] = num
	}
	fmt.Println(part1(depths))
	fmt.Println(part2(depths))
}

func part1(depths []int) int {
	var count int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return count
}

func part2(depths []int) int {
	var count int
	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			count++
		}
	}
	return count
}
