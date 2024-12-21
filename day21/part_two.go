package day21

import (
	"io"
)

func PartTwo(r io.Reader) any {
	return calculateComplexitySum(r, partTwoRobots)
}
