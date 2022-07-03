package main

import (
	"flag"
	"vampire/pkg/game"
	"vampire/pkg/txtreader"
)

var file string

func init() {
	flag.StringVar(&file, "f", "./assets/_Thousand Year Old Vampire.txt", "Specify the file to load game text from.")
	flag.Parse()
}

func main() {
	prompts := txtreader.ReadPrompts(file)
	game.NewGame(prompts)
}
