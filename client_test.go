package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	// "github.com/go-test/deep"
)

const (
	mockAPIAuthHeader = "Bearer bar"
	gameId            = "foo"
	initResponse      = `{"gameId":"foo","status":{"card":10,"money":0,"players":[{"money":11,"cards":[],"name":"ure"},{"money":11,"cards":[],"name":"KoodiKonna-77"},{"money":11,"cards":[],"name":"PähkinäPomo X10"},{"money":11,"cards":[],"name":"BittiBotti-6000"}],"cardsLeft":24,"finished":false}}`
)

func TestInitGame(t *testing.T) {

	// Create a local mock server to handle requests.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: The mock API token has to be present in the environment variables for example
		// via `env` call. Without it present, the .env file will be searched for the actual key.
		if r.Header.Get("Authorization") != mockAPIAuthHeader {
			t.Errorf("Expected 'Authorization: Bearer' header %s, got %s", mockAPIAuthHeader, r.Header.Get("Authorization"))
		}
		// Respond
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(initResponse))
	}))
	defer ts.Close()

	var gs GameState
	_, err := initGame(ts.URL, &gs)
	if err != nil {
		t.Fatalf("Game initialization error.")
	}

	// The sought after struct state after above response
	want := GameState{
		GameId: gameId,
		Status: Status{
			Card:  10,
			Money: 0,
			Players: []Player{
				{Money: 11, Name: "ure", Cards: [][]int8{}},
				{Money: 11, Name: "KoodiKonna-77", Cards: [][]int8{}},
				{Money: 11, Name: "PähkinäPomo X10", Cards: [][]int8{}},
				{Money: 11, Name: "BittiBotti-6000", Cards: [][]int8{}},
			},
			CardsLeft: 24,
			Finished:  false,
		},
	}

	if !cmp.Equal(want, gs) {
		t.Error("Decoded game state structs do not match:\nwant:\t", want, "\nwas:\t", gs)
	}
}

func TestPostAction(t *testing.T) {

}
