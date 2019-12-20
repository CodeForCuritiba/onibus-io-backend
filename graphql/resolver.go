package graphql

//go:generate go run github.com/99designs/gqlgen
import (
	context "context"

	"github.com/codeforcuritiba/onibus-io-backend/core/business"
	"github.com/codeforcuritiba/onibus-io-backend/core/model"
)

type Resolver struct {
	LinhaBO *business.LinhaBO
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// func (r *Resolver) Linha() LinhaResolver {
// 	return &linhaResolver{r}
// }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Linhas(ctx context.Context) ([]*model.Linha, error) {
	return r.LinhaBO.Linhas()
}

func (r *queryResolver) Linha(ctx context.Context, codigo string) (*model.Linha, error) {
	return r.LinhaBO.Linha(codigo)
}

type linhaResolver struct{ *Resolver }

// func (l *linhaResolver) Veiculos(ctx context.Context, linha *model.Linha) ([]*model.Veiculo, error) {
// 	veiculos, err := l.Store.VeiculosLinha(linha.Codigo)
// 	return veiculos, err
// }
