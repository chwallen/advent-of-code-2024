package day06

import (
	"aoc/common"
)

type patrolTile struct {
	char        rune
	visited     bool
	visitedDirs [4]bool
}

func createGrid(lines []string) (grid common.Grid[patrolTile], start common.Point) {
	grid = common.NewGrid[patrolTile](len(lines[0]), len(lines))

	for y, line := range lines {
		for x, char := range line {
			grid.Set(x, y, patrolTile{char: char})
			if char == '^' {
				start = common.NewPoint(x, y)
			}
		}
	}
	return grid, start
}

func traversePatrolPath(
	grid common.Grid[patrolTile],
	start common.Point,
) int {
	current := start
	dir := common.Up
	cardinalIndex := dir.GetCardinalIndex()
	uniqueTiles := 0

	for {
		tile := grid.GetPoint(current)
		if !tile.visited {
			tile.visited = true
			uniqueTiles += 1
		} else if tile.visitedDirs[cardinalIndex] {
			return -1
		}
		tile.visitedDirs[cardinalIndex] = true
		grid.Set(current.X(), current.Y(), tile)

		next := current.Add(dir)
		if !grid.IsPointWithinBounds(next) {
			return uniqueTiles
		}

		if grid.GetPoint(next).char != '#' {
			current = next
		} else {
			dir = dir.TurnRight()
			cardinalIndex = dir.GetCardinalIndex()
		}
	}
}
