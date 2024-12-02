package day02

import (
	"io"

	"aoc/common"
)

func getNumberOfSafeReports(r io.Reader, isSafe func(report []int) bool) int {
	safeReports := 0
	report := make([]int, 0, 10)
	for line := range common.ReadLinesLazy(r) {
		report = common.SplitToInts(line, " ", report)
		if isSafe(report) {
			safeReports += 1
		}
		report = report[:0]
	}
	return safeReports
}

func isReportSafe(report []int) bool {
	lastDiff := 0
	for i, item := range report[1:] {
		// i starts at 0 but range starts at 1
		prev := report[i]
		diff := item - prev

		// Sign change
		if diff*lastDiff < 0 {
			return false
		}
		absDiff := common.Abs(item - prev)
		if absDiff == 0 || absDiff > 3 {
			return false
		}
		lastDiff = diff
	}
	return true
}
