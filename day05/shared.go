package day05

import (
	"io"
	"slices"

	"aoc/common"
)

func getSumOfGoodUpdates(
	r io.Reader,
	isGoodUpdate func(update, fixedUpdate []int) bool,
) int {
	lines := common.ReadLinesEager(r)
	rules := make(map[int][]int)

	var i int
	for i = 0; lines[i] != ""; i++ {
		precedingPage, subsequentPage := common.CutToInts(lines[i], "|")
		rules[precedingPage] = append(rules[precedingPage], subsequentPage)
	}

	fixUpdate := func(pageA, pageB int) int {
		if slices.Contains(rules[pageA], pageB) {
			return -1
		} else if slices.Contains(rules[pageB], pageA) {
			return 1
		}
		return 0
	}

	update := make([]int, 0, 30)
	updateCopy := make([]int, 30)
	sum := 0
	for _, line := range lines[i+1:] {
		update = common.SplitToInts(line, ",", update)
		copy(updateCopy, update)

		fixedUpdate := updateCopy[:len(update)]
		slices.SortFunc(fixedUpdate, fixUpdate)

		if isGoodUpdate(update, fixedUpdate) {
			sum += fixedUpdate[len(fixedUpdate)/2]
		}

		update = update[:0]
	}
	return sum
}
