package day07

import (
	"io"
)

func PartOne(r io.Reader) any {
	return getSumOfValidTestValues(r, add, mul)
}
