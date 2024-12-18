package day18

import (
	"math"

	"aoc/common"
)

type tile struct {
	corrupted bool
	score     int
}

type mazePath struct {
	pos   common.Point
	tile  *tile
	score int
}

const (
	partOneCorruption = 1024
	gridHeight        = 71
	gridWidth         = 71
	startX            = 0
	startY            = 0
)

var start = common.NewPoint(startX, startY)

func createTileGrid(lines []string, corruptionLimit int) common.Grid[*tile] {
	grid := common.NewGrid[*tile](gridWidth, gridHeight)
	sizeX, sizeY := grid.Size()
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			grid.Set(x, y, &tile{corrupted: false, score: math.MaxInt})
		}
	}
	for i := 0; i < corruptionLimit; i++ {
		setTileCorruption(grid, lines[i], true)
	}

	return grid
}

func setTileCorruption(grid common.Grid[*tile], line string, corrupt bool) {
	x, y := common.CutToInts(line, ",")
	grid.Get(x, y).corrupted = corrupt
}

func getFewestStepsToTraverse(grid common.Grid[*tile]) int {
	queue := common.Queue[mazePath]{}
	startTile := grid.GetPoint(start)
	startTile.score = 0
	queue.Push(mazePath{start, startTile, 0})

	fewestSteps := math.MaxInt
	for queue.Len() > 0 {
		item := queue.Pop()
		if item.score >= fewestSteps {
			break
		} else if item.pos.X() == gridWidth-1 && item.pos.Y() == gridHeight-1 {
			fewestSteps = item.score
		} else {
			steps := item.score + 1
			for _, neighbor := range item.pos.Neighbors() {
				if !grid.IsPointWithinBounds(neighbor) {
					continue
				}

				nextTile := grid.GetPoint(neighbor)
				if nextTile.corrupted || steps >= nextTile.score {
					continue
				}
				nextTile.score = steps
				queue.Push(mazePath{neighbor, nextTile, steps})
			}
		}
	}

	return fewestSteps
}
