package day12

import (
	"io"
	"iter"

	"aoc/common"
)

type groupType int

const (
	unvisitedSameCrop groupType = iota
	visitedSameCrop
	otherCrop
)

type cropRegion struct {
	area      int
	perimeter int
	sides     int
}

type gardenTile struct {
	point   common.Point
	crop    rune
	visited bool
}

var directions = []common.Direction{common.Up, common.Right, common.Down, common.Left, common.Up}

func getCropRegions(r io.Reader) iter.Seq[cropRegion] {
	lines := common.ReadLinesEager(r)
	grid := common.NewGrid[*gardenTile](len(lines[0]), len(lines))
	sizeX, sizeY := grid.Size()

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			grid.Set(x, y, &gardenTile{
				point: common.NewPoint(x, y),
				crop:  rune(lines[y][x]),
			})
		}
	}

	return func(yield func(region cropRegion) bool) {
		for y := 0; y < sizeY; y++ {
			for x := 0; x < sizeX; x++ {
				tile := grid.Get(x, y)
				if tile.visited {
					continue
				}

				region := getCropRegion(grid, tile)
				if !yield(region) {
					return
				}
			}
		}
	}
}

func getCropRegion(
	grid common.Grid[*gardenTile],
	origin *gardenTile,
) cropRegion {
	queue := common.Queue[*gardenTile]{}
	region := cropRegion{}
	crop := origin.crop

	groupNeighbors := func(neighbor common.Point) groupType {
		t := getTileOrDefault(grid, neighbor)
		switch {
		case t.crop != crop:
			return otherCrop
		case t.visited:
			return visitedSameCrop
		default:
			return unvisitedSameCrop
		}
	}

	queue.Push(origin)
	for queue.Len() > 0 {
		tile := queue.Pop()
		// Unvisited tiles may become visited while in the queue
		if tile.visited {
			continue
		}
		tile.visited = true

		groups := common.GroupBy(tile.point.Neighbors(), groupNeighbors)
		region.area += 1
		region.perimeter += len(groups[otherCrop])
		region.sides += countCorners(grid, tile.point, crop)

		for _, p := range groups[unvisitedSameCrop] {
			queue.Push(grid.GetPoint(p))
		}
	}

	return region
}

func countCorners(
	grid common.Grid[*gardenTile],
	origin common.Point,
	crop rune,
) int {
	corners := 0
	for i := 0; i < len(directions)-1; i++ {
		dir1 := directions[i]
		dir2 := directions[i+1]

		side1Crop := getTileOrDefault(grid, origin.Add(dir1)).crop
		side2Crop := getTileOrDefault(grid, origin.Add(dir2)).crop
		diagonalCrop := getTileOrDefault(grid, origin.Add(dir1).Add(dir2)).crop

		if (crop != side1Crop && crop != side2Crop) ||
			(crop == side1Crop && crop == side2Crop && crop != diagonalCrop) {
			corners += 1
		}
	}
	return corners
}

func getTileOrDefault(grid common.Grid[*gardenTile], point common.Point) gardenTile {
	if grid.IsPointWithinBounds(point) {
		return *grid.GetPoint(point)
	}
	return gardenTile{}
}
