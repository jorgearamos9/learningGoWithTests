package main

import (
	"fmt"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	fmt.Printf("Inside GetPlayerScore, name received: %s\n", name)
	return 123
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	fmt.Printf("Inside RecordWin, name received: %s\n", name)
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
