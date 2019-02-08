package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//create a new type of deck
//wich is a slice s strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clus"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	serializedDeck := []byte(d.toString())
	return ioutil.WriteFile(filename, serializedDeck, 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
