package store

import (
	"context"

	"github.com/codeforcuritiba/onibus-io-backend/config"
	"github.com/codeforcuritiba/onibus-io-backend/core"
	"github.com/codeforcuritiba/onibus-io-backend/core/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoStore é uma store que se comunica com uma base de dados mongodb.
type MongoStore struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

// NewMongoStore cria uma Store para uma base de dados mongodb.
func NewMongoStore(ctx context.Context, client *mongo.Client, config config.Configuration) (store core.LinhaStorer) {
	store = &MongoStore{client: client, db: client.Database(config.MongoDB.DBName)}
	return
}

// Linhas retorna uma lista de linhas sem seus pontos e suas tabelas.
func (ms *MongoStore) Linhas() (linhas []*model.Linha, err error) {
	cur, err := ms.db.Collection("linhas").Find(ms.ctx, bson.D{})
	if err != nil {
		return
	}
	for cur.Next(ms.ctx) {
		var linha *model.Linha
		err = cur.Decode(&linha)
		if err != nil {
			linhas = []*model.Linha{}
			return
		}
		linhas = append(linhas, linha)
	}
	return
}

// Linha busca uma linha através do código
func (ms *MongoStore) Linha(codigo string) (linha *model.Linha, err error) {
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

// Veiculo retorna uma lista de veiculo por codigo
func (ms *MongoStore) Veiculo(codigo string) (veiculos model.Veiculos, err error) {

	cur, err := ms.db.Collection("veiculos").Find(ms.ctx, bson.M{"cod": codigo})
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

// VeiculosLinha retorna uma lista dos veiculos de uma linha
func (ms *MongoStore) VeiculosLinha(codigo string) (veiculos []*model.Veiculo, err error) {

	cur, err := ms.db.Collection("veiculos").Find(ms.ctx, bson.M{"codigolinha": codigo})
	if err != nil {
		return
	}
	for cur.Next(ms.ctx) {
		var veiculo model.Veiculo
		err = cur.Decode(&veiculo)
		if err != nil {
			veiculos = []*model.Veiculo{}
			return
		}
		veiculos = append(veiculos, &veiculo)
	}
	return
}

func (ms *MongoStore) TabelaLinha(codigoLinha,numeroPonto string) ([]*model.Parada, error) {
	pipeline := mongo.Pipeline{

			bson.D{{"$match",bson.D{{"cod",codigoLinha}}}},
			bson.D{{"$unwind", bson.M{
					"path": "$tabela",
					"includeArrayIndex": "string",
					"preserveNullAndEmptyArrays": false,
			},
			}},
			bson.D{{"$match", bson.M{
				"tabela.num": numeroPonto,
				}},
			},
			bson.D{{"$replaceRoot", bson.D{{
					"newRoot", "$tabela",
			}},
			}},
		}
	
	cur , err := ms.db.Collection("linhas").Aggregate(ms.ctx,pipeline);
	if err != nil {
		return nil, err
	}
	var tabela []*model.Parada
	for cur.Next(ms.ctx){
		var parada model.Parada
		cur.Decode(&parada)
		tabela = append(tabela,&parada)
	}
	return tabela, nil
}
