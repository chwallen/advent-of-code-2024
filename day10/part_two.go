package day10

import (
	"io"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	sum := 0
	_, trailRatings := hikeTrails(lines)
	for _, rating := range trailRatings {
		sum += rating
	}
	return sum
}
