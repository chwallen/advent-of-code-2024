package day04

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	maxY := len(lines) - 3
	maxX := len(lines[0]) - 3

	xmasAppearances := 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {
				if x < maxX {
					if isXmasSequence(lines, x, y, common.Right) {
						xmasAppearances += 1
					}
					if y < maxY && isXmasSequence(lines, x, y, common.DownRight) {
						xmasAppearances += 1
					}
					if y >= 3 && isXmasSequence(lines, x, y, common.UpRight) {
						xmasAppearances += 1
					}
				}
				if x >= 3 {
					if isXmasSequence(lines, x, y, common.Left) {
						xmasAppearances += 1
					}
					if y < maxY && isXmasSequence(lines, x, y, common.DownLeft) {
						xmasAppearances += 1
					}
					if y >= 3 && isXmasSequence(lines, x, y, common.UpLeft) {
						xmasAppearances += 1
					}
				}
				if y < maxY && isXmasSequence(lines, x, y, common.Down) {
					xmasAppearances += 1
				}
				if y >= 3 && isXmasSequence(lines, x, y, common.Up) {
					xmasAppearances += 1
				}
			}
		}
	}

	return xmasAppearances
}

func isXmasSequence(lines []string, x, y int, dir common.Direction) bool {
	dx, dy := common.Point(dir).XY()
	return lines[y+dy][x+dx] == 'M' &&
		lines[y+2*dy][x+2*dx] == 'A' &&
		lines[y+3*dy][x+3*dx] == 'S'
}
