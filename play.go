package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PlayerIndex uint8

// Go version of enums
const (
	LORE PlayerIndex = iota
	Bot1
	Bot2
	Bot3
)

const (
	MAX_DISTANCE = 35 - 3
	ADJACENT     = 1
	GOAL_COST    = 12
)

// Calculates and returns a player's score
func (p Player) score() int16 {
	var streakSum int16
	for _, streak := range p.Cards {
		streakSum += int16(streak[0])
	}
	return streakSum - int16(p.Money)
}

// Prints the game state in a somewhat sane form
func (gs *GameState) printState() {
	fmt.Println("Cards left: ", gs.Status.CardsLeft)
	fmt.Println("Cards left: ", gs.Status.CardsLeft)
	fmt.Println("Money on table: ", gs.Status.Money)
	fmt.Println("Card value: ", gs.Status.Card)
	for _, p := range gs.Status.Players {
		fmt.Printf("%s: %d coins\n", p.Name, p.Money)
		fmt.Println("Score: ", p.score())
		fmt.Println("Cards:")
		for _, streak := range p.Cards {
			fmt.Println("\t", streak)
		}
		fmt.Println()
	}
	fmt.Println()
}

// Gets the card's distance from this player's closest streak.
// A streak means a consecutive set of cards.
// A returned distance of 1 is a card that fits to a streak.
func (p Player) distanceFromStreak(card int8) int8 {
	abs := func(n int8) int8 {
		if n < 0 {
			return -1 * n
		} else {
			return n
		}
	}

	var distance int8 = MAX_DISTANCE
	for _, streak := range p.Cards {
		low := streak[0]
		high := streak[len(streak)-1]
		distance = min(min(abs(low-card), distance), min(abs(high-card), distance))
	}
	return distance
}

// Tells if the card is 'urgent' to pick for 'player'. A card is marked urgent
// if the other players have an urgent reason to pick it, usually a fitting
// streak. An urgent card should be picked as soon as possible if it's
// important for LORE's streak.
func (player Player) isCardUrgent(gs *GameState) bool {
	for _, p := range gs.Status.Players {
		// Skip self
		if &p == &player {
			continue
		}
		if p.distanceFromStreak(gs.Status.Card) == ADJACENT {
			return true
		}
	}
	return false
}

// Calculates the 'cost' of taking the card on this turn.
func (player Player) getCost(gs *GameState) int8 {
	return gs.Status.Card - int8(gs.Status.Money)
}

// Decides whether the player 'p' should take the card (true)
// or not (false).
// The strategy of this algorithm is to aim for the middle: cards above 9 and
// below 21. The goal is to build some streaks with reasonable risk in the
// high-20 range, while getting to poach some coins from the other players.
// The idea is that the low cards go fast and can't be obtained as easily for
// streaks, and high cards are just a risk even if they could be utilized to
// steal more coins.
func (p Player) decidePlay(gs *GameState) bool {
	// No money, no honey. Got to take.
	if p.Money == 0 {
		fmt.Println("NO MONEY, money: ", p.Money)
		return true
	}

	// A good 'cost' for a card is <= 12. Calculated as card-money.
	cost := p.getCost(gs)

	// Unless a card fits to a streak...
	distance := p.distanceFromStreak(gs.Status.Card)
	if distance == ADJACENT {
		fmt.Println("ADJACENT, distance: ", distance)
		return true
	}
	// ...we simply skip every card outside of the range [10,20].
	if gs.Status.Card < 10 || gs.Status.Card > 20 {
		fmt.Println("OUT OF BOUNDS, Card: ", gs.Status.Card)
		return false
	}

	// We take the ones within our range that are cheap enough.
	if cost <= GOAL_COST {
		fmt.Println("BELOW GOAL, cost: ", cost)
		return true
	}
	// Otherwise just bet.
	return false
}

// Update the game state with the response, returning whether the game
// is finished.
func (gs *GameState) update(res *http.Response) bool {
	if err := json.NewDecoder(res.Body).Decode(&gs); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("New state in Update:\n", gs)
	return gs.Status.Finished
}

// Method on a game instance that plays a single turn, i.e. decides on the
// action, sends it to server and updates the game state with the response.
// Returns the information whether the game has concluded.
func (gs *GameState) playTurn() {

	// Inspect game state and decide on the correct play
	takeCard := gs.Status.Players[LORE].decidePlay(gs)
	fmt.Println("Decision: ", takeCard)

	// POST the decision to server and return the response
	response := postAction(takeCard, URL+"/game/"+gs.GameId+"/action")
	defer response.Body.Close()

	// Update the game state
	gs.update(response)
}
