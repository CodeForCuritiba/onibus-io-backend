package store

import (
	"github.com/codeforcuritiba/onibus-io-backend/model"
)

// Storer é uma interface que representa as estruturas que terão acesso ao banco de dados.
type Storer interface {
	Linhas() (model.Linhas, error)
	Veiculos() (model.Veiculos, error)
	Linha(codigo string) (model.Linha, error)
	Veiculo(codigo string) (model.Veiculo, error)
	VeiculosLinhas(codigoLinha string) (model.Veiculos, error)
}
