package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jtyler258/mbp/api/db"
	"github.com/jtyler258/mbp/api/gqlgen"
	"github.com/jtyler258/mbp/api/resolver"
	"github.com/jtyler258/mbp/api/service"
)

const defaultPort = "8080"
const defaultDBConnection = "mongodb://db/mbp"

func main() {
	dbConnection := os.Getenv("DATABASE_CONNECTION")
	if dbConnection == "" {
		dbConnection = defaultDBConnection
	}

	database, err := db.NewDatabase(dbConnection)
	if (err != nil) {
		log.Fatal(err)
	}

	err = database.Connect()
	if (err != nil) {
		log.Fatal(err)
	}

	defer database.Disconnect(context.Background())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolver.Resolver{
		RecipeResolver: &resolver.RecipeResolver{
			RecipeService: service.NewRecipeService(db.NewRecipeRepository(database)),
		},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
