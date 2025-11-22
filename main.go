package main

import (
	"fmt"

	"github.com/jason870721/GGPoker/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
