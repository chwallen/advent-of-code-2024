package day06

import (
	"io"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)

	initialGrid, start := createGrid(lines)
	patrolledGrid := initialGrid.Clone()
	loopGrid := initialGrid.Clone()
	maxX, maxY := initialGrid.Size()

	traversePatrolPath(patrolledGrid, start)

	loops := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if patrolledGrid.Get(x, y).visited {
				loopGrid.Set(x, y, patrolTile{char: '#'})
				if tiles := traversePatrolPath(loopGrid, start); tiles == -1 {
					loops += 1
				}
				loopGrid.CopyFrom(initialGrid)
			}
		}
	}

	return loops
}
