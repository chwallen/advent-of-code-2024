package day10

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	uniqueTrails, _ := hikeTrails(lines)
	return uniqueTrails.Len()
}
