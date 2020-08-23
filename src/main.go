package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"ufc.com/deti/go-dad/src/migrate"
	r "ufc.com/deti/go-dad/src/routes"
)

func main() {
	mux := r.NewRouter()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"POST", "GET", "DELETE", "PUT", "PATCH"},
	})
	handler := c.Handler(mux)
	fmt.Println("Server running on 8080 port ... ")
	// database.CreateDB()
	migrate.Migrate()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
