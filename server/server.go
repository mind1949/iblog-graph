package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	iblog "github.com/mind1949/iblog"
	"github.com/mind1949/iblog/db"
	"github.com/mind1949/iblog/resolvers"
)

const defaultPort = "8080"

func main() {
	defer db.PGDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(iblog.NewExecutableSchema(iblog.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
