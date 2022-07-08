package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/devinherron/dice"
)

var letters = map[int]string{
	0: "a",
	1: "b",
	2: "c",
}

type Game struct {
	Prompts [][]string
	Visited [][]bool
	Prompt  int
	Entry   int
}

func Run(game Game, new bool) {
	finished := false

	if new {
		Advance(&game)
	}

	for !finished {
		fmt.Printf("[%d%s] %s\n", game.Prompt, letters[game.Entry], game.Prompts[game.Prompt][game.Entry])
		game.Visited[game.Prompt][game.Entry] = true
		game.Entry = 0

		if game.Prompt > 71 {
			finished = true
		}

		var input string
		fmt.Scanln(&input)

		Advance(&game)
		if input == "quit" {
			finished = true
		} else if input == "save" {
			Save(game)
		}
	}
}

func Load() {
	f, err := ioutil.ReadFile("save.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var game Game
	err = json.Unmarshal(f, &game)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	Run(game, false)
}

func New(prompts [][]string) {
	visited := make([][]bool, 81)

	Run(Game{prompts, visited, 1, 0}, true)
}

func Save(game Game) {
	f, err := os.OpenFile("save.json", os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	encoder := json.NewEncoder(f)
	encoder.Encode(game)
}

func Advance(game *Game) {
	result := Roll()
	if game.Prompt+result < 1 {
		game.Prompt = 1
	} else {
		game.Prompt += result
	}

	if len(game.Visited[game.Prompt]) == 0 {
		game.Visited[game.Prompt] = make([]bool, 3)
	}

	for game.Visited[game.Prompt][game.Entry] {
		if game.Entry > 1 {
			game.Prompt++
			game.Entry = 0

			if len(game.Visited[game.Prompt]) == 0 {
				game.Visited[game.Prompt] = make([]bool, 3)
			}
		} else {
			game.Entry++
		}
	}
}

func Roll() int {
	d6 := dice.Roll(6)
	d10 := dice.Roll(10)

	return (d10 - d6)
}
