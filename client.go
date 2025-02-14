package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// For marsalling (JSON encoding) an action sent to the server
type Action struct {
	TakeCard bool `json:"takeCard"`
}

// Requests a new game and returns its ID on success.
func initGame(url string, gs *GameState) (string, error) {
	var buf io.Reader

	client := http.DefaultClient

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}
	key, found := os.LookupEnv("LORE_API_KEY")
	if !found {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+key)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&gs); err != nil {
		return "", err
	}
	return gs.GameId, nil
}

func postAction(action bool, url string) *http.Response {
	a := Action{TakeCard: action}
	msg, err := json.Marshal(a)

	req, err := http.NewRequest("POST", url, bytes.NewReader(msg))
	if err != nil {
		log.Fatal(err)
	}

	key, found := os.LookupEnv("LORE_API_KEY")
	if !found {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
