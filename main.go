package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

// Game state populated by API responses.
type GameState struct {
    GameId string `json:"gameId"`
    Status Status `json:"status"`
}

type Status struct {
    Card uint8 `json:"card"`
    Money uint8 `json:"money"`
    Players []Player `json:"players"`
    CardsLeft uint8 `json:"cardsLeft"`
    Finished bool `json:"finished"`
}

type Player struct {
    Name string `json:"name"`
    Money uint16 `json:"money"`
    Cards [][]uint16 `json:"cards"`
}

const (
    URL = "https://koodipahkina.monad.fi/api"
)

// Requests a new game and returns its ID on success, and writes error on failure.
func initGame(url string, gs *GameState) (string, error) {
    var buf io.Reader

    client := http.DefaultClient

    req, err := http.NewRequest("POST", url, buf)
    if err != nil {
        log.Fatal(err)
    }
    key, found := os.LookupEnv("LORE_API_KEY")
    if !found {
        log.Fatal(err)
    }
    req.Header.Add("Authorization", "Bearer "+key)

    res, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    err = json.NewDecoder(res.Body).Decode(&gs)
    if err != nil {
        return "", err
    }
    return gs.GameId, nil
}

// ping the server for a happy and stress free start to the project
func ping() string {
    log.Println("Pinging server")
    resp, err := http.Get(URL)
    if err != nil {
        log.Fatal(err)
    }
    if resp.StatusCode != 200 {
        log.Fatal("Server returned status code " + fmt.Sprint(resp.StatusCode))
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Can't read response body.")
    }
    return string(body)
}

func main() {
    // This will not override any keys, so an API key loaded to the test process via `env` command
    // will be used instead of anything in .env. Github CI tests should be using a mock key. Should. :D
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    var gs GameState
    gameId, err := initGame(URL+"/game", &gs)
    if err != nil || gameId == "" {
        log.Fatal("Game initialization returned error.", err)
    }
    // fmt.Println("Ping:", ping())
}
