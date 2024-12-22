package day22

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	result := 0
	for line := range common.ReadLinesLazy(r) {
		secret := common.Atoi(line)
		for i := 0; i < steps; i++ {
			secret = calculateNextSecret(secret)
		}
		result += secret
	}

	return result
}
