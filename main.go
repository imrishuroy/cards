package main

func main() {

	// var card string = "Ace of Spades"

	// card := "Ace of Spades"
	// note - walrus operator := is only used when declaring/initializing a new variable
	// card = "Five of Diamonds"

	// card := newCard()
	// fmt.Println(card)

	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}

	// cards = append(cards, "Six of Spades")

	//cards := newDeck()
	//fmt.Println([]byte(cards.toString()))

	//cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println("")
	// print("Remaining Deck: ", cards)
	// fmt.Println("")
	// remainingDeck.print()

	//cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
