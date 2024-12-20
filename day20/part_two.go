package day20

import (
	"io"
)

func PartTwo(r io.Reader) any {
	return findNumberOfGoodCheats(r, partTwoCheatTiles)
}
