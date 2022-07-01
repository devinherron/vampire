package game

import (
	"strconv"
	"testing"
)

func Test_Roll(t *testing.T) {
	result := Roll()
	if result < -5 || result > 9 {
		t.Errorf("Invalid Roll result: " + strconv.Itoa(result))
	}
}
