package day09

import (
	"io"

	"aoc/common"
)

type block struct {
	start  int
	length int
	id     int
}

func parseBlocks(r io.Reader, compact bool) (fileBlocks, emptyBlocks []block) {
	line := common.ReadLinesEager(r)[0]
	size := (len(line) + 1) / 2
	if !compact {
		size *= 5
	}
	fileBlocks = make([]block, 0, size)
	emptyBlocks = make([]block, 0, size)

	id := 0
	start := 0
	for i, char := range line {
		length := int(char - '0')
		if i%2 == 0 {
			if compact {
				fileBlocks = append(fileBlocks, block{start, length, id})
			} else {
				for j := 0; j < length; j++ {
					fileBlocks = append(fileBlocks, block{start + j, 1, id})
				}
			}
			id += 1
		} else if compact {
			emptyBlocks = append(emptyBlocks, block{start, length, -1})
		} else {
			for j := 0; j < length; j++ {
				emptyBlocks = append(emptyBlocks, block{start + j, 1, -1})
			}
		}
		start += length
	}

	return fileBlocks, emptyBlocks
}

func calculateChecksum(fileBlocks []block) int {
	checksum := 0
	for _, fileBlock := range fileBlocks {
		for i := 0; i < fileBlock.length; i++ {
			checksum += fileBlock.id * (fileBlock.start + i)
		}
	}
	return checksum
}
