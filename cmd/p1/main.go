package main

import (
	"fmt"
	"strconv"
  "aoc/utils"
)


func part1(path string) int {
	var exists = struct{}{}
	lines := utils.ReadLinesOfFile(path)
	const GOAL = 2020
	numbers := make(map[int]struct{})
	for _, line := range lines {
		entry, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}
		complement := GOAL - entry
		_, ok := numbers[complement]
		if ok {
			return entry * complement
		} else {
			numbers[entry] = exists
		}
	}
	return 0
}

func checkGoal(numbers map[int]struct{}, goal int) (int, bool) {
  for entry := range numbers {
		complement := goal - entry
    if _, ok := numbers[complement]; ok {
      return complement * entry, true
    }
  }
  return 0, false
}

func part2(path string) int {
	var exists = struct{}{}
	lines := utils.ReadLinesOfFile(path)
	const GOAL = 2020
	numbers := make(map[int]struct{})
	for _, line := range lines {
		entry, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}
		complement := GOAL - entry
		product, ok := checkGoal(numbers, complement)
		if ok {
			return entry * product
		} else {
			numbers[entry] = exists
		}
	}
	return 0
}

func main() {
	p1 := part1("input.txt")
  fmt.Println("Star1:", p1)
	p2 := part2("input.txt")
  fmt.Println("Star2:", p2)
}
