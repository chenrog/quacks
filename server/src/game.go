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

func (g *Game) PlayToken() {
	token := g.Bag.GrabToken()
	g.Pot.PlaceToken(token)
}

func (g *Game) Name() string {
	return "Game:\n" + g.Bag.Name() + "\n" + g.Pot.Name()
}
