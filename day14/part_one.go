package day14

import (
	"io"
)

func PartOne(reader io.Reader) any {
	var finalState []robot
	for _, finalState = range simulateRobotMovements(reader, partOneSeconds) {
	}

	topLeftQuadrant := 0
	topRightQuadrant := 0
	bottomRightQuadrant := 0
	bottomLeftQuadrant := 0
	for _, r := range finalState {
		x, y := r.position.XY()
		if x > width/2 {
			if y > height/2 {
				bottomRightQuadrant += 1
			} else if y < height/2 {
				topRightQuadrant += 1
			}
		} else if x < width/2 {
			if y > height/2 {
				bottomLeftQuadrant += 1
			} else if y < height/2 {
				topLeftQuadrant += 1
			}
		}
	}
	return topLeftQuadrant * topRightQuadrant * bottomRightQuadrant * bottomLeftQuadrant
}
