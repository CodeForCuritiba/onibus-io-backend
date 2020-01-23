package business_test

import (
	"errors"
	"testing"

	"github.com/codeforcuritiba/onibus-io-backend/core/business"
	"github.com/codeforcuritiba/onibus-io-backend/core/model"
)

func TestLinhaBusiness(t *testing.T) {
	linhaBO := business.NewLinhaBO(mockStore{})
	t.Run("Retorna as linhas obtidas no banco", func(t *testing.T) {
		linhas, err := linhaBO.Linhas()
		if err != nil {
			t.Errorf("Erro ao obter as linhas: %v", err)
		}
		if len(linhas) != len(mockLinhas) {
			t.Errorf("Esperava-se %d linhas, obteve-se %d", len(mockLinhas), len(linhas))
		}
	})

	t.Run("Busca uma linha passando seu código como parâmetro", func(t *testing.T) {
		want := mockLinhas[1]
		linha, err := linhaBO.Linha(want.Codigo)
		if err != nil {
			t.Fatalf("Erro ao obter a linha: %v", err)
		}
		if linha == nil || linha.Codigo != want.Codigo {
			t.Errorf("Esperava-se linha %v, obteve-se %v", want, linha)
		}
	})

	t.Run("Retorna um erro caso a linha não exista", func(t *testing.T) {
		linha, err := linhaBO.Linha("xyz")
		if err == nil {
			t.Fatalf("Esperava-se um erro")
		}
		if linha != nil {
			t.Errorf("Esperava-se linha %v, obteve-se %v", nil, linha)
		}
	})

	t.Run("Retorna os veiculos de uma linha", func(t *testing.T) {
		veiculos, err := linhaBO.VeiculosLinha("666")
		if err != nil {
			t.Fatalf("Erro ao obter veiculos de uma linha: %v", err)
		}
		if len(veiculos) != len(mockVeiculos) {
			t.Errorf("Esperava-se %d veiculos, obteve-se %d", len(mockVeiculos), len(veiculos))
		}
	})
}

type mockStore struct {
}

func (ms mockStore) Linhas() ([]*model.Linha, error) {
	return mockLinhas, nil
}

func (ms mockStore) Linha(codigo string) (*model.Linha, error) {
	for _, linha := range mockLinhas {
		if linha.Codigo == codigo {
			return linha, nil
		}
	}
	return nil, errors.New("Linha não encontrada")
}

func (ms mockStore) VeiculosLinha(codigoLinha string) (veiculos []*model.Veiculo, err error) {
	for _, veiculo := range mockVeiculos {
		if veiculo.CodigoLinha == codigoLinha {
			veiculos = append(veiculos, veiculo)
		}
	}
	return
}

var mockLinhas = []*model.Linha{
	&model.Linha{
		Codigo: "666",
	},
	&model.Linha{
		Codigo: "667",
	},
	&model.Linha{
		Codigo: "668",
	},
}

var mockVeiculos = []*model.Veiculo{
	&model.Veiculo{
		Codigo:      "123",
		CodigoLinha: "666",
	},
	&model.Veiculo{
		Codigo:      "323",
		CodigoLinha: "666",
	},
}
