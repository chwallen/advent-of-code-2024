package day05

import (
	"io"
	"slices"
)

func PartTwo(r io.Reader) any {
	return getSumOfGoodUpdates(r, notEqual)
}

func notEqual(update, fixedUpdate []int) bool {
	return !slices.Equal(update, fixedUpdate)
}
