package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 4)

	hand.print()
	fmt.Println("====================")
	remainingCards.print()

}
