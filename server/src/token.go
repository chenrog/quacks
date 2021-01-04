package main

import "strconv"

type Token struct {
	IngredientType string
	Value          int
}

func NewToken(ingredientType string, value int) *Token {
	return &Token{ingredientType, value}
}

func (t *Token) Name() string {
	return t.IngredientType + "-" + strconv.Itoa(t.Value)
}
