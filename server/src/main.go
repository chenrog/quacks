package main

import "fmt"

var (
	game = NewGame()
)

func main() {
	fmt.Println(game.Name())
	game.PlayToken()
	game.PlayToken()
	game.PlayToken()
	fmt.Println(game.Name())
}
