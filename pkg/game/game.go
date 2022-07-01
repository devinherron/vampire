package game

import (
	"github.com/devinherron/dice"
)

func NewGame(prompts [][]string) {
	println(Roll())
}

func Roll() int {
	d6 := dice.Roll(6)
	d10 := dice.Roll(10)

	return (d10 - d6)
}
