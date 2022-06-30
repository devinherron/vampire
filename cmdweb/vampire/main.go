package main

import (
	"fmt"

	"github.com/devinherron/dice"
)

func main() {
	test := roll()
	fmt.Println(test)
}

func roll() int {
	_, d6Result := dice.Roll(6, 1)
	_, d10Result := dice.Roll(10, 1)

	return d10Result - d6Result
}
