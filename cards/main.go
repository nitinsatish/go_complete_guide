package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	strval := cards.toString()
	fmt.Println("cards as string:", strval)
	cards.toFile("cards.txt")
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining Deck:")
	remainingDeck.print()
	// cards.print()
	sdeck := cards.shuffle()
	fmt.Println("Shuffled deck :")
	sdeck.print()
	fdeck := deckFromFile("cards.txt")
	fmt.Println("Deck from File:")
	fdeck.print()

}

func newCard() string {
	return "Ace of Hearts"
}
