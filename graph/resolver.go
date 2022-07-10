package graph

import "graphql-go-example/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver() *Resolver {

	users := make([]*model.User, 0)

	users = append(users, &model.User{ID: "1", Name: "fphilip"})

	users = append(users, &model.User{ID: "2", Name: "lturanga"})

	return &Resolver{

		todos: make([]*model.Todo, 0),

		users: users,

		lastUserId: 3,
	}

}

type Resolver struct {
	todos []*model.Todo

	users []*model.User

	lastTodoId int

	lastUserId int
}
