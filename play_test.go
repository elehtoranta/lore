package main

import (
	"testing"
)

// Test objects
var TESTSTATE = GameState{
	GameId: "Foo",
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

var TESTPLAYER = Player{
	"Foo",
	42,
	[][]int8{
		[]int8{1, 2, 3},
		[]int8{11, 12, 13},
		[]int8{30},
		[]int8{34},
	},
}

func TestDistanceFromStreak(t *testing.T) {
	is := TESTPLAYER.distanceFromStreak(10)
	not := TESTPLAYER.distanceFromStreak(32)

	if is != ADJACENT {
		t.Errorf("Distance from streak failed for IS:\nwanted: %d\ngot: %d", ADJACENT, is)
	}
	if not != 2 {
		t.Errorf("Distance from streak failed for NOT:\nwanted: %d\ngot: %d", 2, not)
	}
}
