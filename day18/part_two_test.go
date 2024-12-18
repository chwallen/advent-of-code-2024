package day18

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartTwo(t *testing.T) {
	expected := "26,50"
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartTwo(r).(string)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
