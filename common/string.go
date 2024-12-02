package common

import (
	"fmt"
	"strconv"
	"strings"
)

// Atoi is equivalent to strconv.Atoi but panics on error.
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err == nil {
		return v
	}
	panic(fmt.Errorf("cannot convert string \"%s\" to int", s))
}

// CutToInts cuts s at sep using strings.Cut and converts the parts to ints.
// Panics if the parts cannot be converted.
func CutToInts(s string, sep string) (l, r int) {
	a, b, _ := strings.Cut(s, sep)
	return Atoi(a), Atoi(b)
}

// SplitToInts splits s at sep using strings.Split, converts and adds each part
// to b, and returns the updated slice. Panics if a part cannot be converted.
func SplitToInts(s string, sep string, b []int) []int {
	for _, v := range strings.Split(s, sep) {
		b = append(b, Atoi(v))
	}
	return b
}
