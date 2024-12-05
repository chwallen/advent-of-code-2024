package day05

import (
	"io"
	"slices"
)

func PartOne(r io.Reader) any {
	return getSumOfGoodUpdates(r, slices.Equal)
}
