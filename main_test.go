package main

import (
    "testing"
)

func TestPing(t *testing.T) {
    want := "{\"ok\":true}"
    res := ping()
    if res != want {
        t.Fatalf(`Pinging failed, response body was: %s instead of %s.`, res, want)
    }
}

