package day24

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartTwo(t *testing.T) {
	expected := "fvw,grf,mdb,nwq,wpq,z18,z22,z36"
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartTwo(r).(string)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
