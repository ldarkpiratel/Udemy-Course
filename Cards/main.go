package main

func main() {

	card := newDeck()
	//card := deck{"Ace of Diamonds", newCard()} // create a slice of type string using customized type deck
	//card = append(card, "Six of Spades")       // add a new string to the slice

	// range is used to iterate over a slice or array
	//for i, cards := range card { // loop over slice card, and print the index and its element
	//	fmt.Println(i+1, cards)
	//}
	card.print()
}

// explicitly state data type string, to return a string (fits for other data types)
//func newCard() string {
//	return "Five of Diamonds"
//}
