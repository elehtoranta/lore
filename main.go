package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Game state populated by API responses.
type GameState struct {
	GameId string `json:"gameId"`
	Status Status `json:"status"`
}

type Status struct {
	Card      int8     `json:"card"`
	Money     uint8    `json:"money"`
	Players   []Player `json:"players"`
	CardsLeft uint8    `json:"cardsLeft"`
	Finished  bool     `json:"finished"`
}

type Player struct {
	Name  string   `json:"name"`
	Money uint16   `json:"money"`
	Cards [][]int8 `json:"cards"`
}

const (
	URL = "https://koodipahkina.monad.fi/api"
)

func main() {
	// This will not override any keys, so an API key loaded to the test process via `env` command
	// will be used instead of anything in .env. Github CI tests should be using a mock key.
	if len(os.Args) <= 1 {
		fmt.Println("Please give the number of games to play as an argument.")
		return
	}

	// The number of games played is given as the only argument.
	playNGames, err := strconv.Atoi(os.Args[1])
	if err != nil || playNGames < 1 || playNGames > 100 {
		fmt.Println("Please give a number between 1-100 as the number of games.")
		return
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for i := 0; i < playNGames; i++ {
		var gs GameState
		gameId, err := initGame(URL+"/game", &gs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Game ID: ", gameId)

		for !gs.Status.Finished {
			gs.playTurn()
			gs.printState()
			fmt.Println("---------NEW TURN----------")
		}
	}
}
