package day18

import (
	"io"
	"math"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)

	// Start backwards by maximizing corruption
	grid := createTileGrid(lines, len(lines))

	var cutoffTile string
	for i := len(lines) - 1; i >= 0; i-- {
		if getFewestStepsToTraverse(grid) < math.MaxInt {
			cutoffTile = lines[i+1]
			break
		}

		for y := 0; y < gridHeight; y++ {
			for x := 0; x < gridWidth; x++ {
				grid.Get(x, y).score = math.MaxInt
			}
		}
		setTileCorruption(grid, lines[i], false)
	}

	return cutoffTile
}
