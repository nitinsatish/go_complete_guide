package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	var cards deck
	cardSuits := []string{"Diamond", "Clubs", "Hearts", "Spade"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handCount int) (deck, deck) {
	return d[:handCount], d[handCount:]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) toFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)

}

func (d deck) shuffle() deck {
	for i := range d {
		randno := rand.Intn(len(d) - 1)
		d[i], d[randno] = d[randno], d[i]
	}
	return d
}

func deckFromFile(f string) deck {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}
