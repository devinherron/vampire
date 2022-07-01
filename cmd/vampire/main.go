package main

import (
	"fmt"
	"vampire/pkg/txtreader"

	"github.com/devinherron/dice"
)

func main() {
	prompts := txtreader.ReadFile("./assets/_Thousand Year Old Vampire.txt")
	fmt.Println(prompts[22][0])
	test := roll()
	fmt.Println(test)
}

func roll() int {
	_, d6Result := dice.Roll(6, 1)
	_, d10Result := dice.Roll(10, 1)

	return d10Result - d6Result
}
