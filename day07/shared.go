package day07

import (
	"io"
	"strconv"
	"strings"

	"aoc/common"
)

type operation int

const (
	add operation = iota
	mul
	concat
)

func getSumOfValidTestValues(r io.Reader, ops ...operation) int {
	sum := 0
	values := make([]int, 0, 50)
	for line := range common.ReadLinesLazy(r) {
		left, right, _ := strings.Cut(line, ": ")
		expected := common.Atoi(left)
		values = common.SplitToInts(right, " ", values)

		if hasSolution(values, len(values)-1, expected, ops) {
			sum += expected
		}

		values = values[:0]
	}
	return sum
}

// Tries to find the solution backwards as it's much faster to prune invalid
// branches that way.
func hasSolution(values []int, i, value int, ops []operation) bool {
	if i == 0 {
		return value == values[0]
	}

	for _, op := range ops {
		switch op {
		case add:
			v := value - values[i]
			if v >= values[i-1] && hasSolution(values, i-1, v, ops) {
				return true
			}
		case mul:
			v, remainder := common.DivRem(value, values[i])
			if remainder == 0 && hasSolution(values, i-1, v, ops) {
				return true
			}
		case concat:
			lhs := strconv.Itoa(value)
			rhs := strconv.Itoa(values[i])
			if strings.HasSuffix(lhs, rhs) && len(lhs) > len(rhs) {
				v := common.Atoi(lhs[:len(lhs)-len(rhs)])
				if hasSolution(values, i-1, v, ops) {
					return true
				}
			}
		}
	}
	return false
}
