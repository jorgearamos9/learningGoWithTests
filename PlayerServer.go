package main

import (
	"net/http"
	"fmt"
	"strings"
)

// PlayerServer implemenents the http server logic 
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Printf("Got a request with URL %s, for player: %s\n",r.URL.Path,player)
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
    if name == "Pepper" {
        return "20"
    }

    if name == "Floyd" {
        return "10"
    }

    return ""
}