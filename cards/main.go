package main

func main() {

	// cards := newDeck()
	cards := newDeckFromFile("my_cards.txt")
	cards.print()
	// hand, remainingCards := deal(cards, 4)
	//
	// hand.print()
	// fmt.Println("====================")
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.txt")
}
