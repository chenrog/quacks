package main

import "fmt"

var (
	game = NewGame()
)

func main() {
	fmt.Println(game.Bag.Name())
	fmt.Println(game.Pot.Name())
}
