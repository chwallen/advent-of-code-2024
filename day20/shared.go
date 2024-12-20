package day20

import (
	"fmt"
	"io"

	"aoc/common"
)

const (
	partOneCheatTiles        = 2
	partTwoCheatTiles        = 20
	minStepsToSaveByCheating = 100
)

type tile struct {
	point common.Point
	steps int
	char  rune
	next  *tile
}

func findNumberOfGoodCheats(r io.Reader, maxCheatDistance int) int {
	lines := common.ReadLinesEager(r)
	grid := common.NewGrid[*tile](len(lines[0]), len(lines))
	var start *tile

	for y, line := range lines {
		for x, char := range line {
			t := tile{point: common.NewPoint(x, y), char: char}
			if char == 'S' {
				start = &t
			}
			grid.Set(x, y, &t)
		}
	}
	if start == nil {
		panic(fmt.Errorf("could not find start tile"))
	}

	steps := 0
	direction := common.Up
	for !traverseDirection(grid, start, &steps, direction) {
		direction = direction.TurnRight()
	}

	cheats := 0
	currentTile := start
	for i := 0; i < steps-minStepsToSaveByCheating; i++ {
		cheats += findCheatsAroundTile(grid, currentTile, maxCheatDistance)
		currentTile = currentTile.next
	}

	return cheats
}

func traverseDirection(
	grid common.Grid[*tile],
	current *tile,
	steps *int,
	dir common.Direction,
) bool {
	s := *steps
	next := grid.GetPoint(current.point.Add(dir))
	for next.char != '#' {
		s += 1
		next.steps = s
		current.next = next
		current = next
		next = grid.GetPoint(current.point.Add(dir))
	}
	if current.char == 'E' {
		return true
	}
	if s > *steps {
		*steps = s
		return traverseDirection(grid, current, steps, dir.TurnRight()) ||
			traverseDirection(grid, current, steps, dir.TurnLeft())
	}
	return false
}

// findCheatsAroundTile finds all cheats around t within the manhattan distance
// specified by maxDistance that save at least minStepsToSaveByCheating steps.
func findCheatsAroundTile(grid common.Grid[*tile], t0 *tile, maxDistance int) int {
	sizeX, sizeY := grid.Size()
	x0, y0 := t0.point.XY()
	steps := t0.steps
	cheats := 0

	for y1 := -maxDistance; y1 <= maxDistance; y1++ {
		y := y0 + y1
		if y < 1 {
			continue
		} else if y == sizeY-1 {
			// No point in trying outside the boundaries of the grid.
			break
		}
		absY1 := common.Abs(y1)
		for x1 := -maxDistance; x1 <= maxDistance; x1++ {
			x := x0 + x1
			manhattanDistance := absY1 + common.Abs(x1)
			if x < 1 || manhattanDistance > maxDistance {
				continue
			} else if x == sizeX-1 {
				// No point in trying outside the boundaries of the grid.
				break
			}
			t := grid.Get(x, y)
			savedSteps := t.steps - steps - manhattanDistance
			if t.char == '#' || savedSteps < minStepsToSaveByCheating {
				continue
			}
			cheats += 1
		}
	}
	return cheats
}
