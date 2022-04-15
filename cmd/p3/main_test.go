package main

import (
	"testing"
)

func TestDay(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		p1 := part1("input_test.txt")
		expected := 7
		if p1 != expected {
			t.Errorf("Part1 fail. Expected %d received %d", expected, p1)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		p2 := part2("input_test.txt")
		expected := 336
		if p2 != expected {
			t.Errorf("Part2 fail. Expected %d received %d", expected, p2)
		}
	})
}
