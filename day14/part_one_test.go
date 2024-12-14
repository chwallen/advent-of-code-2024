package day14

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartOne(t *testing.T) {
	expected := 222062148
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartOne(r).(int)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
