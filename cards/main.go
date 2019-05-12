package main

func main() {

	// cards := newDeck()
	// cards.savetofile("my_cards")
	// hand, remainingcard := deal(cards, 5)
	// hand.print()
	// remainingcard.print()
	// cards := newDeckFromfile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newcard() string { //by convention variable name of receiver variable is one or two letter starting with type of variable i.e. here its d

	return "King of hearts"
}
