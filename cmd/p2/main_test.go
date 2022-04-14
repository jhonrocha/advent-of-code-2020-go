package main

import (
	"testing"
)

func TestDay(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		p1 := part1("test.txt")
		expected := 2
		if p1 != expected {
			t.Errorf("Part1 fail. Expected %d received %d", expected, p1)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		p2 := part2("test.txt")
		expected := 1
		if p2 != expected {
			t.Errorf("Part2 fail. Expected %d received %d", expected, p2)
		}
	})
}
