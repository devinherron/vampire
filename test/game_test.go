package game

import (
	"strconv"
	"testing"
	"vampire/pkg/game"
)

func Test_Roll(t *testing.T) {
	result := game.Roll()
	if result < -5 || result > 9 {
		t.Errorf("Invalid Roll result: " + strconv.Itoa(result))
	}
}
