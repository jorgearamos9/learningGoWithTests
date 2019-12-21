package main

import (
	"net/http"
	"fmt"
)

// PlayerServer implemenents the http server logic 
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}