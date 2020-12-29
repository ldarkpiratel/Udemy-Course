package main

func main() {

	//card := newDeck()
	//card := deck{"Ace of Diamonds", newCard()} // create a slice of type string using customized type deck
	//card = append(card, "Six of Spades")       // add a new string to the slice

	// range is used to iterate over a slice or array
	//for i, cards := range card { // loop over slice card, and print the index and its element
	//	fmt.Println(i+1, cards)
	//}
	//card.print()
	//hand, remainingCards := deal(card, 5)

	//hand.print()
	//remainingCards.print()

	//greeting := "Hi there!"

	// convert string to byte slice
	//fmt.Println([]byte(greeting))

	//cards := newDeck()
	//fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// explicitly state data type string, to return a string (fits for other data types)
//func newCard() string {
//	return "Five of Diamonds"
//}
