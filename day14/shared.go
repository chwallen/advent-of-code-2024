package day14

import (
	"io"
	"iter"
	"strings"

	"aoc/common"
)

const (
	height         = 103
	width          = 101
	partOneSeconds = 100
	partTwoSeconds = height * width
)

type robot struct {
	position common.Point
	vector   common.Point
}

func simulateRobotMovements(
	reader io.Reader,
	seconds int,
) iter.Seq2[int, []robot] {
	robots := make([]robot, 0, 500)

	for line := range common.ReadLinesLazy(reader) {
		position, velocity, _ := strings.Cut(line, " ")
		pX, pY := common.CutToInts(strings.TrimPrefix(position, "p="), ",")
		vX, vY := common.CutToInts(strings.TrimPrefix(velocity, "v="), ",")

		robots = append(robots, robot{
			position: common.NewPoint(pX, pY),
			vector:   common.NewPoint(vX, vY),
		})
	}

	return func(yield func(int, []robot) bool) {
		for second := 1; second <= seconds; second++ {
			for i, r := range robots {
				pX, pY := r.position.XY()
				vX, vY := r.vector.XY()
				nextX := (pX + vX + width) % width
				nextY := (pY + vY + height) % height
				robots[i].position = common.NewPoint(nextX, nextY)
			}
			if !yield(second, robots) {
				return
			}
		}
	}
}
