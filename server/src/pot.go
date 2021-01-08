package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type potToken struct {
	token    Token
	potIndex int
}

func NewPotToken(token Token, potIndex int) *potToken {
	return &potToken{token, potIndex}
}

type Pot struct {
	PotSpaces []*PotSpace
	potTokens []*potToken
}

func NewPot() *Pot {
	return &Pot{}
}

func NewPotFromJson(jsonPath string) *Pot {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pot Pot
	json.Unmarshal([]byte(byteValue), &pot)

	return &pot
}

func (p *Pot) Name() string {
	name := "Pot: "

	// for _, space := range p.PotSpaces {
	// 	name += space.Name()
	// 	name += " | "
	// }

	// name += " ||| "

	for _, token := range p.potTokens {
		name += token.token.Name()
		name += " on "
		name += strconv.Itoa(token.potIndex)
		name += " | "
	}

	return name
}

func (p *Pot) PlaceToken(token Token) {
	index := -1
	if p.potTokens != nil {
		index = p.potTokens[len(p.potTokens)-1].potIndex
	}

	p.potTokens = append(p.potTokens, NewPotToken(token, index+token.Value))
}
