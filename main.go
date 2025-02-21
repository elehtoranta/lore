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
	URL     = "https://koodipahkina.monad.fi/api"
	API_KEY = "LORE_API_KEY"
)

func main() {
	// This will not override any keys, so an API key loaded to the test process via `env` command
	// will be used instead of anything in .env. Github CI tests should be using a mock key.
	if len(os.Args) <= 1 {
		fmt.Println("Please give the number of games to play as an argument.")
		return
	}

	// The number of games played is given as the only argument.
	nGames, err := strconv.Atoi(os.Args[1])
	if err != nil || nGames < 1 || nGames > 100 {
		fmt.Println("Please give a number between 1-100 as the number of games.")
		return
	}

	// Look for API_KEY in environment, and if missing, otherwise load from .env.
	if _, found := os.LookupEnv(API_KEY); !found {
		fmt.Printf("No %s in environment, searching from .env\n", API_KEY)
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Create a channel for n games
	gameChan := make(chan GameState, nGames)

	// Play the games in goroutines
	for i := 0; i < nGames; i++ {
		fmt.Println("Playing game #", i)
		go playGame(gameChan, cap(gameChan), i)
	}
	// Print state for all finished games received from channel.
	for game := range gameChan {
		game.printState()
	}
}
