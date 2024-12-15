package day15

import (
	"io"
)

func PartOne(r io.Reader) any {
	return getBoxCoordinateSum(r, false)
}
