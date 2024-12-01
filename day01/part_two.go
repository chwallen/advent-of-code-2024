package day01

import (
	"io"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	linesLen := len(lines)
	leftColumn := make([]int, linesLen)
	rightColumn := make(map[int]int, linesLen)

	for i, line := range lines {
		leftItem, rightItem := common.CutToInts(line, "   ")
		leftColumn[i] = leftItem
		rightColumn[rightItem] += 1
	}

	similarity := 0
	for _, leftItem := range leftColumn {
		similarity += leftItem * rightColumn[leftItem]
	}

	return similarity
}
