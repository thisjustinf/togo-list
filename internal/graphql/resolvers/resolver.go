package resolvers
//go:generate go run github.com/99designs/gqlgen generate
import (
	"gorm.io/gorm"

	gen "github.com/thisjustinf/togo-list/internal/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func NewRootResolvers(db *gorm.DB) gen.Config {
	c := gen.Config{
		Resolvers: &Resolver{
			DB: db,
		},
	}

	return c
}

// func (r *Resolver) User() UserResolver {
// 	return &userResolver{r}
// }

// func (r *Resolver) Todo() TodoResolver{
// 	return &todoResolver{r}
// }
