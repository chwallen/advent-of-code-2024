package day06

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	grid, start := createGrid(lines)
	return traversePatrolPath(grid, start)
}
