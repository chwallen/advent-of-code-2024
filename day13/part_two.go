package day13

import (
	"io"
)

const (
	prizePointCoordinateModifier = 10_000_000_000_000
)

func PartTwo(r io.Reader) any {
	return getFewestTokensToWinAllPrizes(r, prizePointCoordinateModifier)
}
