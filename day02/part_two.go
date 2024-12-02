package day02

import (
	"io"
	"slices"
)

func PartTwo(r io.Reader) any {
	return getNumberOfSafeReports(r, isReportSafeWhenModified)
}

func isReportSafeWhenModified(report []int) bool {
	for toSkip := 0; toSkip < len(report); toSkip++ {
		item := report[toSkip]
		report = append(report[:toSkip], report[toSkip+1:]...)
		if isReportSafe(report) {
			return true
		}
		report = slices.Insert(report, toSkip, item)
	}
	return false
}
