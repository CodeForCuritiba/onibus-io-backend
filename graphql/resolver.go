package graphql

//go:generate go run github.com/99designs/gqlgen
import (
	context "context"

	"github.com/codeforcuritiba/onibus-io-backend/core/business"
	"github.com/codeforcuritiba/onibus-io-backend/core/model"
)

// Resolver das queries graphql
type Resolver struct {
	LinhaBO *business.LinhaBO
}

// Query resolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Linha resolver
func (r *Resolver) Linha() LinhaResolver {
	return &linhaResolver{r}
}

type queryResolver struct{ *Resolver }

// Linhas query retorna lista de linhas
func (r *queryResolver) Linhas(ctx context.Context) ([]*model.Linha, error) {
	return r.LinhaBO.Linhas()
}

// Linha query retorna linha por código
func (r *queryResolver) Linha(ctx context.Context, codigo string) (*model.Linha, error) {
	return r.LinhaBO.Linha(codigo)
}

type mutationResolver struct{ *Resolver }

type linhaResolver struct{ *Resolver }

// Veiculos resolve o field Veiculos de linhas
func (l *linhaResolver) Veiculos(ctx context.Context, linha *model.Linha) ([]*model.Veiculo, error) {
	veiculos, err := l.LinhaBO.VeiculosLinha(linha.Codigo)
	return veiculos, err
}
