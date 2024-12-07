package day07

import (
	"io"
)

func PartTwo(r io.Reader) any {
	return getSumOfValidTestValues(r, add, mul, concat)
}
