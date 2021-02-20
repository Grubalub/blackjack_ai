package blackjack

import (
	deck "Grubalub/Deck"
	"fmt"
)

const (
	statePlayerTurn State = iota
	stateDealerTurn
	stateHandOver
)

type state int8

func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
}

type Game struct {
	// unexported fields
	deck     []deck.Card
	state    state
	Player   []deck.Card
	Dealer   []deck.Card
	dealerAI AI
	balance  int
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.State {
	case StatePlayerTurn:
		return &g.player
	case StateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any players turn")
	}
}

func deal(gs Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.Deck = draw(g.deck)
		g.dealer = append(g.dealer, card)

	}
	g.State = statePlayerTurn

}

func (g *Game) Play(ai AI) int {
	gs.deck = deck.New(deck.Deck(3), deck.Shuffle)

	for i := 0; i < 2; i++ {
		deal(g)

		for g.state == statePlayerTurn {
			han := make([]deck.Card, len(g.Player))
			copy(hand, g.player)
			move := ai.Play(g.player, g.dealer[0])
			move(g)
		}

		for g.state == StateDealerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endHand(g, ai)
	}
	return 0
}

type Move func(*Game)

func MoveHit(g *Game) {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
	return ret
}

func MoveStand(gs *Game) {
	g.state++
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func endHand(g *Game, ai AI) {

	pScore, dScore := Score(g.player...), Score(g.dealer...)
	//todo : figure out winnings
	switch {
	case pScore > 21:
		fmt.Println("You bust fool!")
		g.balance--
	case dScore > 21:
		fmt.Println("The deala bust fool!")
		g.balance++
	case pScore > dScore:
		fmt.Println("You won sucka!")
		g.balance++
	case pScore < dScore:
		fmt.Println("You lost sucka!")
		g.balance--
	case pScore == dScore:
		fmt.Println("It's a draw fool")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil

}

//score will take in a hand and return the best possible blackjack hand
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, we are changing it to be worth 11
			return minScore + 10
		}
	}
	return minScore
}

// soft returns true if the score of a ahand is a soft score
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}