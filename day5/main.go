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
	all := read()
	fmt.Println(part1(all))
	fmt.Println(part2(all))

}

// probably the most convuluted way to do this.
func part2(vents []vent) int {
	type pos struct {
		x int
		y int
	}
	var hasVent = make(map[pos]int)
	for _, v := range vents {
		d, dist := getD(v)
		x, y := v.p[0][0], v.p[0][1]
		for i := 0; i < dist; i++ {
			hasVent[pos{
				x: x,
				y: y,
			}]++
			x += d[0]
			y += d[1]
		}
	}
	total := 0
	for _, count := range hasVent {
		if count > 1 {
			total++
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if count, ok := hasVent[pos{x: i, y: j}]; ok {
				fmt.Printf(" %d ", count)
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println()
	}
	return total
}

func getD(v vent) (d [2]int, dist int) {
	dx := v.p[1][0] - v.p[0][0]
	dy := v.p[1][1] - v.p[0][1]
	dist = int(math.Abs(float64(dx))) + 1
	if dist == 1 {
		dist = int(math.Abs(float64(dy))) + 1
	}
	if dx > 0 {
		d[0] = 1
	} else if dx < 0 {
		d[0] = -1
	}
	if dy > 0 {
		d[1] = 1
	} else if dy < 0 {
		d[1] = -1
	}
	return d, dist
}

func part1(all []vent) int {
	vents := filterHV(all)
	type pos struct {
		x int
		y int
	}
	var hasVent = make(map[pos]int)
	for _, v := range vents {
		if v.p[0][0] == v.p[1][0] {
			start, end := v.p[0][1], v.p[1][1]
			if start > end {
				start, end = end, start
			}
			for j := start; j <= end; j++ {
				hasVent[pos{x: v.p[0][0], y: j}]++
			}
		} else if v.p[0][1] == v.p[1][1] {
			start, end := v.p[0][0], v.p[1][0]
			if start > end {
				start, end = end, start
			}
			for j := start; j <= end; j++ {
				hasVent[pos{x: j, y: v.p[0][1]}]++
			}
		}
	}
	total := 0
	for _, count := range hasVent {
		if count > 1 {
			total++
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if count, ok := hasVent[pos{x: i, y: j}]; ok {
				fmt.Printf(" %d ", count)
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println()
	}
	return total
}

func filterHV(vents []vent) []vent {
	var res []vent
	for _, v := range vents {
		if v.p[0][0] == v.p[1][0] {
			res = append(res, v)
			continue
		}
		if v.p[0][1] == v.p[1][1] {
			res = append(res, v)
			continue
		}
	}
	return res
}

type vent struct {
	p [2][2]int
}

func read() (vents []vent) {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		// x,y -> x,y
		splt := strings.Split(line, "->")
		var v vent
		for i := 0; i < 2; i++ {
			p1 := strings.Split(splt[i], ",")
			for j := 0; j < 2; j++ {
				v.p[i][j], _ = strconv.Atoi(strings.TrimSpace(p1[j]))
			}
		}
		vents = append(vents, v)
	}
	return vents
}
