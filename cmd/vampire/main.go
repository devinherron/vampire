package main

import (
	"flag"
	"fmt"
	"strings"
	"vampire/pkg/game"
	"vampire/pkg/txtreader"
)

var file string

func init() {
	flag.StringVar(&file, "f", "./assets/_Thousand Year Old Vampire.txt", "Specify the file to load game text from.")
	flag.Parse()
}

func main() {
	var input string
	fmt.Println("(N)ew Game or (L)oad Game")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	switch input {
	case "n":
		new()
	case "l":
		load()
	}
}

func load() {
	game.Load()
}

func new() {
	prompts := txtreader.ReadPrompts(file)
	game.New(prompts)
}
