package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
	"errors"
)

// Size of a deck of cards
const deckSize = 52;

// All cvalues and suits of a deck of cards
var values = []string{"1","2","3","4","5","6","7","8","9","10","J","Q","K"}
var suits = []string{"D","C","H","S"}

// Card has a value and a suit
type Card struct {
	Value string
	Suit string
}

// Get the index of a particular suit
func (c Card) suitIndex() (int, error) {
	var index *int
	for i,s := range suits {
		if s == c.Suit {
			index = &i
			break
		}
	}
	if index == nil {
		return 0,errors.New(fmt.Sprintf("Invalid suit %s", c.Suit))
	}

	return *index, nil
}

// Get the index of a particular vaule
func (c Card) valueIndex() (int, error) {
	var index *int
	for i,s := range values {
		if s == c.Value {
			index = &i
			break
		}
	}
	if index == nil {
		return 0,errors.New(fmt.Sprintf("Invalid value %s", c.Value))
	}

	return *index, nil
}

// Check if two cards are equal
func (c Card) equal(card Card) bool {
	if ((c.Value == card.Value) &&
	    (c.Suit == card.Suit)) {
	    return true
        }
	return false
}

// Get the numerically linear representation of a card
func (c Card) orderedValue() (int) {
	card_suit_index, err := c.suitIndex()
	if err != nil { panic(err) }

	card_value_index, err := c.valueIndex()
	if err != nil { panic(err) }
	
	return ((card_suit_index*len(values)) + card_value_index)
}

// Check if one card is greater than another
func (c Card) Less(card Card) (bool) {
	current_card_value := c.orderedValue()
	arg_card_value := card.orderedValue()
	return current_card_value < arg_card_value
}

// String represenation of acard
func (c Card) toString() string {
	return c.Value + c.Suit
}


// A deck is an array of cards
type Deck struct {
	Cards []Card
}

// Lenght of a deck is the size of the array of cards
func (d Deck) Len() int { return len(d.Cards) }

// Swap two cards in a deck
func (d Deck) Swap(i,j int) { d.Cards[i],d.Cards[j] = d.Cards[j],d.Cards[i] }

// Check if one card is less than another
func (d Deck) Less(i,j int) (bool) {
	return d.Cards[i].Less(d.Cards[j])
}

// Get the index of a card in a deck
func (d Deck) Index(card Card) int {
	for i,deckCard := range d.Cards {
		if deckCard.equal(card) {
			return i
		}
	}
	return -1
}

// remove card from deck and return
func (d *Deck) GetCard(index int) Card {
	return_card := d.Cards[index]
	if index == d.Len() {
		d.Cards = append(d.Cards[:d.Len()-1])
	} else if index == 0 {
		d.Cards = d.Cards[index+1:]
	} else {
		d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	}
	return return_card
}

// Add card to the deck
func (d *Deck) AddCard(index int, card Card) {
	if index == 0 {
		d.Cards = append([]Card{card}, d.Cards...)
	} else if index == d.Len() {
		d.Cards = append(d.Cards[:d.Len()], []Card{card}...)
	} else {
		d.Cards = append(d.Cards[:index], append([]Card{card}, d.Cards[index:]...)...)
	}
}

// Create a new deck of cards
func NewDeck() (deck Deck) {
	var cards []Card
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, Card{value,suit})
		}
	}
	return Deck{cards}
}

// Shuffle a deck of cards
func Shuffle(deck Deck) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(deck.Cards) > 0 {
		n := len(deck.Cards)
		randIndex := r.Intn(n)
		deck.Cards[n-1], deck.Cards[randIndex] = deck.Cards[randIndex], deck.Cards[n-1]
		deck.Cards = deck.Cards[:n-1]
	}
}

func main() {

	// 1. Create a deck of cards and print all sorted spades
	var deck = NewDeck()
	var counter int
	for _,card := range deck.Cards {
		if (card.Suit == "S" && (counter < 5)) {
			fmt.Printf(card.toString() + " ")
			counter++
		}
	}
	fmt.Println()
	
	// 2. Shuffle cards and print 5 random cards
	deck = NewDeck()
	Shuffle(deck)
	for i,card := range deck.Cards {
		if i < 5 {
			fmt.Printf(card.toString() + " ")
		}
	}
	fmt.Println()
	
	// 3. Sort a shuffled deck of cards
	deck = NewDeck()
	Shuffle(deck)
	sort.Sort(deck)
	for _,card := range deck.Cards {
		fmt.Printf(card.toString() + " ")
	}
	fmt.Println()
		
	// 4. Compare random cards against eachother by [spades, hearts, clubs, diamonds] in that order
	deck = NewDeck()
	Shuffle(deck)
	card1, card2, sign := deck.Cards[rand.Intn(deckSize)], deck.Cards[rand.Intn(deckSize)], ""
	if card1.Less(card2) {
		sign = "<"
	} else if card2.Less(card1) {
		sign = ">"
	} else {
		sign = "="
	}
	fmt.Println(card1.toString(), sign, card2.toString())

	// 5. Cut a deck of cards
	deck = NewDeck()
	Shuffle(deck)
	split := rand.Intn(deckSize)
	deck.Cards = append(deck.Cards[split:], deck.Cards[:split]...)
	
	// 6. simulate a simple game of "war"
	deck = NewDeck()
	Shuffle(deck)
	split = int(deckSize / 2)
	player1_deck, player2_deck := &Deck{deck.Cards[split:]}, &Deck{deck.Cards[:split]}

	iters := 0
	for ((player1_deck.Len() != 0) && (player2_deck.Len() != 0)) {

		// Count iters
		iters++

		// Get player cards
		player1_card, player2_card := player1_deck.GetCard(0), player2_deck.GetCard(0)

		// determine winner
		if player1_card.orderedValue() < player2_card.orderedValue() {
			player2_deck.AddCard(player2_deck.Len(), player1_card)
			player2_deck.AddCard(player2_deck.Len(), player2_card)
		} else if player1_card.orderedValue() > player2_card.orderedValue() {
			player1_deck.AddCard(player1_deck.Len(), player2_card)
			player1_deck.AddCard(player1_deck.Len(), player1_card)
		} else {
			player1_deck.AddCard(player1_deck.Len(), player1_card)
			player2_deck.AddCard(player2_deck.Len(), player2_card)
		}
	}
	fmt.Println(fmt.Sprintf("Player 1: %d Player 2: %d Iters: %d", player1_deck.Len(), player2_deck.Len(), iters))
}
