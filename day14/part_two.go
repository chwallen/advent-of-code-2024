package day14

import (
	"io"

	"aoc/common"
)

func PartTwo(reader io.Reader) any {
	positions := make(map[common.Point]int, 547)
	for second, robots := range simulateRobotMovements(reader, partTwoSeconds) {
		if hasNoRobotOverlaps(robots, positions) {
			return second
		}
		clear(positions)
	}
	panic("simulation did not find a christmas tree")
}

func hasNoRobotOverlaps(robots []robot, positions map[common.Point]int) bool {
	for _, r := range robots {
		robotsOnSameTile := positions[r.position] + 1
		if robotsOnSameTile > 1 {
			return false
		}
		positions[r.position] = robotsOnSameTile
	}
	return true
}
