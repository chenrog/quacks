package main

type Game struct {
	Bag *Bag
	Pot *Pot
}

func NewGame() *Game {
	game := &Game{}

	game.Bag = NewBagFromJson("./domain/tokens.json")
	game.Pot = NewPotFromJson("./domain/pot_spaces.json")

	return game
}
