package main

import (
	"fmt"
	io "io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Diamond", "Heart", "Spade", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cardDeck := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cardDeck = append(cardDeck, value+" Of "+suit)
		}
	}

	return cardDeck
}

func (d deck) print() {
	fmt.Println("\nPrinting deck:-")
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deckToString() string {
	return strings.Join(d, ", ")

}

func (d deck) saveToFile(filename string) error {
	return io.WriteFile(filename, []byte(d.deckToString()), 0666)
}

func readDeckFromFile(filename string) deck {
	var d deck

	if bs, error := io.ReadFile(filename); error != nil {
		fmt.Println("Some error occured:- ", error)
		os.Exit(1)
	} else {
		return strings.Split(string(bs), ",")
		// for _, card := range strings.Split(string(bs), ","){
		// 	d = append(d, card)
		// }
	}

	return d

}

func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)

		r := random.Intn(len(d) - 1)
		fmt.Println(i, r)
	}
}
