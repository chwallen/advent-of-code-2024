package day10

import (
	"fmt"

	"aoc/common"
)

func hikeTrails(
	lines []string,
) (uniqueTrails common.Set[string], trailRatings map[common.Point]int) {
	maxY := len(lines)
	maxX := len(lines[0])

	uniqueTrails = common.NewSet[string]()
	trailRatings = make(map[common.Point]int, 1000)

	for y, line := range lines {
		for x, char := range line {
			if char == '0' {
				trailStart := common.NewPoint(x, y)
				travelTrail(lines, trailStart, trailStart, maxX, maxY, uniqueTrails, trailRatings, '1')
			}
		}
	}

	return uniqueTrails, trailRatings
}

func travelTrail(
	lines []string,
	trailStart common.Point,
	previous common.Point,
	maxX int,
	maxY int,
	uniqueTrails common.Set[string],
	trailRatings map[common.Point]int,
	soughtSymbol uint8,
) {
	for _, neighbor := range previous.Neighbors() {
		x, y := neighbor.XY()
		if neighbor.IsWithinBounds(0, 0, maxX, maxY) && lines[y][x] == soughtSymbol {
			if soughtSymbol == '9' {
				key := fmt.Sprintf("%d,%d,%d,%d", trailStart.X(), trailStart.Y(), x, y)
				uniqueTrails.Add(key)
				trailRatings[neighbor]++
			} else {
				travelTrail(lines, trailStart, neighbor, maxX, maxY, uniqueTrails, trailRatings, soughtSymbol+1)
			}
		}
	}
}
