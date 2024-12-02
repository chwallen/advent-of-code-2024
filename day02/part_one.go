package day02

import (
	"io"
)

func PartOne(r io.Reader) any {
	return getNumberOfSafeReports(r, isReportSafe)
}
