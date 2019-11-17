package store

import (
	"context"

	"github.com/codeforcuritiba/onibus-io-backend/config"
	"github.com/codeforcuritiba/onibus-io-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStore é uma store que se comunica com uma base de dados mongodb.
type MongoStore struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

// NewMongoStore cria uma Store para uma base de dados mongodb.
func NewMongoStore(ctx context.Context, client *mongo.Client, config config.Configurer) (store Storer) {
	store = &MongoStore{client: client, db: client.Database(config.DBName())}
	return
}

// Linhas retorna uma lista de linhas sem seus pontos e suas tabelas.
func (ms *MongoStore) Linhas() (linhas model.Linhas, err error) {
	projection := map[string]int{
		"_id":    0,
		"pontos": 0,
		"tabela": 0,
	}
	cur, err := ms.db.Collection("linhas").Find(ms.ctx, bson.D{}, &options.FindOptions{Projection: projection})
	if err != nil {
		return
	}
	for cur.Next(ms.ctx) {
		var linha model.Linha
		err = cur.Decode(&linha)
		if err != nil {
			linhas = model.Linhas{}
			return
		}
		linhas = append(linhas, linha)
	}
	return
}

// Linha busca uma linha através do código
func (ms *MongoStore) Linha(codigo string) (linha model.Linha, err error) {
	err = ms.db.Collection("linhas").FindOne(ms.ctx, map[string]string{"cod": codigo}).Decode(&linha)
	return
}

// Linhas retorna uma lista de veiculos
func (ms *MongoStore) Veiculos() (veiculos model.Veiculos, err error) {

	cur, err := ms.db.Collection("veiculos").Find(ms.ctx, bson.D{})
	if err != nil {
		return
	}
	for cur.Next(ms.ctx) {
		var veiculo model.Veiculo
		err = cur.Decode(&veiculo)
		if err != nil {
			veiculos = model.Veiculos{}
			return
		}
		veiculos = append(veiculos, veiculo)
	}
	return
}
