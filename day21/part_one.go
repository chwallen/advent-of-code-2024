package day21

import (
	"io"
)

func PartOne(r io.Reader) any {
	return calculateComplexitySum(r, partOneRobots)
}
