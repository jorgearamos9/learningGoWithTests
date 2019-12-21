package main

import (
	"net/http"
	"fmt"
	"strings"
)

type PlayerStore interface {
    GetPlayerScore(name string) int
} 

// PlayerServer implemenents the http server logic 
type PlayerServer struct {
    store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Printf("Got a request with URL %s, for player: %s\n",r.URL.Path,player)
	
	score := p.store.GetPlayerScore(player)

    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

    fmt.Fprint(w, score)
}
