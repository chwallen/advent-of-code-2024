package day02

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartTwo(t *testing.T) {
	expected := 577
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartTwo(r).(int)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
