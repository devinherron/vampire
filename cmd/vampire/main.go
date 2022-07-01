package main

import (
	"vampire/pkg/game"
	"vampire/pkg/txtreader"
)

func main() {
	prompts := txtreader.ReadFile("./assets/_Thousand Year Old Vampire.txt")
	game.NewGame(prompts)
}
