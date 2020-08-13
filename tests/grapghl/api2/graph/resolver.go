package graph

import (
    "context"
    "fmt"
    "gotests/tests/grapghl/api2/graph/model"
)

//go:generate gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *Resolver) Jobs(ctx context.Context) (result []*model.Job, err error) {
    count := 4
    for i := 0; i < count; i++ {
        id := new(int)
        *id = i
        name := new(string)
        *name = fmt.Sprintf("name_%d", i)
        job := &model.Job{
            ID:   id,
            Name: name,
        }
        result = append(result, job)
    }
    return result, nil
}

func (r *Resolver) JobToTaskQuery(ctx context.Context, arg1 int, arg2 *model.Arg2) (task *model.Task, err error) {
    task = &model.Task{
        ID:   5555,
        Name: "JobToTaskQuery",
    }
    return
}

func (r *Resolver) JobToTaskMut(ctx context.Context, arg1 int, arg2 model.Arg2) (task *model.Task, err error) {
    task = &model.Task{
        ID:   5555,
        Name: "JobToTaskMut",
    }
    return
}