package main

import "fmt"

func main() {

	cards := newDeck()
    hand, remainingCards := deal(cards, 5)
	fmt.Println("Original hand\n======================")
    hand.print()

	fmt.Println("\nRemaining deck\n======================")
    remainingCards.print()

    cards.saveToFile("cards_file.csv")

	// cards := newDeckFromFile("my_cards.txt")
	// fmt.Println("Original hand\n======================")
	// cards.print()
	// cards.shuffle()
	// fmt.Println("Shuffled hand\n======================")
	// cards.print()

	// hand, remainingCards := deal(cards, 4)
	//
	// hand.print()
	// fmt.Println("====================")
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.txt")
}
