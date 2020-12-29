package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// any variable of type 'deck' gets access to the 'print()' method
// 'd' is a reference to the instance of type 'deck' variable
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// take first argument of type deck, and a second argument of type int
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function
func (d deck) toString() string {
	// convert slice type deck to type string then join every element in a slice into 1 string
	// the strings.Join() function receives the slice as the first argument, then the separation
	// character after joining the slice
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	// ioutil.WriteFile() requires 3 arguments: the name of the file to create, the byte slice and the permissions for the file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 - log the error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return deck(s)

}

func (d deck) shuffle() {

	// generate new source (seed) value for random numbers to be generated
	// use time package to use type Int64 as argument to the function
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// pseudorandom number generator
		// get random index from deck slice
		newPosition := r.Intn(len(d) - 1)

		// shuffle elements on slice to new positions using random index generated
		// this swaps elements in the slice
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
