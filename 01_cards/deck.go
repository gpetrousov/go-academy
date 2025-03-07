package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
type deck []string

// Create a receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
    // Returns two different decks split by handSize
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
    // Instantiate a new deck of cards

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func check_error(e error) {
    if e != nil {
        panic(e)
    }
}

func (d deck) saveToFile(filename string) {
    content := d.toString()
    //fmt.Println(content)
    err := os.WriteFile(filename, []byte(content), 0664)
    check_error(err)
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
