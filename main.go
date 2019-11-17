package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codeforcuritiba/onibus-io-backend/config"
	"github.com/codeforcuritiba/onibus-io-backend/db"
	"github.com/codeforcuritiba/onibus-io-backend/router"
	"github.com/codeforcuritiba/onibus-io-backend/store"
)

func main() {
	config := &config.EnvConfigurer{}

	log.Println("Criando cliente mongodb...")
	ctx := context.Background()
	client, err := db.NewMongoClient(ctx, config.DBStrConn())
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %q", err)
		os.Exit(1)
	}
	log.Println("Cliente mongodb criado com sucesso")

	log.Println("Criando store...")
	s := store.NewMongoStore(ctx, client, config)
	log.Println("Store criada com sucesso")

	log.Println("Criando rotas...")
	r := router.NewRouter(s)
	log.Println("Rotas criadas")
	log.Printf("Iniciando servidor na porta %s\n", config.Port())
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port()), r)

}
