package day05

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartOne(t *testing.T) {
	expected := 4185
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartOne(r).(int)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
