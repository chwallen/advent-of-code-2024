package day15

import (
	"io"
)

func PartTwo(r io.Reader) any {
	return getBoxCoordinateSum(r, true)
}
