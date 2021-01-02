package main

import "fmt"

var (
	bag = NewBag()
)

func main() {
	fmt.Println(bag.Name())
	bag.AddToken(NewToken("White", 2))
    fmt.Println(bag)
}