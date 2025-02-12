package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://koodipahkina.monad.fi/api"

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
    fmt.Println("Ping:", ping())
}
