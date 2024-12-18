package day18

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	grid := createTileGrid(lines, partOneCorruption)
	return getFewestStepsToTraverse(grid)
}
