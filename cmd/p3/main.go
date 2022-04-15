package main

import (
	"aoc/utils"
	"fmt"
)

type Slope struct {
	right int
	down  int
}

func calculate(lines []string, sl Slope) int {
	pos := utils.Coord{X: 0, Y: 0}
	j := len(lines)
	tree := 0
	for pos.Y+sl.down < j {
		pos.Y += sl.down
		lineSize := len(lines[pos.Y])
		pos.X = (pos.X + sl.right) % lineSize
		v := lines[pos.Y][pos.X]
		if v == '#' {
			tree++
		}
	}
	return tree
}

func part1(path string) int {
	lines := utils.ReadLinesOfFile(path)
	return calculate(lines, Slope{3, 1})
}

func generatorCalculate(c chan int, lines []string, slope Slope) {
	go func() {
		c <- calculate(lines, slope)
	}()
}

func part2(path string) int {
	lines := utils.ReadLinesOfFile(path)
	c := make(chan int)
	generatorCalculate(c, lines, Slope{1, 1})
	generatorCalculate(c, lines, Slope{3, 1})
	generatorCalculate(c, lines, Slope{5, 1})
	generatorCalculate(c, lines, Slope{7, 1})
	generatorCalculate(c, lines, Slope{1, 2})
	resp := 1
	for i := 0; i < 5; i++ {
		resp *= <-c
	}
	return resp
}

func main() {
	p1 := make(chan int)
	p2 := make(chan int)
	go func() {
		p1 <- part1("input.txt")
	}()
	go func() {
		p2 <- part2("input.txt")
	}()
	fmt.Println("Star1:", <-p1)
	fmt.Println("Star2:", <-p2)
}
