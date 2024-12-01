package day01

import (
	"io"
	"slices"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	linesCount := len(lines)
	data := make([]int, 2*linesCount)
	leftColumn := data[:linesCount]
	rightColumn := data[linesCount:]

	for i, line := range lines {
		leftColumn[i], rightColumn[i] = common.CutToInts(line, "   ")
	}

	slices.Sort(leftColumn)
	slices.Sort(rightColumn)

	distance := 0
	for i := 0; i < linesCount; i++ {
		distance += common.Abs(leftColumn[i] - rightColumn[i])
	}

	return distance
}
