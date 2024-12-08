package day08

import (
	"io"
	"math"

	"aoc/common"
)

func isAntennaWithinBounds(a common.Point, xMax, yMax int) bool {
	return 0 <= a.X() && a.X() < xMax && 0 <= a.Y() && a.Y() < yMax
}

func getUniqueAntiNodesCount(r io.Reader, unlimitedDistance bool) int {
	lines := common.ReadLinesEager(r)
	antennas := make(map[rune][]common.Point)
	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			p := common.NewPoint(x, y)
			antennas[char] = append(antennas[char], p)
		}
	}

	antiNodes := make(map[common.Point]bool)
	maxY := len(lines)
	maxX := len(lines[0])
	for _, points := range antennas {
		for i := 0; i < len(points)-1; i++ {
			pointA := points[i]
			for j := i + 1; j < len(points); j++ {
				pointB := points[j]
				diffX := common.Abs(pointA.X() - pointB.X())
				diffY := common.Abs(pointA.Y() - pointB.Y())

				aDiff, bDiff := diffX, diffX
				if pointA.X() > pointB.X() {
					bDiff = -diffX
				} else {
					aDiff = -diffX
				}

				end := 1
				if unlimitedDistance {
					end = math.MaxInt
					antiNodes[pointA] = true
					antiNodes[pointB] = true
				}
				for k := 1; k <= end; k++ {
					antiNodeOne := common.NewPoint(pointA.X()+k*aDiff, pointA.Y()-k*diffY)
					antiNodeTwo := common.NewPoint(pointB.X()+k*bDiff, pointB.Y()+k*diffY)

					exit := true
					if isAntennaWithinBounds(antiNodeOne, maxX, maxY) {
						antiNodes[antiNodeOne] = true
						exit = false
					}
					if isAntennaWithinBounds(antiNodeTwo, maxX, maxY) {
						antiNodes[antiNodeTwo] = true
						exit = false
					}
					if exit {
						break
					}
				}
			}
		}
	}

	return len(antiNodes)
}
