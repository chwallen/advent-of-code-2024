package day17

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartOne(t *testing.T) {
	expected := "4,1,7,6,4,1,0,2,7"
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartOne(r).(string)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
