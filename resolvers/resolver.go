//go:generate go run github.com/99designs/gqlgen init -v
package resolvers

import (
	"github.com/mind1949/iblog"
)

type Resolver struct{}

func (r *Resolver) Mutation() iblog.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() iblog.QueryResolver {
	return &queryResolver{r}
}
