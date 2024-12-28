package day23

import (
	"strings"
	"testing"

	"aoc/common"
)

func TestPartTwo(t *testing.T) {
	expected := "aj,ds,gg,id,im,jx,kq,nj,ql,qr,ua,yh,zn"
	testInput := common.ReadFileToString("input.txt")
	r := strings.NewReader(testInput)

	actual := PartTwo(r).(string)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
