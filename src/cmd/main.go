package main

import (
	server "app/src"
	models "app/src/models"
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

func main() {
	server := server.NewBooksServer(models.NewPGBookDatabse())
	fmt.Printf("Listening on http://localhost:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), server); err != nil {
		log.Fatalf("Could not listen on port %d", PORT)
	}
}
