package day03

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	sum := 0
	for line := range common.ReadLinesLazy(r) {
		matches := multiplicationRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sum += common.Atoi(match[1]) * common.Atoi(match[2])
		}
	}
	return sum
}
