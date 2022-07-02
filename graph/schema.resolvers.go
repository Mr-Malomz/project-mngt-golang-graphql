package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"project-mngt-golang-graphql-gin/graph/generated"
	"project-mngt-golang-graphql-gin/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Owner(ctx context.Context, input *model.FetchOwner) (*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Project(ctx context.Context, input *model.FetchProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
