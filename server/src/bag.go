package main

type Bag struct {
	Tokens []*Token
}

func NewBag() *Bag {
	bag := &Bag{}
	return bag
}

func (bag *Bag) AddToken(token *Token) {
	bag.Tokens = append(bag.Tokens, token)
}

func (bag *Bag) Name() string {
	name := "Bag: "

	for _, token := range bag.Tokens {
		name += token.Name()
		name += ", "
	}

	return name
}