package day22

import (
	"io"
	"math"

	"aoc/common"
)

type priceChangeSequence struct {
	first  int
	second int
	third  int
	fourth int
}

func PartTwo(r io.Reader) any {
	bananasForSequence := make(map[priceChangeSequence]int, 50_000)
	knownSequences := common.NewSet[priceChangeSequence]()

	for line := range common.ReadLinesLazy(r) {
		secret := common.Atoi(line)

		previousPrice := secret % 10
		seq := priceChangeSequence{}
		for i := 0; i < steps; i++ {
			secret = calculateNextSecret(secret)

			price := secret % 10
			seq = priceChangeSequence{
				first:  seq.second,
				second: seq.third,
				third:  seq.fourth,
				fourth: price - previousPrice,
			}
			previousPrice = price

			if i >= 3 && knownSequences.Add(seq) {
				bananasForSequence[seq] += price
			}
		}
		knownSequences.Clear()
	}

	maxBananas := math.MinInt
	for _, bananas := range bananasForSequence {
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}
	return maxBananas
}
