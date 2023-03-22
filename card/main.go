package main

// import "fmt"

func main() {
	// Creating decks
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("\n")
	// remainingCards.print()

	// // saving the deck into a file
	// fmt.Println(cards.toString())
	// cards.saveToFile("text.txt")

	// // Reading decks from a file
	// cards := newDeckFromFile("text.txt")
	// cards.print()

	// shuffling cars
	cards.shuffle()
	cards.print()
}
