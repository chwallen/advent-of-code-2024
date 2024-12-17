package day17

import (
	"io"
	"slices"
	"strings"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	p := parseIntoProgram(lines)

	tokens := strings.TrimPrefix(lines[4], "Program: ")
	expectedOutput := common.SplitToInts(tokens, ",", make([]int, 0, 20))

	// A must be at least 8^(len(expectedOutput)-1) to produce len(expectedOutput) symbols
	initialA := common.IntPow(8, len(expectedOutput)-1)
	for {
		p.a = initialA
		actualOutput := p.execute()
		if slices.Equal(expectedOutput, actualOutput) {
			return initialA
		}

		for i := len(expectedOutput) - 1; i >= 0; i-- {
			if expectedOutput[i] != actualOutput[i] {
				// The token at position i changes every 8^i increment of A
				initialA += common.IntPow(8, i)
				p.outputIndex = 0
				p.b = 0
				p.c = 0
				break
			}
		}
	}
}
