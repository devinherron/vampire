package game

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/devinherron/dice"
)

type Game struct {
	Prompts [][]string
	Visited [][]bool
}

func NewGame(prompts [][]string) {

	visited := make([][]bool, 81)
	finished := false

	currentPrompt := 1
	currentEntry := 0

	for !finished {
		result := Roll()
		if currentPrompt+result < 1 {
			currentPrompt = 1
		} else {
			currentPrompt += result
		}

		if len(visited[currentPrompt]) == 0 {
			visited[currentPrompt] = make([]bool, 3)
		}

		for visited[currentPrompt][currentEntry] {
			if currentEntry > 1 {
				currentPrompt++
				currentEntry = 0

				if len(visited[currentPrompt]) == 0 {
					visited[currentPrompt] = make([]bool, 3)
				}
			} else {
				currentEntry++
			}
		}

		fmt.Printf("[%d] %s\n", currentPrompt, prompts[currentPrompt][currentEntry])
		visited[currentPrompt][currentEntry] = true
		currentEntry = 0

		if currentPrompt > 71 {
			finished = true
		}

		var input string
		fmt.Scanln(&input)
		if input == "quit" {
			finished = true
		} else if input == "save" {
			Save(prompts, visited)
		}
	}

}

func Save(prompts [][]string, visited [][]bool) {
	game := Game{prompts, visited}

	f, err := os.OpenFile("save.json", os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	encoder := json.NewEncoder(f)
	encoder.Encode(game)
}

func Roll() int {
	d6 := dice.Roll(6)
	d10 := dice.Roll(10)

	return (d10 - d6)
}
