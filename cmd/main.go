package main

import (
	b "app"
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

func main() {
	server := b.NewBooksServer(b.NewInMemoryBookDatabase())
	fmt.Printf("Listening on http://localhost:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), server); err != nil {
		log.Fatalf("Could not listen on port %d", PORT)
	}
}
