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
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	diagnostic := strings.Split(string(contents), "\n")
	gamma, epsilon := part1(diagnostic)
	fmt.Printf("%d gamma x %d epsilon = %d\n", gamma, epsilon, gamma*epsilon)
	o2, co2 := part2(diagnostic)
	fmt.Printf("%d o2 x %d co2 = %d\n", o2, co2, o2*co2)

}

func part1(diagnostics []string) (gamma, epsilon int) {
	result := counts(diagnostics)

	for i, count := range result {
		if count > 0 {
			gamma += int(math.Pow(2, float64(len(result)-(i+1))))
		} else {
			epsilon += int(math.Pow(2, float64(len(result)-(i+1))))
		}
	}
	return gamma, epsilon
}

func counts(diagnostics []string) []int {
	result := make([]int, len(diagnostics[0]))
	for _, d := range diagnostics {
		for i, v := range []rune(d) {
			if v == '1' {
				result[i]++
			} else {
				result[i]--
			}
		}
	}
	return result
}

func part2(diagnostics []string) (co2, o2 int) {
	return o2Scrubber(diagnostics), co2Scrubber(diagnostics)
}

func o2Scrubber(diagnostics []string) int {
	for i := 0; i < len(diagnostics[0]); i++ {
		if len(diagnostics) <= 1 {
			break
		}
		counts := counts(diagnostics)
		if counts[i] >= 0 {
			// only keep those with 1s
			diagnostics = filter(diagnostics, i, '1')
		} else {
			// only keep those with 0s
			diagnostics = filter(diagnostics, i, '0')
		}
	}
	o2, _ := strconv.ParseInt(diagnostics[0], 2, 64)
	return int(o2)
}

func co2Scrubber(diagnostics []string) int {
	for i := 0; i < len(diagnostics[0]); i++ {
		if len(diagnostics) <= 1 {
			break
		}
		counts := counts(diagnostics)
		if counts[i] >= 0 {
			// only keep those with 0s
			diagnostics = filter(diagnostics, i, '0')
		} else {
			// only keep those with 1s
			diagnostics = filter(diagnostics, i, '1')
		}
	}
	o2, _ := strconv.ParseInt(diagnostics[0], 2, 64)
	return int(o2)
}

func filter(diagnostics []string, i int, r rune) (res []string) {
	for _, s := range diagnostics {
		if rune(s[i]) != r {
			continue
		}
		res = append(res, s)
	}
	return res
}
