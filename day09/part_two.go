package day09

import (
	"io"
	"slices"
)

func PartTwo(r io.Reader) any {
	fileBlocks, emptyBlocks := parseBlocks(r, true)

	for fileIndex, file := range slices.Backward(fileBlocks) {
		for emptyBlockIndex, emptyBlock := range emptyBlocks {
			if emptyBlock.start >= file.start {
				break
			}
			if emptyBlock.length >= file.length {
				fileBlocks[fileIndex].start = emptyBlock.start
				emptyBlocks[emptyBlockIndex].start += file.length
				emptyBlocks[emptyBlockIndex].length -= file.length
				break
			}
		}
	}

	return calculateChecksum(fileBlocks)
}
