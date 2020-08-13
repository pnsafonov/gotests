package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gotests/tests/grapghl/api2/graph/generated"
	"gotests/tests/grapghl/api2/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) JobToTaskMut(ctx context.Context, arg1 int, arg2 model.Arg2) (*model.Task, error) {
	return r.Resolver.JobToTaskMut(ctx, arg1, arg2)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("todoss")
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Jobs(ctx context.Context) ([]*model.Job, error) {
	return r.Resolver.Jobs(ctx)
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) JobToTaskQuery(ctx context.Context, arg1 int, arg2 *model.Arg2) (*model.Task, error) {
	return r.Resolver.JobToTaskQuery(ctx, arg1, arg2)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) JobToTask1Query(ctx context.Context, arg1 int, arg2 *model.Arg2) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) JobToTask2(ctx context.Context, arg1 int, arg2 model.Arg2) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) JobToTask1(ctx context.Context, arg1 int, arg2 *model.Arg2) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) JobToTask(ctx context.Context, arg1 int, arg2 model.Arg2) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) JobToTask(ctx context.Context, arg1 int, arg2 *model.Arg2) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
