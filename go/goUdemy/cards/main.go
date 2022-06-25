package main

import "fmt"

func main(){
	filename := "deckfile.txt"
	deck := newDeck()
	if err := deck.saveToFile(filename); err != nil {
		fmt.Println("Some error occured:- ", err)
	}

	// deck.print()

	// hand, remainingDeck := deal(deck, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(deck.deckToString())

	d := readDeckFromFile(filename)
	d.print()

	d.shuffle()
}

func newCard() string{
	return "Ace of star"
}
