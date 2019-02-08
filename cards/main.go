package main

func main() {
	cards := newDeck()

	hand, remainigCards := deal(cards, 5)

	hand.print()
	remainigCards.print()

	hand.saveToFile("hand.txt")

}

func newCard() string {
	return "Five of Diamonds"
}
