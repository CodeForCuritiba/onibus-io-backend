package store

import (
	"github.com/codeforcuritiba/onibus-io-backend/model"
)

// Storer é uma interface que representa as estruturas que terão acesso ao banco de dados.
type Storer interface {
	Linhas() ([]*model.Linha, error)
	Linha(codigo string) (*model.Linha, error)
	Veiculos() (model.Veiculos, error)
	Veiculo(codigo string) (model.Veiculos, error)
	VeiculosLinha(codigoLinha string) ([]*model.Veiculo, error)
}
