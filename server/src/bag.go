package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

func NewBagFromJson(jsonPath string) *Bag {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var bag Bag
	json.Unmarshal([]byte(byteValue), &bag)

	return &bag
}

func (b *Bag) Name() string {
	name := "Bag: "

	for _, token := range b.Tokens {
		name += token.Name()
		name += ", "
	}

	return name
}
