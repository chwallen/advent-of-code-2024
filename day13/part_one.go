package day13

import (
	"io"
)

func PartOne(r io.Reader) any {
	return getFewestTokensToWinAllPrizes(r, 0)
}
