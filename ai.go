package blackjack

import (
	deck "Grubalub/Deck"
	"fmt"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Result(hand [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct {}

func (ai dealerAI) Bet() int {
	// nothing
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
		if dScore() <= 16 || (dScore() == 17 && Soft(hand...)) {
	return MoveHit
		}
			return MoveStand
		}



func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
// nothing
}

func HumanAI() AI {
	return humanAI{}
}
type HhmanAI struct{}

func (ai humanAI) Bet() int {
	return 1
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
	fmt.Println("Players:", hand)
	fmt.Println("Dealer:", dealer)
	fmt.Println("What do you want to do? (H)it, (S)tand")
		var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "h":
		return MoveHit
	case "s":
		return MoveStand
	default:
		fmt.Println("Invald option:", input)
	}
}

func (ai humanAI) Results(hand [][]deck.Card, dealer[]deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Players:", hand)
	fmt.Println("Dealer:", dealer)
}


// filler for later 


