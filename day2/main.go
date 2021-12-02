package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	x, y int
}

func main() {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	plan := make([]move, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			log.Fatalf("expected direction got: %s", line)
		}
		num, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		switch split[0] {
		case "forward":
			plan[i].x = num
		case "up":
			plan[i].y = -num
		case "down":
			plan[i].y = num
		}
	}

	x, y := part1(plan)
	fmt.Printf("%d forward x %d down = %d\n", x, y, x*y)
	x, y = part2(plan)
	fmt.Printf("%d forward x %d down = %d\n", x, y, x*y)
}

func part1(plan []move) (int, int) {
	var distX, distY int
	for _, m := range plan {
		distX += m.x
		distY += m.y
	}
	return distX, distY
}

func part2(plan []move) (int, int) {
	var distX, distY, aim int
	for _, m := range plan {
		distX += m.x
		distY += m.x * aim
		aim += m.y
	}
	return distX, distY
}
