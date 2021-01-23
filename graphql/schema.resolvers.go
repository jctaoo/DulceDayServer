package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DulceDayServer/graphql/generated"
	"DulceDayServer/graphql/model"
	"context"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todo := &model.Todo{
		ID:   "123",
		Text: "haha",
		Done: false,
		User: &model.User{
			ID:   "1",
			Name: "bob",
		},
	}
	return []*model.Todo{todo}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
