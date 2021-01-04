package main

import "fmt"

var (
	game = NewGame()
)

func main() {
	fmt.Println(game.Bag.Name())
}
