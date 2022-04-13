package main

import (
	"testing"
)

func TestDay(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		p1 := part1("test.txt")
		if p1 != 514579 {
			t.Errorf("Part1 fail")
		}
	})

	t.Run("Part2", func(t *testing.T) {
		p2 := part2("test.txt")
		if p2 != 241861950 {
			t.Errorf("Part2 fail")
		}
	})
}
