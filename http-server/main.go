package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on part 5000 %v", err)
	}
}
