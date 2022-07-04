package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
    "context"
    "project-mngt-golang-graphql-gin/configs"
    "project-mngt-golang-graphql-gin/graph/generated"
    "project-mngt-golang-graphql-gin/graph/model"
)

var (
    db = configs.ConnectDB()
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
    project, err := db.CreateProject(&input)

    return project, err
}

func (r *mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (*model.Owner, error) {
    owner, err := db.CreateOwner(&input)

    return owner, err
}

func (r *queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
    owners, err := db.GetOwners()

    return owners, err
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
    projects, err := db.GetProjects()

    return projects, err
}

func (r *queryResolver) Owner(ctx context.Context, input *model.FetchOwner) (*model.Owner, error) {
    owner, err := db.SingleOwner(input.ID)

    return owner, err
}

func (r *queryResolver) Project(ctx context.Context, input *model.FetchProject) (*model.Project, error) {
    project, err := db.SingleProject(input.ID)

    return project, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }