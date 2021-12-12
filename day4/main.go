package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	game, boards := readBoards()
	fmt.Println(part1(game, boards))
	fmt.Println(part2(game, boards))
}

type board struct {
	cells [5][5]int
}

func readBoards() ([]int, []board) {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	snums := strings.Split(lines[0], ",")
	nums := make([]int, len(snums))
	for i, n := range snums {
		nums[i], _ = strconv.Atoi(n)
	}
	var boards []board
	for i := 2; i < len(lines); i += 6 {
		var b board
		for j := 0; j < 5; j++ {
			line := strings.Split(lines[i+j], " ")
			var k int
			for l := 0; l < len(line); l++ {
				if len(line[l]) == 0 {
					continue
				}
				v, _ := strconv.Atoi(line[l])
				b.cells[j][k] = v
				k++
			}
		}
		boards = append(boards, b)
	}
	return nums, boards
}

func part1(game []int, boards []board) int {
	scores := convert(boards)
	var winner []*scoreBoard
	for _, v := range game {
		winner = updateScores(v, scores[v])
		if len(winner) > 0 {
			break
		}
	}
	return calculateScore(winner[0])
}

func part2(game []int, boards []board) int {
	scores := convert(boards)
	var winner []*scoreBoard
	for _, v := range game {
		winner = append(winner, updateScores(v, scores[v])...)
	}
	return calculateScore(winner[len(winner)-1])
}

func calculateScore(s *scoreBoard) int {
	var sum int
	for v := range s.vals {
		called := false
		for _, cv := range s.called {
			if v == cv {
				called = true
				break
			}
		}
		if called {
			continue
		}
		sum += v
	}
	return sum * s.called[len(s.called)-1]
}

func updateScores(v int, boards []*scoreBoard) (completed []*scoreBoard) {
	for _, b := range boards {
		if b.completed {
			continue
		}
		b.called = append(b.called, v)
		c := b.vals[v]
		b.rowSums[c.i] += 1
		b.colSums[c.j] += 1
		if c.i == c.j {
			b.diagonalSums[0] += 1
		}
		if c.i+c.j == 5 {
			b.diagonalSums[1] += 1
		}
		if b.rowSums[c.i] == 5 || b.colSums[c.j] == 5 || b.diagonalSums[0] == 5 || b.diagonalSums[1] == 5 {
			completed = append(completed, b)
			b.completed = true
		}
	}
	return completed
}

type cell struct {
	i, j int
}

type scoreBoard struct {
	i            int
	vals         map[int]cell
	rowSums      [5]int
	colSums      [5]int
	diagonalSums [2]int
	called       []int
	completed    bool
}

func convert(boards []board) map[int][]*scoreBoard {
	init := make(map[int][]*scoreBoard)
	for idx, b := range boards {
		var s scoreBoard
		s.vals = make(map[int]cell)
		s.i = idx
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				val := b.cells[i][j]
				s.vals[val] = cell{
					i: i,
					j: j,
				}
				init[val] = append(init[val], &s)
			}
		}
	}
	return init
}
