package resolver

//go:generate go run github.com/99designs/gqlgen
import (
	context "context"
	"log"

	"github.com/codeforcuritiba/onibus-io-backend/graphql"
	"github.com/codeforcuritiba/onibus-io-backend/model"
	"github.com/codeforcuritiba/onibus-io-backend/store"
)

type Resolver struct {
	Store store.Storer
}

func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Linha() graphql.LinhaResolver {
	return &linhaResolver{r}
}

func (l *linhaResolver) Veiculos(ctx context.Context, linha *model.Linha) ([]*model.Veiculo, error) {
	veiculos, err := l.Store.VeiculosLinha(linha.Codigo)
	log.Printf("veiculos: %v, err %v\n", veiculos, err)
	return veiculos, err
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Linhas(ctx context.Context) ([]*model.Linha, error) {
	return r.Store.Linhas()
}

func (r *queryResolver) Linha(ctx context.Context, codigo string) (*model.Linha, error) {
	return r.Store.Linha(codigo)
}

type linhaResolver struct{ *Resolver }
