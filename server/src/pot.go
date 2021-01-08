package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Pot struct {
	Spaces []*PotSpace
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

	for _, space := range p.Spaces {
		name += space.Name()
		name += " | "
	}

	return name
}
