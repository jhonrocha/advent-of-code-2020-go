package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func part1(path string) int {
	lines := utils.ReadLinesOfFile(path)
	valid := 0
	for _, line := range lines {
		var (
			min, max int
			repeat   string
			psw      string
		)
		fmt.Sscanf(line, "%d-%d %1s: %s", &min, &max, &repeat, &psw)
		count := strings.Count(psw, repeat)
		if count >= min && count <= max {
			valid++
		}
	}
	return valid
}

func part2(path string) int {
	lines := utils.ReadLinesOfFile(path)
	valid := 0
	for _, line := range lines {
		var (
			min, max int
			repeat   byte
			psw      string
		)
		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &repeat, &psw)
		hits := 0
    l := len(psw)
		if l >= min && psw[min-1] == repeat {
			hits++
		}
		if l >= max && psw[max-1] == repeat {
			hits++
		}
		if hits == 1 {
			valid++
		}
	}
	return valid
}

func main() {
	p1 := part1("input.txt")
	fmt.Println("Star1:", p1)
	p2 := part2("input.txt")
	fmt.Println("Star2:", p2)
}
