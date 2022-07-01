package game

import (
	"fmt"

	"github.com/devinherron/dice"
)

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

		if currentPrompt > 71 {
			finished = true
		}

		var input string
		fmt.Scanln(&input)

		if input == "quit" {
			finished = true
		}
	}

}

func Roll() int {
	d6 := dice.Roll(6)
	d10 := dice.Roll(10)

	return (d10 - d6)
}
