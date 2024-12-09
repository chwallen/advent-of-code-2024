package day09

import (
	"io"
	"slices"
)

func PartOne(r io.Reader) any {
	fileBlocks, emptyBlocks := parseBlocks(r, false)

	emptyBlockIndex := 0
	for fileIndex, file := range slices.Backward(fileBlocks) {
		emptyBlock := emptyBlocks[emptyBlockIndex]

		if file.start < emptyBlock.start {
			break
		}

		fileBlocks[fileIndex].start = emptyBlock.start
		emptyBlocks[emptyBlockIndex].start = file.start
		emptyBlockIndex += 1
	}

	return calculateChecksum(fileBlocks)
}
