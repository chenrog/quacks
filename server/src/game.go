package main

type Game struct {
	Bag *Bag
	Pot *Pot
}

func NewGame() *Game {
	game := &Game{}

	game.Bag = NewBagFromJson("./domain/tokens.json")
	game.Pot = NewPot()

	return game
}
