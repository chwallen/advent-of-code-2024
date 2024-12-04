package day04

import (
	"io"

	"aoc/common"
)

func PartTwo(r io.Reader) any {
	lines := common.ReadLinesEager(r)
	xmasAppearances := 0
	for i := 0; i < len(lines)-2; i++ {
		line := lines[i]
		for j := 0; j < len(line)-2; j++ {
			if lines[i+1][j+1] == 'A' {
				topLeft := line[j]
				topRight := line[j+2]
				bottomLeft := lines[i+2][j]
				bottomRight := lines[i+2][j+2]

				if ((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) &&
					((topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')) {
					xmasAppearances += 1
				}
			}
		}
	}

	return xmasAppearances
}
