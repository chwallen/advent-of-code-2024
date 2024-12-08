package day08

import (
	"io"
)

func PartOne(r io.Reader) any {
	return getUniqueAntiNodesCount(r, false)
}
