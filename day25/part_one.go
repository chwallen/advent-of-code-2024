package day25

import (
	"io"

	"aoc/common"
)

func PartOne(r io.Reader) any {
	lines := common.ReadLinesEager(r)

	locks := make([][]int, 0, 300)
	keys := make([][]int, 0, 300)
	for i := 0; i < len(lines); i += 8 {
		item := make([]int, 5)
		for col := 0; col < 5; col++ {
			for row := 0; row <= 6; row++ {
				if lines[i+row][col] == '#' {
					item[col] += 1
				}
			}
		}
		if lines[i][0] == '#' {
			locks = append(locks, item)
		} else {
			keys = append(keys, item)
		}
	}

	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			if keyFitsLock(lock, key) {
				result += 1
			}
		}
	}
	return result
}

func keyFitsLock(lock, key []int) bool {
	for i := 0; i < len(key); i++ {
		if lock[i]+key[i] > 7 {
			return false
		}
	}
	return true
}
