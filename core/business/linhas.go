// Package business é onde fica a regra de negócio do sistema.
// A ideia é manter o código limpo e desacoplado das tecnologias de I/O (db, API rest, etc)
package business

import (
	"github.com/codeforcuritiba/onibus-io-backend/core"
	"github.com/codeforcuritiba/onibus-io-backend/core/model"
)

// LinhaBO é o Business Object de Linhas da URBS
type LinhaBO struct {
	store core.LinhaStorer
}

// NewLinhaBO cria um LinhaBO
func NewLinhaBO(store core.LinhaStorer) *LinhaBO {
	return &LinhaBO{store}
}

// Linhas retorna todas as linhas existente
func (l LinhaBO) Linhas() ([]*model.Linha, error) {
	return l.store.Linhas()
}

// Linha retorna uma linha de acordo com seu código
func (l LinhaBO) Linha(codigo string) (*model.Linha, error) {
	return l.store.Linha(codigo)
}

func (l LinhaBO) VeiculosLinha(codigo string) ([]*model.Veiculo, error) {
	return l.store.VeiculosLinha(codigo)
}

func (l LinhaBO) Tabela(codigo, numeroPonto string) ([]*model.Parada, error) {
	return l.store.TabelaLinha(codigo, numeroPonto)
}
