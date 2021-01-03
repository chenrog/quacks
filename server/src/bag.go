package main

type Bag struct {
	Tokens []*Token
}

func NewBag() *Bag {
	bag := &Bag{}
	return bag
}

func (b *Bag) AddToken(t *Token) {
	b.Tokens = append(b.Tokens, t)
}

func (b *Bag) Name() string {
	name := "Bag: "

	for _, token := range b.Tokens {
		name += token.Name()
		name += ", "
	}

	return name
}
