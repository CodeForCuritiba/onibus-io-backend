package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/codeforcuritiba/onibus-io-backend/config"
	"github.com/codeforcuritiba/onibus-io-backend/resolver"
	"github.com/codeforcuritiba/onibus-io-backend/db"
	"github.com/codeforcuritiba/onibus-io-backend/graphql"
	"github.com/codeforcuritiba/onibus-io-backend/store"
	"github.com/luizvnasc/gonfig"
)

func main() {
	config := config.Configuration{}
	gonfig.Load(&config)

	log.Println("Criando cliente mongodb...")
	ctx := context.Background()
	client, err := db.NewMongoClient(ctx, config.MongoDB.StrConn)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %q", err)
		os.Exit(1)
	}
	log.Println("Cliente mongodb criado com sucesso")

	log.Println("Criando store...")
	s := store.NewMongoStore(ctx, client, config)
	log.Println("Store criada com sucesso")

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver.Resolver{s}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))

}
