package main

import (
	"strconv"
)

type Token struct {
	IngredientType string
	Value 		   int
}

func NewToken(ingredientType string, value int) *Token {
	return &Token{ingredientType, value}
}

func (token *Token) Name() string {
	return token.IngredientType + " " + strconv.Itoa(token.Value)
}