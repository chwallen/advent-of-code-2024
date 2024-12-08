package day08

import (
	"io"
)

func PartTwo(r io.Reader) any {
	return getUniqueAntiNodesCount(r, true)
}
