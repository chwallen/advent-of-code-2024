package day20

import (
	"io"
)

func PartOne(r io.Reader) any {
	return findNumberOfGoodCheats(r, partOneCheatTiles)
}
