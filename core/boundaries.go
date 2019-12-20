package core

// Aqui ficam os limites do app, as interfaces que se comunicam com os inputs/outputs

import (
	"github.com/codeforcuritiba/onibus-io-backend/core/model"
)

// Storer é uma interface que representa as estruturas que terão acesso ao banco de dados.
type Storer interface {
	Veiculos() (model.Veiculos, error)
	Veiculo(codigo string) (model.Veiculos, error)
	VeiculosLinha(codigoLinha string) ([]*model.Veiculo, error)
}

// LinhaStorer é uma interface que representa a fronteira de comunicação para se obter as linhas.
type LinhaStorer interface {
	Linhas() ([]*model.Linha, error)
	Linha(codigo string) (*model.Linha, error)
}
